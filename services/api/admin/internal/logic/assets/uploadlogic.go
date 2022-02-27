package assets

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"k8scommerce/internal/storage/asset"
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
	// path looks like this /v1/asset/:productId/:variantId/:kind
	// fwe pop them off in reverse
	// image first
	// variantId second
	// productId third
	pathParts := strings.Split(l.r.URL.Path, "/")

	kind := pathParts[len(pathParts)-1]
	if err != nil {
		return
	}

	variantId, err := strconv.ParseInt(pathParts[len(pathParts)-2], 10, 64)
	if err != nil {
		return
	}

	productId, err := strconv.ParseInt(pathParts[len(pathParts)-3], 10, 64)
	if err != nil {
		return
	}

	req := types.UploadAssetRequest{
		ProductId: productId,
		VariantId: variantId,
		Kind:      kind,
	}

	// startTime := time.Now()
	multipartReader, err := l.r.MultipartReader()
	if err != nil {
		return
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	stream, err := l.svcCtx.CatalogRpc.UploadAsset(ctx,
		grpc.MaxCallRecvMsgSize(int(l.svcCtx.Config.MaxBytes)),
		grpc.MaxCallSendMsgSize(int(l.svcCtx.Config.MaxBytes)),
	)
	if err != nil {
		return resp, err
	}

	partCount := int64(0)
	for {
		part, partErr := multipartReader.NextPart()
		if partErr != nil {
			if partErr == io.EOF {
				break
			} else {
				return resp, fmt.Errorf("error getting a part %v", partErr)
			}
		} else {
			if len(part.FileName()) > 0 {
				if partCount == 0 {
					kind := asset.Kind(req.Kind)

					streamReq := &catalog.UploadAssetRequest{
						Data: &catalog.UploadAssetRequest_Asset{
							Asset: &catalog.Asset{
								StoreId:   l.ctx.Value(types.StoreKey).(int64),
								Name:      part.FileName(),
								ProductId: req.ProductId,
								VariantId: req.VariantId,
								Kind:      catalog.AssetKind(kind.Int32()),
							},
						},
					}
					err = stream.Send(streamReq)
					if err != nil {
						return resp, fmt.Errorf("cannot send image info to server: %s", err)
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
					break
				}

				partCount++
			}
		}
	}

	uploadResponse, err := stream.CloseAndRecv()
	if err != nil {
		return resp, fmt.Errorf("cannot receive response: %s, %#v", err, uploadResponse)
	}

	resp.Id = uploadResponse.Id
	resp.Kind = int(uploadResponse.Kind.Number())
	resp.Name = uploadResponse.Name
	resp.ProductId = uploadResponse.ProductId
	resp.VariantId = uploadResponse.VariantId
	resp.Url = uploadResponse.Url
	resp.ContentType = uploadResponse.ContentType
	resp.DisplayName = uploadResponse.DisplayName
	resp.SortOrder = uploadResponse.SortOrder

	// stopTime := time.Now()
	// timeDiff := (stopTime.UnixNano()-startTime.UnixNano())/(1000*1000) + 1
	// throughput := (1000 * partBytes) / timeDiff
	// partSize := int64(0)
	// if partCount <= 0 {
	// 	partSize = 0
	// } else {
	// 	partSize = partBytes / partCount
	// }
	// log.Printf("Upload: time = %dms, size = %d B, throughput = %d B/s, partSize = %d B", timeDiff, partBytes, throughput, partSize)

	// log.Printf("#%v", uploadResponse
	return resp, nil
}
