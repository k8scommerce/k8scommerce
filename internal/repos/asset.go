package repos

import (
	"context"
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

type Asset interface {
	Exists() bool
	Deleted() bool
	Create(asset *models.Asset) error
	CreateTx(asset *models.Asset, tx *sqlx.Tx) error
	Update(asset *models.Asset) error
	Save() error
	Upsert() error
	Delete(id int64) error
	GetAssetById(id int64) (res *models.Asset, err error)
	GetAssetsByVariantId(customerId int64) (res []*models.Asset, err error)
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

func (m *assetRepo) CreateTx(asset *models.Asset, tx *sqlx.Tx) error {
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

func (m *assetRepo) GetAssetsByVariantId(variantId int64) (res []*models.Asset, err error) {
	nstmt, err := m.db.PrepareNamed(`
		SELECT 
			*
		FROM asset
		WHERE variant_id = :variant_id
	`)
	if err != nil {
		return nil, fmt.Errorf("error::GetAssetsByVariantId::%s", err.Error())
	}

	var assets []*models.Asset
	err = nstmt.Select(&assets,
		map[string]interface{}{
			"variant_id": variantId,
		})
	if err != nil {
		return nil, err
	}

	return assets, nil
}
