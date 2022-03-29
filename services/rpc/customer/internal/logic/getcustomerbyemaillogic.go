package logic

import (
	"context"
	"time"

	"github.com/k8scommerce/k8scommerce/services/rpc/customer/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/customer/pb/customer"

	"github.com/k8scommerce/k8scommerce/internal/gcache"
	"github.com/k8scommerce/k8scommerce/internal/groupctx"
	"github.com/k8scommerce/k8scommerce/internal/models"
	"github.com/k8scommerce/k8scommerce/internal/utils"

	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const Group_GetCustomerByEmail = "GetCustomerByEmail"

var Group_GetCustomerByEmailKey = func(storeId int64, email string) string {
	return gcache.ToKey(Group_GetCustomerByEmail, storeId, email)
}

type GetCustomerByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCustomerByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCustomerByEmailLogic {
	return &GetCustomerByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCustomerByEmailLogic) GetCustomerByEmail(in *customer.GetCustomerByEmailRequest) (*customer.GetCustomerByEmailResponse, error) {
	l.ctx = groupctx.SetStoreId(l.ctx, in.StoreId)
	l.ctx = groupctx.SetCustomerEmail(l.ctx, in.Email)
	res := &customer.GetCustomerByEmailResponse{}
	err := l.cache().Get(l.ctx, Group_GetCustomerByEmailKey(in.StoreId, in.Email), groupcache.ProtoSink(res))
	return res, err
}

func (l *GetCustomerByEmailLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetCustomerByEmail, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			found, err := l.svcCtx.Repo.Customer().GetCustomerByEmail(
				groupctx.GetStoreId(ctx),
				groupctx.GetCustomerEmail(ctx),
			)
			if err != nil {
				logx.Infof("error: %s", err)
			}

			cust := &customer.Customer{}
			utils.TransformObj(found, &cust)

			// fetch the addresses in parallel
			err = mr.Finish(func() error {
				addresses := getAddressesByKind(l.svcCtx.Repo, cust.Id, models.AddressKindBilling)
				cust.BillingAddresses = addresses
				return nil
			}, func() error {
				addresses := getAddressesByKind(l.svcCtx.Repo, cust.Id, models.AddressKindShipping)
				cust.ShippingAddresses = addresses
				return nil
			})
			if err != nil {
				logx.Error(status.Errorf(codes.Internal, "could not fetch addresses in parallel: %s", err.Error()))
			}

			// Set the groupcache to expire after 24 hours
			if err := dest.SetProto(&customer.GetCustomerByEmailResponse{
				Customer: cust,
			}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
