package assets

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
	"k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadLogic(r *http.Request, ctx context.Context, svcCtx *svc.ServiceContext) UploadLogic {
	return UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadLogic) Upload() (resp *types.Asset, err error) {
	resp = &types.Asset{}

	// we can't use the standard parser
	// as it will cause the MultipartReader to fail
	// so we need to parse the params ourselves
	pathParts := strings.Split(l.r.URL.Path, "/")
	productId, err := strconv.ParseInt(pathParts[len(pathParts)-1], 10, 64)
	if err != nil {
		return
	}
	variantId, err := strconv.ParseInt(pathParts[len(pathParts)-2], 10, 64)
	if err != nil {
		return
	}

	req := types.UploadAssetRequest{
		ProductId: productId,
		VariantId: variantId,
	}

	startTime := time.Now()
	multipartReader, err := l.r.MultipartReader()
	if err != nil {
		return
	}

	var stream catalog.CatalogClient_UploadAssetClient

	partBytes := int64(0)
	partCount := int64(0)
	for {
		//DOS problem .... what if this header is very large?  (Intentionally)
		part, partErr := multipartReader.NextPart()
		if partErr != nil {
			if partErr == io.EOF {
				break //just an eof...not an error
			} else {
				return resp, fmt.Errorf("error getting a part %v", partErr)
			}
		} else {

			if len(part.FileName()) > 0 {
				if partCount == 0 {

					// size := 1024 * 1024 * 12
					// opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(size)))
					// conn, err := grpc.Dial(address, opts...)

					stream, err = l.svcCtx.CatalogRpc.UploadAsset(l.ctx,
						grpc.MaxCallRecvMsgSize(int(l.svcCtx.Config.MaxBytes)),
						grpc.MaxCallSendMsgSize(int(l.svcCtx.Config.MaxBytes)),
					)
					if err != nil {
						return resp, err
					}

					streamReq := &catalog.UploadAssetRequest{
						Data: &catalog.UploadAssetRequest_Asset{
							Asset: &catalog.Asset{
								StoreId:   l.ctx.Value(types.StoreKey).(int64),
								Name:      part.FileName(),
								ProductId: req.ProductId,
								VariantId: req.VariantId,
							},
						},
					}
					err = stream.Send(streamReq)
					if err != nil {
						return resp, fmt.Errorf("cannot send image info to server: %s, %#v", err, stream.RecvMsg(nil))
					}
				}

				// copy the multipart chunk to a buffer
				// then to bytes
				buffer := bytes.NewBuffer(nil)
				if _, err := io.Copy(buffer, part); err != nil {
					return nil, err
				}

				chunkReq := &catalog.UploadAssetRequest{
					Data: &catalog.UploadAssetRequest_Chunk{
						Chunk: buffer.Bytes(),
					},
				}

				err = stream.Send(chunkReq)
				if err != nil {
					return resp, fmt.Errorf("cannot send chunk to server: %s, %#v", err, stream.RecvMsg(nil))
				}

				//Could take an *indefinite* amount of time!!
				// partBytesIncr, partCountIncr := h.serveHTTPUploadPOSTDrain(fileName, w, part)
				// partBytes += partBytesIncr
				// partCount += partCountIncr

				partCount++
			}
		}
	}

	uploadResponse, err := stream.CloseAndRecv()
	if err != nil {
		return resp, fmt.Errorf("cannot receive response: %s, %#v", err, uploadResponse)
	}

	stopTime := time.Now()
	timeDiff := (stopTime.UnixNano()-startTime.UnixNano())/(1000*1000) + 1
	throughput := (1000 * partBytes) / timeDiff
	partSize := int64(0)
	if partCount <= 0 {
		partSize = 0
	} else {
		partSize = partBytes / partCount
	}
	log.Printf("Upload: time = %dms, size = %d B, throughput = %d B/s, partSize = %d B", timeDiff, partBytes, throughput, partSize)

	log.Printf("#%v", uploadResponse)

	return resp, nil
}
