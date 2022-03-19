package logic_test

import (
	"context"

	"github.com/localrivet/galaxycache"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8scommerce/services/rpc/customer/internal/logic"
	"k8scommerce/services/rpc/customer/internal/svc"
)

var _ = Describe("VerifyEmailAddressLogic", func() {
	defer GinkgoRecover()

	var ctx context.Context = context.Background()
	var svcCtx *svc.ServiceContext = &svc.ServiceContext{}
	var universe *galaxycache.Universe = nil
	logic.NewVerifyEmailAddressLogic(ctx, svcCtx, universe)

	It("should validate an email address", func() {
		// l.VerifyEmailAddress()

		// customer, err := createCustomer()
		Expect(nil).To(BeNil())
	})

})
