package repos

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/k8scommerce/k8scommerce/internal/models"

	"github.com/google/uuid"
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
	Delete(cartId uuid.UUID, sku string, force bool) error

	GetByCartId(cartId uuid.UUID) ([]*models.CartItem, error)
	ClearItems(cartId uuid.UUID, force bool) (err error)
	AddItem(cartId uuid.UUID, item *models.CartItem) (res *models.CartItem, err error)
	UpdateQuantity(cartId uuid.UUID, sku string, quantity int) (res *models.CartItem, err error)
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

// cart items
func (m *cartItemRepo) GetByCartId(cartId uuid.UUID) ([]*models.CartItem, error) {
	nstmt, err := m.db.PrepareNamed(`
		SELECT 
			cart_id,
			sku,
			quantity,
			price,
			expires_at
		FROM cart_item
		WHERE cart_id = :cartId
		AND abandoned_at IS NULL
	`)
	if err != nil {
		return nil, fmt.Errorf("error::GetCartItemsByCustomerId::%s", err.Error())
	}

	var items []*models.CartItem

	err = nstmt.Select(&items,
		map[string]interface{}{
			"cartId": cartId,
		})

	return items, err
}

func (m *cartItemRepo) AddItem(cartId uuid.UUID, item *models.CartItem) (res *models.CartItem, err error) {
	nstmt, err := m.db.PrepareNamed(`
		INSERT INTO cart_item (
			cart_id,
			sku,
			quantity,
			price,
			expires_at
		) 
		VALUES (
			:cart_id, 
			:sku,
			:quantity,
			:price,
			:expires_at
		)
		ON CONFLICT (cart_id, sku) DO UPDATE 
		SET quantity = cart_item.quantity
		RETURNING 
			cart_id,
			sku,
			quantity,
			price,
			expires_at
	`)
	if err != nil {
		return nil, fmt.Errorf("error::AddItem::%s", err.Error())
	}

	var response []*models.CartItem

	err = nstmt.Select(&response,
		map[string]interface{}{
			"cart_id":    cartId,
			"sku":        item.Sku,
			"quantity":   item.Quantity,
			"price":      item.Price,
			"expires_at": item.ExpiresAt,
		})

	if len(response) > 0 {
		return response[0], nil
	}

	return nil, err
}

func (m *cartItemRepo) UpdateQuantity(cartId uuid.UUID, sku string, quantity int) (res *models.CartItem, err error) {
	nstmt, err := m.db.PrepareNamed(`
		UPDATE cart_item SET quantity = :quantity
			WHERE cart_id = :cart_id
			AND sku = :sku
		RETURNING 
			cart_id,
			sku,
			quantity,
			price,
			expires_at
	`)
	if err != nil {
		return nil, fmt.Errorf("error::UpdateQuantity::%s", err.Error())
	}

	var response []*models.CartItem

	err = nstmt.Select(&response,
		map[string]interface{}{
			"cart_id":  cartId,
			"sku":      sku,
			"quantity": quantity,
		})

	if len(response) > 0 {
		return response[0], nil
	}

	return nil, err
}

func (m *cartItemRepo) Create(cartItem *models.CartItem) error {
	if err := cartItem.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *cartItemRepo) Update(cartItem *models.CartItem) error {
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

func (m *cartItemRepo) Delete(cartId uuid.UUID, sku string, force bool) error {
	cartItem, err := models.CartItemByCartIDSku(m.ctx, m.db, cartId, sku)
	if err != nil {
		return err
	}

	if force {
		err = cartItem.Delete(m.ctx, m.db)
	} else {
		cartItem.AbandonedAt = sql.NullTime{Time: time.Now(), Valid: true}
		err = cartItem.Upsert(m.ctx, m.db)
	}
	return err
}

func (m *cartItemRepo) ClearItems(cartId uuid.UUID, force bool) (err error) {
	var nstmt *sqlx.NamedStmt
	if force {
		nstmt, err = m.db.PrepareNamed(`
			DELETE FROM cart_item WHERE cart_id = :cartId
		`)
		if err != nil {
			return fmt.Errorf("error::ClearItems with force::%s", err.Error())
		}
		_, err = nstmt.Exec(map[string]interface{}{
			"cartId": cartId,
		})
	} else {
		nstmt, err = m.db.PrepareNamed(`
			UPDATE cart_item SET abandoned_at = :abandonedAt
			WHERE cart_id = :cartId
		`)
		if err != nil {
			return fmt.Errorf("error::ClearItems without force::%s", err.Error())
		}
		_, err = nstmt.Exec(map[string]interface{}{
			"cartId":      cartId,
			"abandonedAt": time.Now(),
		})
	}

	return err
}
