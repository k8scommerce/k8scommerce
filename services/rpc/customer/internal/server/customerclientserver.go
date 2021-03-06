// Code generated by goctl. DO NOT EDIT!
// Source: customer.proto

package server

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/customer/internal/logic"
	"github.com/k8scommerce/k8scommerce/services/rpc/customer/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/customer/pb/customer"
)

type CustomerClientServer struct {
	svcCtx *svc.ServiceContext
	customer.UnimplementedCustomerClientServer
}

func NewCustomerClientServer(svcCtx *svc.ServiceContext) *CustomerClientServer {
	return &CustomerClientServer{
		svcCtx: svcCtx,
	}
}

func (s *CustomerClientServer) CreateCustomer(ctx context.Context, in *customer.CreateCustomerRequest) (*customer.CreateCustomerResponse, error) {
	l := logic.NewCreateCustomerLogic(ctx, s.svcCtx)
	return l.CreateCustomer(in)
}

func (s *CustomerClientServer) UpdateCustomer(ctx context.Context, in *customer.UpdateCustomerRequest) (*customer.UpdateCustomerResponse, error) {
	l := logic.NewUpdateCustomerLogic(ctx, s.svcCtx)
	return l.UpdateCustomer(in)
}

func (s *CustomerClientServer) GetCustomerByEmail(ctx context.Context, in *customer.GetCustomerByEmailRequest) (*customer.GetCustomerByEmailResponse, error) {
	l := logic.NewGetCustomerByEmailLogic(ctx, s.svcCtx)
	return l.GetCustomerByEmail(in)
}

func (s *CustomerClientServer) SetPassword(ctx context.Context, in *customer.SetPasswordRequest) (*customer.SetPasswordResponse, error) {
	l := logic.NewSetPasswordLogic(ctx, s.svcCtx)
	return l.SetPassword(in)
}

func (s *CustomerClientServer) Login(ctx context.Context, in *customer.LoginRequest) (*customer.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *CustomerClientServer) SendForgotPasswordEmail(ctx context.Context, in *customer.SendForgotPasswordEmailRequest) (*customer.SendForgotPasswordEmailResponse, error) {
	l := logic.NewSendForgotPasswordEmailLogic(ctx, s.svcCtx)
	return l.SendForgotPasswordEmail(in)
}

func (s *CustomerClientServer) SendConfirmEmailAddressEmail(ctx context.Context, in *customer.SendConfirmEmailAddressEmailRequest) (*customer.SendConfirmEmailAddressEmailResponse, error) {
	l := logic.NewSendConfirmEmailAddressEmailLogic(ctx, s.svcCtx)
	return l.SendConfirmEmailAddressEmail(in)
}

func (s *CustomerClientServer) VerifyEmailAddress(ctx context.Context, in *customer.VerifyEmailAddressRequest) (*customer.VerifyEmailAddressResponse, error) {
	l := logic.NewVerifyEmailAddressLogic(ctx, s.svcCtx)
	return l.VerifyEmailAddress(in)
}
