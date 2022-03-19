package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// AdminPasswordReset represents a row from 'public.admin_password_reset'.
type AdminPasswordReset struct {
	ID         int64        `json:"id" db:"id"`                   // id
	AdminID    int64        `json:"admin_id" db:"admin_id"`       // admin_id
	Token      string       `json:"token" db:"token"`             // token
	RedeemedAt sql.NullTime `json:"redeemed_at" db:"redeemed_at"` // redeemed_at
	ExpiredAt  sql.NullTime `json:"expired_at" db:"expired_at"`   // expired_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the AdminPasswordReset exists in the database.
func (apr *AdminPasswordReset) Exists() bool {
	return apr._exists
}

// Deleted returns true when the AdminPasswordReset has been marked for deletion from
// the database.
func (apr *AdminPasswordReset) Deleted() bool {
	return apr._deleted
}

// Insert inserts the AdminPasswordReset to the database.
func (apr *AdminPasswordReset) Insert(ctx context.Context, db DB) error {
	switch {
	case apr._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case apr._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.admin_password_reset (` +
		`admin_id, token, redeemed_at, expired_at` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING id`
	// run
	logf(sqlstr, apr.AdminID, apr.Token, apr.RedeemedAt, apr.ExpiredAt)
	if err := db.QueryRowContext(ctx, sqlstr, apr.AdminID, apr.Token, apr.RedeemedAt, apr.ExpiredAt).Scan(&apr.ID); err != nil {
		return logerror(err)
	}
	// set exists
	apr._exists = true
	return nil
}

// Update updates a AdminPasswordReset in the database.
func (apr *AdminPasswordReset) Update(ctx context.Context, db DB) error {
	switch {
	case !apr._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case apr._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.admin_password_reset SET ` +
		`admin_id = $1, token = $2, redeemed_at = $3, expired_at = $4 ` +
		`WHERE id = $5`
	// run
	logf(sqlstr, apr.AdminID, apr.Token, apr.RedeemedAt, apr.ExpiredAt, apr.ID)
	if _, err := db.ExecContext(ctx, sqlstr, apr.AdminID, apr.Token, apr.RedeemedAt, apr.ExpiredAt, apr.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the AdminPasswordReset to the database.
func (apr *AdminPasswordReset) Save(ctx context.Context, db DB) error {
	if apr.Exists() {
		return apr.Update(ctx, db)
	}
	return apr.Insert(ctx, db)
}

// Upsert performs an upsert for AdminPasswordReset.
func (apr *AdminPasswordReset) Upsert(ctx context.Context, db DB) error {
	switch {
	case apr._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.admin_password_reset (` +
		`id, admin_id, token, redeemed_at, expired_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`admin_id = EXCLUDED.admin_id, token = EXCLUDED.token, redeemed_at = EXCLUDED.redeemed_at, expired_at = EXCLUDED.expired_at `
	// run
	logf(sqlstr, apr.ID, apr.AdminID, apr.Token, apr.RedeemedAt, apr.ExpiredAt)
	if _, err := db.ExecContext(ctx, sqlstr, apr.ID, apr.AdminID, apr.Token, apr.RedeemedAt, apr.ExpiredAt); err != nil {
		return logerror(err)
	}
	// set exists
	apr._exists = true
	return nil
}

// Delete deletes the AdminPasswordReset from the database.
func (apr *AdminPasswordReset) Delete(ctx context.Context, db DB) error {
	switch {
	case !apr._exists: // doesn't exist
		return nil
	case apr._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.admin_password_reset ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, apr.ID)
	if _, err := db.ExecContext(ctx, sqlstr, apr.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	apr._deleted = true
	return nil
}

// AdminPasswordResetByID retrieves a row from 'public.admin_password_reset' as a AdminPasswordReset.
//
// Generated from index 'admin_password_reset_pkey'.
func AdminPasswordResetByID(ctx context.Context, db DB, id int64) (*AdminPasswordReset, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, admin_id, token, redeemed_at, expired_at ` +
		`FROM public.admin_password_reset ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	apr := AdminPasswordReset{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&apr.ID, &apr.AdminID, &apr.Token, &apr.RedeemedAt, &apr.ExpiredAt); err != nil {
		return nil, logerror(err)
	}
	return &apr, nil
}

// AdminPasswordResetByToken retrieves a row from 'public.admin_password_reset' as a AdminPasswordReset.
//
// Generated from index 'admin_password_reset_token_key'.
func AdminPasswordResetByToken(ctx context.Context, db DB, token string) (*AdminPasswordReset, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, admin_id, token, redeemed_at, expired_at ` +
		`FROM public.admin_password_reset ` +
		`WHERE token = $1`
	// run
	logf(sqlstr, token)
	apr := AdminPasswordReset{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, token).Scan(&apr.ID, &apr.AdminID, &apr.Token, &apr.RedeemedAt, &apr.ExpiredAt); err != nil {
		return nil, logerror(err)
	}
	return &apr, nil
}
