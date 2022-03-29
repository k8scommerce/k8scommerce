package logic_test

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/customer/pb/customer"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Setpasswordlogic", func() {
	It("should decode the code and add the password", func() {
		in := &customer.SetPasswordRequest{
			StoreId:  1,
			Code:     "5DK7tC3Ln54d71I7914xM3gg7YH6U53Cz5sQ3I17nx",
			Password: "Pass#123",
		}
		ctx := context.Background()
		response, err := srv.SetPassword(ctx, in)
		Expect(err).To(BeNil())
		Expect(response).ToNot(BeNil())
	})
})
