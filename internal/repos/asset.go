package repos

import (
	"context"
	"database/sql"
	"fmt"

	"k8scommerce/internal/models"

	"github.com/jmoiron/sqlx"
)

func newAsset(repo *repo) Asset {
	return &assetRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

//go:generate mockgen -destination=./mocks/Asset.go -package=mock_repos k8scommerce/internal/repos Asset
type Asset interface {
	Exists() bool
	Deleted() bool
	Create(asset *models.Asset) error
	CreateTx(asset *models.Asset, tx *sql.Tx) error
	Update(asset *models.Asset) error
	Save() error
	Upsert() error
	Delete(id int64) error
	GetAssetById(id int64) (res *models.Asset, err error)
	GetAssetsByVariantId(variantId int64, kind string) (res []*models.Asset, err error)
	AssetExists(storeId int64, name string) (res bool, err error)
}

type assetRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.Asset
}

// asset

func (m *assetRepo) Create(asset *models.Asset) error {
	if err := asset.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *assetRepo) CreateTx(asset *models.Asset, tx *sql.Tx) error {
	if err := asset.Insert(m.ctx, tx); err != nil {
		return err
	}
	return nil
}

func (m *assetRepo) Update(asset *models.Asset) error {
	if asset.ID == 0 {
		return fmt.Errorf("error: can't update asset, missing asset ID")
	}
	if err := asset.Update(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *assetRepo) Save() error {
	return m.Asset.Save(m.ctx, m.db)
}

func (m *assetRepo) Upsert() error {
	return m.Asset.Upsert(m.ctx, m.db)
}

func (m *assetRepo) Delete(id int64) error {
	asset, err := models.AssetByID(m.ctx, m.db, id)
	if err != nil {
		return err
	}
	return asset.Delete(m.ctx, m.db)
}

func (m *assetRepo) GetAssetById(id int64) (res *models.Asset, err error) {
	return models.AssetByID(m.ctx, m.db, id)
}

func (m *assetRepo) GetAssetsByVariantId(variantId int64, kind string) (res []*models.Asset, err error) {
	return models.AssetByVariantIDKind(m.ctx, m.db, variantId, kind)
}

func (m *assetRepo) AssetExists(storeId int64, name string) (res bool, err error) {
	nstmt, err := m.db.PrepareNamed(`
		SELECT 
			id
		FROM asset
		WHERE store_id = :store_id
		AND name = :name
	`)
	if err != nil {
		return false, fmt.Errorf("error::AssetExists::%s", err.Error())
	}

	var result []*struct {
		Id int64 `db:"id"`
	}
	err = nstmt.Select(&result,
		map[string]interface{}{
			"store_id": storeId,
			"name":     name,
		})
	if err != nil {
		return false, fmt.Errorf("error::AssetExists::Query::%s", err.Error())
	}

	if len(result) != 0 {
		return true, nil
	}
	return false, nil
}
