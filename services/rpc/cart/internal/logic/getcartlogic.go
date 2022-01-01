package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/k8scommerce/k8scommerce/pkg/models"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/pb/cart"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type galaxyGetCartLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryCartLogic *galaxyGetCartLogicHelper

type GetCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetCartLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetCartLogic {
	return &GetCartLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetCartLogic) GetCart(in *cart.GetCartRequest) (*cart.GetCartResponse, error) {

	// caching goes logic here
	if entryCartLogic == nil {
		l.mu.Lock()
		entryCartLogic = &galaxyGetCartLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryCartLogic.once.Do(func() {
		fmt.Println(`l.entryCartLogic.Do`)

		// register the galaxy one time

		entryCartLogic.galaxy = gcache.RegisterGalaxyFunc("Cart", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {
				// get the cart from the database
				userId, _ := strconv.ParseInt(key, 10, 64)
				c, err := l.svcCtx.Repo.Cart().GetCartByUserId(int64(userId))
				if err != nil {
					return err
				}

				if c == nil {
					modelCart := &models.Cart{
						UserID: userId,
					}
					response, err := l.svcCtx.Repo.Cart().Create(modelCart)
					if err != nil {
						return err
					}

					c = response
				}

				res := &cart.Cart{
					UserId: c.Cart.UserID,
				}
				var totalPrice int64 = 0
				for _, item := range c.Items {
					res.Items = append(res.Items, &cart.Item{
						UserId:    item.UserID,
						Sku:       item.Sku,
						Quantity:  int32(item.Quantity),
						Price:     item.Price,
						ExpiresAt: timestamppb.New(item.ExpiresAt),
					})

					totalPrice += item.Price
				}

				res.TotalPrice = totalPrice

				out, err := json.Marshal(res)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	res := &cart.GetCartResponse{}

	codec := &galaxycache.ByteCodec{}
	if err := entryCartLogic.galaxy.Get(l.ctx, strconv.Itoa(int(in.UserId)), codec); err != nil {
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
