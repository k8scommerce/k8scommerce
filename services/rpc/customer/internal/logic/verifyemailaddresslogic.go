package logic

import (
	"context"
	"k8scommerce/internal/convert"
	"k8scommerce/internal/events/eventkey"
	"k8scommerce/internal/events/eventkey/eventtype"
	"k8scommerce/services/rpc/customer/internal/svc"
	"k8scommerce/services/rpc/customer/pb/customer"
	"net/mail"
	"strconv"
	"strings"

	"github.com/localrivet/galaxycache"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VerifyEmailAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyEmailAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *VerifyEmailAddressLogic {
	return &VerifyEmailAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyEmailAddressLogic) VerifyEmailAddress(in *customer.VerifyEmailAddressRequest) (*customer.VerifyEmailAddressResponse, error) {
	email, storeId, err := l.decodeCode(in.Code)
	if err != nil {
		return nil, err
	}

	if *storeId != in.StoreId {
		return nil, status.Errorf(codes.Internal, "invalid store id: %d", storeId)
	}

	foundCustomer, err := l.svcCtx.Repo.Customer().GetCustomerByEmail(in.StoreId, *email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not find customer by email: %s", *email)
	}

	if foundCustomer != nil {
		foundCustomer.IsVerified = true
		l.svcCtx.Repo.Customer().Update(foundCustomer)
	}

	res := &customer.VerifyEmailAddressResponse{
		Success: true,
	}

	protoCustomer := &customer.Customer{}
	if foundCustomer != nil {
		convert.ModelCustomerToProtoCustomer(foundCustomer, protoCustomer)
	}

	protoStore, err := getProtoStoreByStoreId(l.svcCtx.Repo, in.StoreId)
	if err != nil {
		return nil, err
	}

	evt := &eventtype.CustomerAccountConfirmedEmail{
		Required: &eventtype.Required{
			Customer: protoCustomer,
			Store:    protoStore,
		},
	}

	if bytes, err := eventkey.CustomerAccountConfirmedEmail.Marshal(evt); err != nil {
		logx.Infof("%d: marshaling event %s failed: %s", codes.Internal, eventkey.CustomerAccountConfirmedEmail, err.Error())
		res.Success = false
	} else {
		// publish event
		err = l.svcCtx.EventManager.Publish(eventkey.CustomerAccountConfirmedEmail.AsKey(), bytes)
		if err != nil {
			logx.Infof("%d: publishing event %s failed: %s", codes.Internal, eventkey.CustomerAccountConfirmedEmail, err.Error())
			res.Success = false
		}
	}

	return res, err
}

func (l *VerifyEmailAddressLogic) decodeCode(code string) (email *string, storeId *int64, err error) {
	decodedStr, err := l.svcCtx.Encrypter.Decrypt(code)
	if err != nil {
		return nil, nil, err
	}

	parts := strings.Split(decodedStr, "|")
	if len(parts) != 2 {
		return nil, nil, status.Errorf(codes.Internal, "could not decode code: %s", code)
	}

	if _, err := mail.ParseAddress(parts[0]); err != nil {
		return nil, nil, status.Errorf(codes.Internal, "invalid email address: %s", parts[0])
	}

	id, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return nil, nil, status.Errorf(codes.Internal, "could not decode store id: %s", parts[1])
	}

	return &parts[0], &id, nil
}
