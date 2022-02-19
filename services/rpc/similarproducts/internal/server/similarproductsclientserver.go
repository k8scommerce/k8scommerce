// Code generated by goctl. DO NOT EDIT!
// Source: similarproducts.proto

package server

import (
	"context"

	"k8scommerce/services/rpc/similarproducts/internal/logic"
	"k8scommerce/services/rpc/similarproducts/internal/svc"
	"k8scommerce/services/rpc/similarproducts/pb/similarproducts"
)

type SimilarProductsClientServer struct {
	svcCtx *svc.ServiceContext
	similarproducts.UnimplementedSimilarProductsClientServer
}

func NewSimilarProductsClientServer(svcCtx *svc.ServiceContext) *SimilarProductsClientServer {
	return &SimilarProductsClientServer{
		svcCtx: svcCtx,
	}
}

func (s *SimilarProductsClientServer) GetSimilarProductsBySku(ctx context.Context, in *similarproducts.GetSimilarProductsBySkuRequest) (*similarproducts.GetSimilarProductsBySkuResponse, error) {
	l := logic.NewGetSimilarProductsBySkuLogic(ctx, s.svcCtx)
	return l.GetSimilarProductsBySku(in)
}
