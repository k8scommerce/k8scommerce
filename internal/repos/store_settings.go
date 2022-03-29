package repos

import (
	"context"
	"fmt"

	"github.com/k8scommerce/k8scommerce/internal/models"

	"github.com/jmoiron/sqlx"
)

func newStoreSetting(repo *repo) StoreSetting {
	return &storeSettingRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

type StoreSetting interface {
	Exists() bool
	Deleted() bool
	Create(store *models.StoreSetting) error
	Update(store *models.StoreSetting) error
	Save() error
	Upsert() error
	Delete(id int64) error
	GetStoreSettingById(storeId int64) (*models.StoreSetting, error)
	// GetAllStores() (*models.StoreSetting, error)
}

type storeSettingRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.StoreSetting
}

func (m *storeSettingRepo) GetStoreSettingById(storeId int64) (*models.StoreSetting, error) {
	storeSetting, err := models.StoreSettingByID(m.ctx, m.db, storeId)
	if err != nil {
		err = &RepoError{
			Err:        err,
			StatusCode: GetStoreByIdErrorCode,
		}
	}
	return storeSetting, err
}

func (m *storeSettingRepo) Create(store *models.StoreSetting) error {
	if err := store.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *storeSettingRepo) Update(store *models.StoreSetting) error {
	if store.ID == 0 {
		return fmt.Errorf("error: can't update store, missing store ID")
	}
	if err := store.Update(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *storeSettingRepo) Save() error {
	return m.StoreSetting.Save(m.ctx, m.db)
}

func (m *storeSettingRepo) Upsert() error {
	return m.StoreSetting.Upsert(m.ctx, m.db)
}

func (m *storeSettingRepo) Delete(id int64) error {
	store, err := models.StoreByID(m.ctx, m.db, id)
	if err != nil {
		return err
	}
	return store.Delete(m.ctx, m.db)
}
