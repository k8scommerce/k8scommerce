package stripe

import (
	"github.com/k8s-commerce/k8s-commerce/services/rpc/payment/internal/config"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/payment/internal/gateway"
	"github.com/stripe/stripe-go/client"
)

func NewStripeGateway(c *config.Config) gateway.Gateway {
	return &stripeGateway{
		sc:  &client.API{},
		cfg: c,
	}
}

type stripeGateway struct {
	sc  *client.API
	cfg *config.Config
}

func (g *stripeGateway) Auth(req gateway.AuthRequest) (gateway.AuthResponse, error) {
	// p := params.(*stripe.ChargeParams)
	g.sc.Init("sk_test_4eC39HqLyjWDarjtT1zdp7dc", nil)
	// g.sc.Charges.Get("ch_3KCgH42eZvKYlo2C15ubypV3", p)
	return gateway.AuthResponse{}, nil
}

func (g *stripeGateway) Sale(req gateway.SaleRequest) (gateway.SaleResponse, error) {
	return gateway.SaleResponse{}, nil
}

func (g *stripeGateway) Refund(req gateway.RefundRequest) (gateway.RefundResponse, error) {
	return gateway.RefundResponse{}, nil
}

func (g *stripeGateway) Credit(req gateway.CreditRequest) (gateway.CreditResponse, error) {
	return gateway.CreditResponse{}, nil
}
