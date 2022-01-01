package gateway

type Gateway interface {
	Auth(req AuthRequest) (AuthResponse, error)
	Sale(req SaleRequest) (SaleResponse, error)
	Refund(req RefundRequest) (RefundResponse, error)
	Credit(req CreditRequest) (CreditResponse, error)
}

//
// Authentication
//

type AuthRequest struct {
	Username string
	Password string
	Key      string
	Secret   string
}

type AuthResponse struct {
}

//
// Sale
//

type SaleRequest struct {
}

type SaleResponse struct {
}

//
// Refund
//

type RefundRequest struct {
}

type RefundResponse struct {
}

//
// Credit
//

type CreditRequest struct {
}

type CreditResponse struct {
}
