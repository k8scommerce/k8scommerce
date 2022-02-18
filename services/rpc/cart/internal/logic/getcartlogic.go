package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/cart/internal/svc"
	"k8scommerce/services/rpc/cart/pb/cart"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/zeromicro/go-zero/core/logx"
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
				customerId, _ := strconv.ParseInt(key, 10, 64)
				c, err := l.svcCtx.Repo.Cart().GetCartByCustomerId(int64(customerId))
				if err != nil {
					return err
				}

				if c == nil {
					modelCart := &models.Cart{
						CustomerID: customerId,
					}
					response, err := l.svcCtx.Repo.Cart().Create(modelCart)
					if err != nil {
						return err
					}

					c = response
				}

				res := &cart.Cart{
					CustomerId: c.Cart.CustomerID,
				}
				var totalPrice int64 = 0
				for _, item := range c.Items {
					res.Items = append(res.Items, &cart.Item{
						CustomerId: item.CustomerID,
						Sku:        item.Sku,
						Quantity:   int32(item.Quantity),
						Price:      item.Price,
						ExpiresAt:  timestamppb.New(item.ExpiresAt),
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
	if err := entryCartLogic.galaxy.Get(l.ctx, strconv.Itoa(int(in.CustomerId)), codec); err != nil {
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
