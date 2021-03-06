package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Archetype represents a row from 'public.archetype'.
type Archetype struct {
	ID      int64  `json:"id" db:"id"`             // id
	StoreID int64  `json:"store_id" db:"store_id"` // store_id
	Name    string `json:"name" db:"name"`         // name
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Archetype exists in the database.
func (a *Archetype) Exists() bool {
	return a._exists
}

// Deleted returns true when the Archetype has been marked for deletion from
// the database.
func (a *Archetype) Deleted() bool {
	return a._deleted
}

// Insert inserts the Archetype to the database.
func (a *Archetype) Insert(ctx context.Context, db DB) error {
	switch {
	case a._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case a._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.archetype (` +
		`store_id, name` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING id`
	// run
	logf(sqlstr, a.StoreID, a.Name)
	if err := db.QueryRowContext(ctx, sqlstr, a.StoreID, a.Name).Scan(&a.ID); err != nil {
		return logerror(err)
	}
	// set exists
	a._exists = true
	return nil
}

// Update updates a Archetype in the database.
func (a *Archetype) Update(ctx context.Context, db DB) error {
	switch {
	case !a._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case a._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.archetype SET ` +
		`store_id = $1, name = $2 ` +
		`WHERE id = $3`
	// run
	logf(sqlstr, a.StoreID, a.Name, a.ID)
	if _, err := db.ExecContext(ctx, sqlstr, a.StoreID, a.Name, a.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Archetype to the database.
func (a *Archetype) Save(ctx context.Context, db DB) error {
	if a.Exists() {
		return a.Update(ctx, db)
	}
	return a.Insert(ctx, db)
}

// Upsert performs an upsert for Archetype.
func (a *Archetype) Upsert(ctx context.Context, db DB) error {
	switch {
	case a._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.archetype (` +
		`id, store_id, name` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`store_id = EXCLUDED.store_id, name = EXCLUDED.name `
	// run
	logf(sqlstr, a.ID, a.StoreID, a.Name)
	if _, err := db.ExecContext(ctx, sqlstr, a.ID, a.StoreID, a.Name); err != nil {
		return logerror(err)
	}
	// set exists
	a._exists = true
	return nil
}

// Delete deletes the Archetype from the database.
func (a *Archetype) Delete(ctx context.Context, db DB) error {
	switch {
	case !a._exists: // doesn't exist
		return nil
	case a._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.archetype ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, a.ID)
	if _, err := db.ExecContext(ctx, sqlstr, a.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	a._deleted = true
	return nil
}

// ArchetypeByName retrieves a row from 'public.archetype' as a Archetype.
//
// Generated from index 'archetype_name_key'.
func ArchetypeByName(ctx context.Context, db DB, name string) (*Archetype, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, store_id, name ` +
		`FROM public.archetype ` +
		`WHERE name = $1`
	// run
	logf(sqlstr, name)
	a := Archetype{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, name).Scan(&a.ID, &a.StoreID, &a.Name); err != nil {
		return nil, logerror(err)
	}
	return &a, nil
}

// ArchetypeByID retrieves a row from 'public.archetype' as a Archetype.
//
// Generated from index 'archetype_pkey'.
func ArchetypeByID(ctx context.Context, db DB, id int64) (*Archetype, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, store_id, name ` +
		`FROM public.archetype ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	a := Archetype{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&a.ID, &a.StoreID, &a.Name); err != nil {
		return nil, logerror(err)
	}
	return &a, nil
}
