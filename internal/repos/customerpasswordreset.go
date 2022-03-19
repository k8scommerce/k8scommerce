package repos

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"k8scommerce/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
)

const (
	CustomerPasswordResetUpdateError = "can't update customer password reset, missing customer ID"
	GetCustomerByPasswordResetToken  = "could not find valid passsword reset token: %s"
)

func newCustomerPasswordReset(repo *repo) CustomerPasswordReset {
	return &customerPasswordResetRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

type CustomerPasswordReset interface {
	Exists() bool
	Deleted() bool
	Create(customerPasswordReset *models.CustomerPasswordReset) error
	Update(customerPasswordReset *models.CustomerPasswordReset) error
	Save() error
	Upsert() error
	Redeem(token string) error
	GetCustomerByPasswordResetToken(token string) (res *models.Customer, err error)
}

type customerPasswordResetRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.CustomerPasswordReset
}

func (m *customerPasswordResetRepo) GetCustomerByPasswordResetToken(token string) (res *models.Customer, err error) {
	customerPasswordReset, err := models.CustomerPasswordResetByToken(m.ctx, m.db, token)
	if err != nil {
		return nil, &RepoError{
			Err:        fmt.Errorf(GetCustomerByPasswordResetToken, token),
			StatusCode: SelectCustomerPasswordResetByTokenErrorCode,
		}
	}

	customer, err := customerPasswordReset.Customer(m.ctx, m.db)
	if err != nil {
		return nil, &RepoError{
			Err:        fmt.Errorf(GetCustomerByPasswordResetToken, token),
			StatusCode: SelectCustomerErrorCode,
		}
	}
	return customer, nil
}

func (m *customerPasswordResetRepo) Create(customerPasswordReset *models.CustomerPasswordReset) error {
	// hash the token
	customerPasswordReset.Token = xid.New().String()
	if err := customerPasswordReset.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *customerPasswordResetRepo) Update(customerPasswordReset *models.CustomerPasswordReset) error {
	if customerPasswordReset.ID == 0 {
		return &RepoError{
			Err:        fmt.Errorf(CustomerPasswordResetUpdateError),
			StatusCode: UpdateErrorCode,
		}
	}
	if err := customerPasswordReset.Update(m.ctx, m.db); err != nil {
		return &RepoError{
			Err:        err,
			StatusCode: UpdateErrorCode,
		}
	}
	return nil
}

func (m *customerPasswordResetRepo) Save() error {
	if err := m.CustomerPasswordReset.Save(m.ctx, m.db); err != nil {
		return &RepoError{
			Err:        err,
			StatusCode: SaveErrorCode,
		}
	}
	return nil
}

func (m *customerPasswordResetRepo) Upsert() error {
	if err := m.CustomerPasswordReset.Upsert(m.ctx, m.db); err != nil {
		return &RepoError{
			Err:        err,
			StatusCode: UpsertErrorCode,
		}
	}
	return nil
}

func (m *customerPasswordResetRepo) Redeem(token string) error {
	customerPasswordReset, err := models.CustomerPasswordResetByToken(m.ctx, m.db, token)
	if err != nil {
		return &RepoError{
			Err:        err,
			StatusCode: RedeemPasswordResetErrorCode,
		}
	}

	customerPasswordReset.RedeemedAt = sql.NullTime{Time: time.Now(), Valid: true}

	if err := customerPasswordReset.Update(m.ctx, m.db); err != nil {
		return &RepoError{
			Err:        err,
			StatusCode: RedeemPasswordResetErrorCode,
		}
	}
	return nil
}
