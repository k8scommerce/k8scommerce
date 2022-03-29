package repos

import (
	"context"
	"fmt"

	"github.com/k8scommerce/k8scommerce/internal/models"

	"github.com/jmoiron/sqlx"
)

const (
	CustomerAddressUpdateError = "can't update customer address"
)

func newCustomerAddress(repo *repo) CustomerAddress {
	return &customerAddressRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

type CustomerAddress interface {
	Exists() bool
	Deleted() bool
	Create(address *models.CustomerAddress) error
	Update(address *models.CustomerAddress) error
	Save() error
	Upsert() error
	Delete(id int64) error
	GetCustomerAddressesByCustomerIdKind(customerId int64, kind models.AddressKind) ([]*models.CustomerAddress, error)
}

type customerAddressRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.CustomerAddress
}

func (m *customerAddressRepo) Create(address *models.CustomerAddress) error {
	if err := address.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *customerAddressRepo) Update(address *models.CustomerAddress) error {
	if address.ID == 0 {
		return &RepoError{
			Err:        fmt.Errorf(CustomerAddressUpdateError),
			StatusCode: UpdateErrorCode,
		}
	}
	if err := address.Update(m.ctx, m.db); err != nil {
		return &RepoError{
			Err:        err,
			StatusCode: UpdateErrorCode,
		}
	}
	return nil
}

func (m *customerAddressRepo) Save() error {
	if err := m.CustomerAddress.Save(m.ctx, m.db); err != nil {
		return &RepoError{
			Err:        err,
			StatusCode: SaveErrorCode,
		}
	}
	return nil
}

func (m *customerAddressRepo) Upsert() error {
	if err := m.CustomerAddress.Upsert(m.ctx, m.db); err != nil {
		return &RepoError{
			Err:        err,
			StatusCode: UpsertErrorCode,
		}
	}
	return nil
}

func (m *customerAddressRepo) Delete(id int64) error {
	address, err := models.CustomerByID(m.ctx, m.db, id)
	if err != nil {
		return &RepoError{
			Err:        err,
			StatusCode: DeleteErrorCode,
		}
	}
	if err := address.Delete(m.ctx, m.db); err != nil {
		return &RepoError{
			Err:        err,
			StatusCode: DeleteErrorCode,
		}
	}
	return nil
}

func (m *customerAddressRepo) GetCustomerAddressesByCustomerIdKind(customerId int64, kind models.AddressKind) ([]*models.CustomerAddress, error) {

	addresses, err := models.CustomerAddressByCustomerIDKind(m.ctx, m.db, customerId, kind)
	if err != nil {
		err = &RepoError{
			Err:        err,
			StatusCode: GetCustomerByEmailErrorCode,
		}
	}
	return addresses, err
}
