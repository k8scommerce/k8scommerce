package paypal

import (
	"context"
	"os"

	"github.com/k8s-commerce/k8s-commerce/services/rpc/payment/internal/config"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/payment/internal/gateway"
	"github.com/plutov/paypal/v4"
)

func NewPaypalGateway(c *config.Config) gateway.Gateway {
	return &paypalGateway{
		pp:  &paypal.Client{},
		ctx: context.Background(),
		cfg: c,
	}
}

type paypalGateway struct {
	pp  *paypal.Client
	ctx context.Context
	cfg *config.Config
}

func (g *paypalGateway) Auth(req gateway.AuthRequest) (gateway.AuthResponse, error) {

	// Create a client instance
	c, err := paypal.NewClient("clientID", "secretID", paypal.APIBaseSandBox)
	c.SetLog(os.Stdout) // Set log to terminal stdout
	g.pp.Client = c.Client

	// accessToken, err := c.GetAccessToken(context.Background())

	return gateway.AuthResponse{}, err
}

func (g *paypalGateway) Sale(req gateway.SaleRequest) (gateway.SaleResponse, error) {
	return gateway.SaleResponse{}, nil
}

func (g *paypalGateway) Refund(req gateway.RefundRequest) (gateway.RefundResponse, error) {
	return gateway.RefundResponse{}, nil
}

func (g *paypalGateway) Credit(req gateway.CreditRequest) (gateway.CreditResponse, error) {
	return gateway.CreditResponse{}, nil
}
