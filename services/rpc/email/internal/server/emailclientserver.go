// Code generated by goctl. DO NOT EDIT!
// Source: email.proto

package server

import (
	"context"

	"email/internal/logic"
	"email/internal/svc"
	"email/pb/email"

	"github.com/localrivet/galaxycache"
)

type EmailClientServer struct {
	svcCtx   *svc.ServiceContext
	universe *galaxycache.Universe
}

func NewEmailClientServer(svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *EmailClientServer {
	return &EmailClientServer{
		svcCtx:   svcCtx,
		universe: universe,
	}
}

func (s *EmailClientServer) SendOrderConfirmation(ctx context.Context, in *email.SendOrderConfirmationRequest) (*email.Empty, error) {
	l := logic.NewSendOrderConfirmationLogic(ctx, s.svcCtx, s.universe)
	return l.SendOrderConfirmation(in)
}
