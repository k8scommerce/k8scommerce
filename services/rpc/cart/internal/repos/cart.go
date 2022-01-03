package repos

import (
	"context"
	"fmt"

	"k8scommerce/internal/models"

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
	Create(cart *models.Cart) (res *CartResponse, err error)
	Update(cart *models.Cart) (res *CartResponse, err error)
	Upsert(cart *models.Cart) (res *CartResponse, err error)
	Delete(userId int64) error
	GetCartByUserId(userId int64) (res *CartResponse, err error)
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

type CartResponse struct {
	Cart  *models.Cart
	Items []models.CartItem
	// Prices []models.Price
}

// products
func (m *cartRepo) GetCartByUserId(userId int64) (res *CartResponse, err error) {
	cart, _ := models.CartByUserID(m.ctx, m.db.DB, userId)
	if cart != nil {
		// get the items if there are any
		var items []models.CartItem
		response, _ := m.repo.CartItem().GetCartItemsByUserId(cart.UserID)
		if response != nil {
			items = append(items, response.Items...)
		}

		res = &CartResponse{
			Cart:  cart,
			Items: items,
		}
	}
	return res, err
}

func (m *cartRepo) Create(cart *models.Cart) (res *CartResponse, err error) {
	out := &CartResponse{}
	if err := cart.Insert(m.ctx, m.db); err != nil {
		return out, err
	}
	out.Cart.UserID = cart.UserID
	return out, nil
}

func (m *cartRepo) Update(cart *models.Cart) (res *CartResponse, err error) {
	if cart.UserID == 0 {
		return nil, fmt.Errorf("error: can't update cart, missing user ID")
	}
	return m.Upsert(cart)
}

func (m *cartRepo) Upsert(cart *models.Cart) (res *CartResponse, err error) {
	_, err = m.db.NamedExec(`
		INSERT INTO cart (
			user_id
		) 
		VALUES (
			:userId
		)
		ON CONFLICT (user_id) DO NOTHING
	`, map[string]interface{}{
		"userId": cart.UserID,
	})

	fmt.Println("ERROR!!!!!!!", err)

	if err != nil {
		return nil, fmt.Errorf("error::Upsert::%s", err.Error())
	}

	out := &CartResponse{
		Cart: &models.Cart{
			UserID: cart.UserID,
		},
	}
	return out, nil
}

func (m *cartRepo) Delete(userID int64) error {
	cart, err := models.CartByUserID(m.ctx, m.db, userID)
	if err != nil {
		return err
	}
	return cart.Delete(m.ctx, m.db)
}
