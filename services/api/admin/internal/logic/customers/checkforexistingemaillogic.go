package customers

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckForExistingEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckForExistingEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckForExistingEmailLogic {
	return CheckForExistingEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckForExistingEmailLogic) CheckForExistingEmail(req types.CheckForExistingEmailRequest) (resp *types.CheckForExistingEmailResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
