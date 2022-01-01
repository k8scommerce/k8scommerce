package logic

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/k8s-commerce/k8s-commerce/services/rpc/cart/internal/svc"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/cart/pb/cart"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/mr"
)

type galaxyClearCartLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryClearCartLogic *galaxyClearCartLogicHelper

type ClearCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewClearCartLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *ClearCartLogic {
	return &ClearCartLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *ClearCartLogic) ClearCart(in *cart.ClearCartRequest) (*cart.ClearCartResponse, error) {
	err := mr.Finish(func() error {
		// clear the existing cart
		if entryCartLogic != nil {
			l.mu.Lock()
			err := entryCartLogic.galaxy.Remove(l.ctx, strconv.FormatInt(in.UserId, 10))
			l.mu.Unlock()
			return fmt.Errorf("error: deleting cart cache: %s", err.Error())
		}
		return nil
	})
	if err != nil {
		log.Printf("clear cart error: %v", err)
		return nil, err
	}

	res := &cart.ClearCartResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "",
	}

	return res, nil
}
