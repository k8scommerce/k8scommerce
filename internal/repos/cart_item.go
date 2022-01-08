package repos

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"k8scommerce/internal/models"

	"github.com/jmoiron/sqlx"
)

func newCartItem(repo *repo) CartItem {
	return &cartItemRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

type CartItem interface {
	SetCartItem(cartItem *models.CartItem)
	Exists() bool
	Deleted() bool
	Create(cartItem *models.CartItem) error
	Update(cartItem *models.CartItem) error
	Save(cartItem *models.CartItem) error
	Upsert(cartItem *models.CartItem) error
	Delete(customerId int64, sku string, force bool) error
	GetCartItemsByCustomerId(customerId int64) (res *cartItemResponse, err error)
	AddItem(cartId int64, item *models.CartItem) (res *cartItemResponse, err error)
}

type cartItemRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.CartItem
}

func (m *cartItemRepo) SetCartItem(cartItem *models.CartItem) {
	m.CartItem = cartItem
}

type cartItemResponse struct {
	Items []models.CartItem
	// Prices []models.Price
}

// cart items
func (m *cartItemRepo) GetCartItemsByCustomerId(customerId int64) (res *cartItemResponse, err error) {
	nstmt, err := m.db.PrepareNamed(`
		SELECT 
			user_id,
			sku,
			quantity,
			price,
			expires_at
		FROM cart_item
		WHERE user_id = :customerId
		AND abandoned_at IS NULL
	`)
	if err != nil {
		return nil, fmt.Errorf("error::GetCartItemsByCustomerId::%s", err.Error())
	}

	var items []models.CartItem

	err = nstmt.Select(&items,
		map[string]interface{}{
			"customerId": customerId,
		})

	out := &cartItemResponse{
		Items: items,
	}
	return out, err
}

func (m *cartItemRepo) AddItem(customerId int64, item *models.CartItem) (res *cartItemResponse, err error) {
	nstmt, err := m.db.PrepareNamed(`
		INSERT INTO cart_item (
			user_id,
			sku,
			quantity,
			price,
			expires_at
		) 
		VALUES (
			:user_id, 
			:sku,
			:quantity,
			:price,
			:expires_at
		)
		ON CONFLICT (user_id, sku) DO UPDATE 
		SET quantity = cart_item.quantity + excluded.quantity
		RETURNING 
		user_id,
			sku,
			quantity,
			price,
			expires_at
	`)
	if err != nil {
		return nil, fmt.Errorf("error::AddItem::%s", err.Error())
	}

	var items []models.CartItem

	err = nstmt.Select(&items,
		map[string]interface{}{
			"user_id":    customerId,
			"sku":        item.Sku,
			"quantity":   item.Quantity,
			"price":      item.Price,
			"expires_at": item.ExpiresAt,
		})

	out := &cartItemResponse{
		Items: items,
	}
	return out, err
}

func (m *cartItemRepo) Create(cartItem *models.CartItem) error {
	if err := cartItem.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *cartItemRepo) Update(cartItem *models.CartItem) error {
	if cartItem.CustomerID == 0 {
		return fmt.Errorf("error: can't update cart item, missing customerId")
	}

	if cartItem.Sku == "" {
		return fmt.Errorf("error: can't update cart item, missing sku")
	}

	if err := cartItem.Update(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *cartItemRepo) Save(cartItem *models.CartItem) error {
	if err := cartItem.Save(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *cartItemRepo) Upsert(cartItem *models.CartItem) error {
	if err := cartItem.Upsert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *cartItemRepo) Delete(customerId int64, sku string, force bool) error {
	cart, err := models.CartItemByCustomerIDSku(m.ctx, m.db, customerId, sku)
	if err != nil {
		return err
	}

	if force {
		err = cart.Delete(m.ctx, m.db)
	} else {
		cart.AbandonedAt = sql.NullTime{Time: time.Now(), Valid: true}
		err = cart.Upsert(m.ctx, m.db)
	}
	return err
}
