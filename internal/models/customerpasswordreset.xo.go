package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// CustomerPasswordReset represents a row from 'public.customer_password_reset'.
type CustomerPasswordReset struct {
	ID         int64        `json:"id" db:"id"`                   // id
	StoreID    int64        `json:"store_id" db:"store_id"`       // store_id
	CustomerID int64        `json:"customer_id" db:"customer_id"` // customer_id
	Token      string       `json:"token" db:"token"`             // token
	RedeemedAt sql.NullTime `json:"redeemed_at" db:"redeemed_at"` // redeemed_at
	ExpiredAt  sql.NullTime `json:"expired_at" db:"expired_at"`   // expired_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the CustomerPasswordReset exists in the database.
func (cpr *CustomerPasswordReset) Exists() bool {
	return cpr._exists
}

// Deleted returns true when the CustomerPasswordReset has been marked for deletion from
// the database.
func (cpr *CustomerPasswordReset) Deleted() bool {
	return cpr._deleted
}

// Insert inserts the CustomerPasswordReset to the database.
func (cpr *CustomerPasswordReset) Insert(ctx context.Context, db DB) error {
	switch {
	case cpr._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case cpr._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.customer_password_reset (` +
		`store_id, customer_id, token, redeemed_at, expired_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) RETURNING id`
	// run
	logf(sqlstr, cpr.StoreID, cpr.CustomerID, cpr.Token, cpr.RedeemedAt, cpr.ExpiredAt)
	if err := db.QueryRowContext(ctx, sqlstr, cpr.StoreID, cpr.CustomerID, cpr.Token, cpr.RedeemedAt, cpr.ExpiredAt).Scan(&cpr.ID); err != nil {
		return logerror(err)
	}
	// set exists
	cpr._exists = true
	return nil
}

// Update updates a CustomerPasswordReset in the database.
func (cpr *CustomerPasswordReset) Update(ctx context.Context, db DB) error {
	switch {
	case !cpr._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case cpr._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.customer_password_reset SET ` +
		`store_id = $1, customer_id = $2, token = $3, redeemed_at = $4, expired_at = $5 ` +
		`WHERE id = $6`
	// run
	logf(sqlstr, cpr.StoreID, cpr.CustomerID, cpr.Token, cpr.RedeemedAt, cpr.ExpiredAt, cpr.ID)
	if _, err := db.ExecContext(ctx, sqlstr, cpr.StoreID, cpr.CustomerID, cpr.Token, cpr.RedeemedAt, cpr.ExpiredAt, cpr.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the CustomerPasswordReset to the database.
func (cpr *CustomerPasswordReset) Save(ctx context.Context, db DB) error {
	if cpr.Exists() {
		return cpr.Update(ctx, db)
	}
	return cpr.Insert(ctx, db)
}

// Upsert performs an upsert for CustomerPasswordReset.
func (cpr *CustomerPasswordReset) Upsert(ctx context.Context, db DB) error {
	switch {
	case cpr._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.customer_password_reset (` +
		`id, store_id, customer_id, token, redeemed_at, expired_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`store_id = EXCLUDED.store_id, customer_id = EXCLUDED.customer_id, token = EXCLUDED.token, redeemed_at = EXCLUDED.redeemed_at, expired_at = EXCLUDED.expired_at `
	// run
	logf(sqlstr, cpr.ID, cpr.StoreID, cpr.CustomerID, cpr.Token, cpr.RedeemedAt, cpr.ExpiredAt)
	if _, err := db.ExecContext(ctx, sqlstr, cpr.ID, cpr.StoreID, cpr.CustomerID, cpr.Token, cpr.RedeemedAt, cpr.ExpiredAt); err != nil {
		return logerror(err)
	}
	// set exists
	cpr._exists = true
	return nil
}

// Delete deletes the CustomerPasswordReset from the database.
func (cpr *CustomerPasswordReset) Delete(ctx context.Context, db DB) error {
	switch {
	case !cpr._exists: // doesn't exist
		return nil
	case cpr._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.customer_password_reset ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, cpr.ID)
	if _, err := db.ExecContext(ctx, sqlstr, cpr.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	cpr._deleted = true
	return nil
}

// CustomerPasswordResetByID retrieves a row from 'public.customer_password_reset' as a CustomerPasswordReset.
//
// Generated from index 'customer_password_reset_pkey'.
func CustomerPasswordResetByID(ctx context.Context, db DB, id int64) (*CustomerPasswordReset, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, store_id, customer_id, token, redeemed_at, expired_at ` +
		`FROM public.customer_password_reset ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	cpr := CustomerPasswordReset{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&cpr.ID, &cpr.StoreID, &cpr.CustomerID, &cpr.Token, &cpr.RedeemedAt, &cpr.ExpiredAt); err != nil {
		return nil, logerror(err)
	}
	return &cpr, nil
}

// CustomerPasswordResetByToken retrieves a row from 'public.customer_password_reset' as a CustomerPasswordReset.
//
// Generated from index 'customer_password_reset_token_key'.
func CustomerPasswordResetByToken(ctx context.Context, db DB, token string) (*CustomerPasswordReset, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, store_id, customer_id, token, redeemed_at, expired_at ` +
		`FROM public.customer_password_reset ` +
		`WHERE token = $1`
	// run
	logf(sqlstr, token)
	cpr := CustomerPasswordReset{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, token).Scan(&cpr.ID, &cpr.StoreID, &cpr.CustomerID, &cpr.Token, &cpr.RedeemedAt, &cpr.ExpiredAt); err != nil {
		return nil, logerror(err)
	}
	return &cpr, nil
}

// CustomerPasswordResetByStoreID retrieves a row from 'public.customer_password_reset' as a CustomerPasswordReset.
//
// Generated from index 'idx_customer_password_reset_store_id'.
func CustomerPasswordResetByStoreID(ctx context.Context, db DB, storeID int64) ([]*CustomerPasswordReset, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, store_id, customer_id, token, redeemed_at, expired_at ` +
		`FROM public.customer_password_reset ` +
		`WHERE store_id = $1`
	// run
	logf(sqlstr, storeID)
	rows, err := db.QueryContext(ctx, sqlstr, storeID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*CustomerPasswordReset
	for rows.Next() {
		cpr := CustomerPasswordReset{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&cpr.ID, &cpr.StoreID, &cpr.CustomerID, &cpr.Token, &cpr.RedeemedAt, &cpr.ExpiredAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &cpr)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Customer returns the Customer associated with the CustomerPasswordReset's (CustomerID).
//
// Generated from foreign key 'customer_password_reset_customer_id_fkey'.
func (cpr *CustomerPasswordReset) Customer(ctx context.Context, db DB) (*Customer, error) {
	return CustomerByID(ctx, db, cpr.CustomerID)
}
