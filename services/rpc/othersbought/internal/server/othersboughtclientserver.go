// Code generated by goctl. DO NOT EDIT!
// Source: othersbought.proto

package server

import (
	"context"

	"k8scommerce/services/rpc/othersbought/internal/logic"
	"k8scommerce/services/rpc/othersbought/internal/svc"
	"k8scommerce/services/rpc/othersbought/pb/othersbought"

	"github.com/localrivet/galaxycache"
)

type OthersBoughtClientServer struct {
	svcCtx   *svc.ServiceContext
	universe *galaxycache.Universe
	othersbought.UnimplementedOthersBoughtClientServer
}

func NewOthersBoughtClientServer(svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *OthersBoughtClientServer {
	return &OthersBoughtClientServer{
		svcCtx:   svcCtx,
		universe: universe,
	}
}

func (s *OthersBoughtClientServer) GetOthersBoughtBySku(ctx context.Context, in *othersbought.GetOthersBoughtBySkuRequest) (*othersbought.GetOthersBoughtBySkuResponse, error) {
	l := logic.NewGetOthersBoughtBySkuLogic(ctx, s.svcCtx, s.universe)
	return l.GetOthersBoughtBySku(in)
}
