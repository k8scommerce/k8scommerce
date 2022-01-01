package logic

import (
	"context"
	"email/internal/svc"
	"email/pb/email"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxySendOrderConfirmationLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entrySendOrderConfirmationLogic *galaxySendOrderConfirmationLogicHelper

type SendOrderConfirmationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewSendOrderConfirmationLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *SendOrderConfirmationLogic {
	return &SendOrderConfirmationLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *SendOrderConfirmationLogic) SendOrderConfirmation(in *email.SendOrderConfirmationRequest) (*email.Empty, error) {

	// caching goes logic here
	if entrySendOrderConfirmationLogic == nil {
		l.mu.Lock()
		entrySendOrderConfirmationLogic = &galaxySendOrderConfirmationLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entrySendOrderConfirmationLogic.once.Do(func() {
		fmt.Println(`l.entrySendOrderConfirmationLogic.Do`)

		// register the galaxy one time
		entrySendOrderConfirmationLogic.galaxy = gcache.RegisterGalaxyFunc("SendOrderConfirmation", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {
				// todo: add your logic here and delete this line
				fmt.Printf("Looking up SendOrderConfirmation record by key: %s", key)

				// uncomment below to get the item from the adapter
				// found, err := l.ca.GetProductBySku(key)
				// if err != nil {
				//	logx.Infof("error: %s", err)
				//	return err
				// }

				// the response struct
				item := &email.Empty{}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	res := &email.Empty{}

	codec := &galaxycache.ByteCodec{}
	if err := entrySendOrderConfirmationLogic.galaxy.Get(l.ctx, in.Id, codec); err != nil {
		res.StatusCode = http.StatusNoContent
		res.StatusMessage = err.Error()
		return res, nil
	}

	b, err := codec.MarshalBinary()
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.StatusMessage = err.Error()
		return res, nil
	}

	err = json.Unmarshal(b, res)
	return res, err

}
