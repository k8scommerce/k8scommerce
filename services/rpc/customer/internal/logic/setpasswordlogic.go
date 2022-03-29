package logic

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	"github.com/k8scommerce/k8scommerce/services/rpc/customer/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/customer/pb/customer"
	"github.com/k8scommerce/k8scommerce/services/rpc/store/pb/store"

	"github.com/k8scommerce/k8scommerce/internal/events/eventkey"
	"github.com/k8scommerce/k8scommerce/internal/events/eventkey/eventtype"
	"github.com/k8scommerce/k8scommerce/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetPasswordLogic {
	return &SetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetPasswordLogic) SetPassword(in *customer.SetPasswordRequest) (*customer.SetPasswordResponse, error) {
	res := &customer.SetPasswordResponse{
		Success: false,
	}

	code, err := l.svcCtx.Encrypter.Decrypt(in.Code)

	if err != nil {
		return res, status.Errorf(codes.FailedPrecondition, "could not decode code: %s", err.Error())
	}

	parts := strings.Split(code, "|")
	if len(parts) < 1 {
		return res, status.Errorf(codes.FailedPrecondition, "could not parse decoded code: %s", err.Error())
	}

	email := parts[0]
	storeId, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return res, status.Errorf(codes.FailedPrecondition, "could not parse store id: %s", err.Error())
	}

	if in.StoreId != storeId {
		return res, status.Errorf(codes.FailedPrecondition, "could not parse decoded code: %s", err.Error())
	}

	// get customer by email first
	found, err := l.svcCtx.Repo.Customer().GetCustomerByEmail(in.StoreId, email)
	if err != nil {
		if !strings.Contains(err.Error(), "sql: no rows in result set") {
			return res, status.Errorf(codes.FailedPrecondition, "could not find customer by email: %s", err.Error())
		}
	}

	if found.ID == 0 {
		return res, status.Errorf(codes.FailedPrecondition, "could not find customer: %s", err.Error())
	}

	firstSetup := false

	if !found.Password.Valid {
		firstSetup = true
	}

	// add the password
	found.Password = sql.NullString{String: in.Password, Valid: true}

	// make sure we're verified
	found.IsVerified = true

	// update the customer
	if err := l.svcCtx.Repo.Customer().Update(found); err != nil {
		return res, status.Errorf(codes.FailedPrecondition, "could not update the customer: %s", err.Error())
	}

	out := &customer.Customer{}
	utils.TransformObj(found, &out)

	required := eventtype.NewRequired().Prepare(l.svcCtx.Repo, &struct {
		StoreId       int64
		CustomerEmail string
		Customer      *customer.Customer
		Store         *store.Store
		StoreSetting  *store.StoreSetting
	}{
		in.StoreId,
		email,
		out,
		nil,
		nil,
	})

	// a new account was created
	if firstSetup {
		if bytes, err := eventkey.CustomerNewAccount.Marshal(&eventtype.CustomerNewAccount{Required: required}); err != nil {
			logx.Infof("%d: marshaling event %s failed: %s", codes.Internal, eventkey.CustomerNewAccount, err.Error())
		} else {
			// publish event
			err = l.svcCtx.EventManager.Publish(eventkey.CustomerNewAccount.AsKey(), bytes)
			if err != nil {
				logx.Infof("%d: publishing event %s failed: %s", codes.Internal, eventkey.CustomerNewAccount, err.Error())
			}
		}
	} else {
		// the password was updated
		if bytes, err := eventkey.CustomerPasswordChanged.Marshal(&eventtype.CustomerPasswordChanged{Required: required}); err != nil {
			logx.Infof("%d: marshaling event %s failed: %s", codes.Internal, eventkey.CustomerPasswordChanged, err.Error())
		} else {
			// publish event
			err = l.svcCtx.EventManager.Publish(eventkey.CustomerPasswordChanged.AsKey(), bytes)
			if err != nil {
				logx.Infof("%d: publishing event %s failed: %s", codes.Internal, eventkey.CustomerPasswordChanged, err.Error())
			}
		}
	}

	res.Success = true
	res.Customer = out
	return res, nil

}
