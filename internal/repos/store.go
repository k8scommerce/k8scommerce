package repos

import (
	"context"
	"fmt"

	"k8scommerce/internal/models"

	"github.com/jmoiron/sqlx"
)

func newStore(repo *repo) Store {
	return &storeRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

type Store interface {
	Exists() bool
	Deleted() bool
	Create(store *models.Store) error
	Update(store *models.Store) error
	Save() error
	Upsert() error
	Delete(id int64) error
	GetStoreById(storeId int64) (*models.Store, error)
	// GetAllStores() (*models.Store, error)
}

type storeRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.Store
}

func (m *storeRepo) GetStoreById(storeId int64) (*models.Store, error) {
	customer, err := models.StoreByID(m.ctx, m.db, storeId)
	if err != nil {
		err = &RepoError{
			Err:        err,
			StatusCode: GetStoreByIdErrorCode,
		}
	}
	return customer, err
}

func (m *storeRepo) Create(store *models.Store) error {
	if err := store.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *storeRepo) Update(store *models.Store) error {
	if store.ID == 0 {
		return fmt.Errorf("error: can't update store, missing store ID")
	}
	if err := store.Update(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *storeRepo) Save() error {
	return m.Store.Save(m.ctx, m.db)
}

func (m *storeRepo) Upsert() error {
	return m.Store.Upsert(m.ctx, m.db)
}

func (m *storeRepo) Delete(id int64) error {
	store, err := models.StoreByID(m.ctx, m.db, id)
	if err != nil {
		return err
	}
	return store.Delete(m.ctx, m.db)
}
