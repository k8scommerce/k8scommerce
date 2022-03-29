package repos

import (
	"context"
	"fmt"

	"github.com/k8scommerce/k8scommerce/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func newCart(repo *repo) Cart {
	return &cartRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

type Cart interface {
	SetCart(cart *models.Cart)
	Exists() bool
	Deleted() bool
	Create(cart *models.Cart) (err error)
	Update(cart *models.Cart) (err error)
	Delete(cartId string) error

	GetByCartId(cartId uuid.UUID) (res *models.Cart, err error)
	// AttachCustomer(AttachCustomerRequest) returns (AttachCustomerResponse);
	// rpc UpdateCustomerDetail(UpdateCustomerDetailRequest)
	// 	returns (UpdateCustomerDetailResponse);
	// rpc UpdateStatus(UpdateStatusRequest) returns (UpdateStatusResponse);

	// rpc GetById(GetByIdRequest) returns (GetByIdResponse);
	// rpc GetBySession(GetBySessionRequest) returns (GetBySessionResponse);

	// rpc AddItem(AddItemRequest) returns (AddItemResponse);
	// rpc UpdateItemQuantity(UpdateItemQuantityRequest)
	// 	returns (UpdateItemQuantityResponse);
	// rpc RemoveItem(RemoveItemRequest) returns (RemoveItemResponse);
	// rpc ClearCart(ClearCartRequest) returns (ClearCartResponse);
}

type cartRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.Cart
}

func (m *cartRepo) SetCart(cart *models.Cart) {
	m.Cart = cart
}

func (m *cartRepo) Create(cart *models.Cart) (err error) {
	// set uuid
	cart.ID = uuid.New()

	// set the status to new
	cart.Status = models.CartStatusNew

	// set the default billing and shipping addresses
	cart.BillingAddress = []byte("{}")
	cart.ShippingAddress = []byte("{}")

	if err := cart.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *cartRepo) GetByCartId(cartId uuid.UUID) (res *models.Cart, err error) {
	cart, err := models.CartByID(m.ctx, m.db, cartId)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (m *cartRepo) Update(cart *models.Cart) (err error) {
	if cart.ID.String() == "" {
		return fmt.Errorf("error: can't update cart, missing ID")
	}
	return cart.Update(m.ctx, m.db)
}

func (m *cartRepo) Delete(cartId string) error {
	// cart, err := models.CartByCustomerID(m.ctx, m.db, userID)
	// if err != nil {
	// 	return err
	// }
	// return cart.Delete(m.ctx, m.db)
	return nil
}
