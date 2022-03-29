package logic

import (
	"context"
	"k8scommerce/internal/convert"
	"k8scommerce/internal/events/eventkey"
	"k8scommerce/internal/events/eventkey/eventtype"
	"k8scommerce/services/rpc/customer/internal/svc"
	"k8scommerce/services/rpc/customer/pb/customer"
	"k8scommerce/services/rpc/store/pb/store"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
)

type SendConfirmEmailAddressEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendConfirmEmailAddressEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendConfirmEmailAddressEmailLogic {
	return &SendConfirmEmailAddressEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendConfirmEmailAddressEmailLogic) SendConfirmEmailAddressEmail(in *customer.SendConfirmEmailAddressEmailRequest) (*customer.SendConfirmEmailAddressEmailResponse, error) {
	res := &customer.SendConfirmEmailAddressEmailResponse{
		Success: true,
	}

	foundCustomer, err := l.svcCtx.Repo.Customer().GetCustomerByEmail(in.StoreId, in.Email)
	if err != nil {
		return nil, err
	}

	protoCustomer := &customer.Customer{}
	if foundCustomer != nil {
		convert.ModelCustomerToProtoCustomer(foundCustomer, protoCustomer)
	}

	code, err := l.svcCtx.Encrypter.Encrypt(strings.Join([]string{protoCustomer.Email, strconv.FormatInt(protoCustomer.StoreId, 10)}, "|"))
	if err != nil {
		return nil, err
	}

	evt := &eventtype.CustomerConfirmationEmail{
		Required: eventtype.NewRequired().Prepare(l.svcCtx.Repo, &struct {
			StoreId       int64
			CustomerEmail string
			Customer      *customer.Customer
			Store         *store.Store
			StoreSetting  *store.StoreSetting
		}{
			in.StoreId,
			in.Email,
			protoCustomer,
			nil,
			nil,
		}),
		Code: code,
	}

	if bytes, err := eventkey.CustomerConfirmationEmail.Marshal(evt); err != nil {
		logx.Infof("%d: marshaling event %s failed: %s", codes.Internal, eventkey.CustomerConfirmationEmail, err.Error())
		res.Success = false
	} else {
		// publish event
		err = l.svcCtx.EventManager.Publish(eventkey.CustomerConfirmationEmail.AsKey(), bytes)
		if err != nil {
			logx.Infof("%d: publishing event %s failed: %s", codes.Internal, eventkey.CustomerConfirmationEmail, err.Error())
			res.Success = false
		}
	}

	return res, err
}
