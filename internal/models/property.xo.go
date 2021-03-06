package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Property represents a row from 'public.property'.
type Property struct {
	ID          int64  `json:"id" db:"id"`                     // id
	StoreID     int64  `json:"store_id" db:"store_id"`         // store_id
	Name        string `json:"name" db:"name"`                 // name
	DisplayName string `json:"display_name" db:"display_name"` // display_name
	Fiterable   bool   `json:"fiterable" db:"fiterable"`       // fiterable
	FilterParam string `json:"filter_param" db:"filter_param"` // filter_param
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Property exists in the database.
func (p *Property) Exists() bool {
	return p._exists
}

// Deleted returns true when the Property has been marked for deletion from
// the database.
func (p *Property) Deleted() bool {
	return p._deleted
}

// Insert inserts the Property to the database.
func (p *Property) Insert(ctx context.Context, db DB) error {
	switch {
	case p._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case p._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.property (` +
		`store_id, name, display_name, fiterable, filter_param` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) RETURNING id`
	// run
	logf(sqlstr, p.StoreID, p.Name, p.DisplayName, p.Fiterable, p.FilterParam)
	if err := db.QueryRowContext(ctx, sqlstr, p.StoreID, p.Name, p.DisplayName, p.Fiterable, p.FilterParam).Scan(&p.ID); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Update updates a Property in the database.
func (p *Property) Update(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case p._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.property SET ` +
		`store_id = $1, name = $2, display_name = $3, fiterable = $4, filter_param = $5 ` +
		`WHERE id = $6`
	// run
	logf(sqlstr, p.StoreID, p.Name, p.DisplayName, p.Fiterable, p.FilterParam, p.ID)
	if _, err := db.ExecContext(ctx, sqlstr, p.StoreID, p.Name, p.DisplayName, p.Fiterable, p.FilterParam, p.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Property to the database.
func (p *Property) Save(ctx context.Context, db DB) error {
	if p.Exists() {
		return p.Update(ctx, db)
	}
	return p.Insert(ctx, db)
}

// Upsert performs an upsert for Property.
func (p *Property) Upsert(ctx context.Context, db DB) error {
	switch {
	case p._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.property (` +
		`id, store_id, name, display_name, fiterable, filter_param` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`store_id = EXCLUDED.store_id, name = EXCLUDED.name, display_name = EXCLUDED.display_name, fiterable = EXCLUDED.fiterable, filter_param = EXCLUDED.filter_param `
	// run
	logf(sqlstr, p.ID, p.StoreID, p.Name, p.DisplayName, p.Fiterable, p.FilterParam)
	if _, err := db.ExecContext(ctx, sqlstr, p.ID, p.StoreID, p.Name, p.DisplayName, p.Fiterable, p.FilterParam); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Delete deletes the Property from the database.
func (p *Property) Delete(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return nil
	case p._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.property ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, p.ID)
	if _, err := db.ExecContext(ctx, sqlstr, p.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	p._deleted = true
	return nil
}

// PropertyByName retrieves a row from 'public.property' as a Property.
//
// Generated from index 'property_name_key'.
func PropertyByName(ctx context.Context, db DB, name string) (*Property, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, store_id, name, display_name, fiterable, filter_param ` +
		`FROM public.property ` +
		`WHERE name = $1`
	// run
	logf(sqlstr, name)
	p := Property{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, name).Scan(&p.ID, &p.StoreID, &p.Name, &p.DisplayName, &p.Fiterable, &p.FilterParam); err != nil {
		return nil, logerror(err)
	}
	return &p, nil
}

// PropertyByID retrieves a row from 'public.property' as a Property.
//
// Generated from index 'property_pkey'.
func PropertyByID(ctx context.Context, db DB, id int64) (*Property, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, store_id, name, display_name, fiterable, filter_param ` +
		`FROM public.property ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	p := Property{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&p.ID, &p.StoreID, &p.Name, &p.DisplayName, &p.Fiterable, &p.FilterParam); err != nil {
		return nil, logerror(err)
	}
	return &p, nil
}
