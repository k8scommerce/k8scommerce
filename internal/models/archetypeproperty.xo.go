package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// ArchetypeProperty represents a row from 'public.archetype_property'.
type ArchetypeProperty struct {
	ArchetypeID int64 `json:"archetype_id" db:"archetype_id"` // archetype_id
	PropertyID  int64 `json:"property_id" db:"property_id"`   // property_id
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the ArchetypeProperty exists in the database.
func (ap *ArchetypeProperty) Exists() bool {
	return ap._exists
}

// Deleted returns true when the ArchetypeProperty has been marked for deletion from
// the database.
func (ap *ArchetypeProperty) Deleted() bool {
	return ap._deleted
}

// Insert inserts the ArchetypeProperty to the database.
func (ap *ArchetypeProperty) Insert(ctx context.Context, db DB) error {
	switch {
	case ap._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ap._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.archetype_property (` +
		`archetype_id, property_id` +
		`) VALUES (` +
		`$1, $2` +
		`)`
	// run
	logf(sqlstr, ap.ArchetypeID, ap.PropertyID)
	if _, err := db.ExecContext(ctx, sqlstr, ap.ArchetypeID, ap.PropertyID); err != nil {
		return logerror(err)
	}
	// set exists
	ap._exists = true
	return nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the ArchetypeProperty from the database.
func (ap *ArchetypeProperty) Delete(ctx context.Context, db DB) error {
	switch {
	case !ap._exists: // doesn't exist
		return nil
	case ap._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	const sqlstr = `DELETE FROM public.archetype_property ` +
		`WHERE archetype_id = $1 AND property_id = $2`
	// run
	logf(sqlstr, ap.ArchetypeID, ap.PropertyID)
	if _, err := db.ExecContext(ctx, sqlstr, ap.ArchetypeID, ap.PropertyID); err != nil {
		return logerror(err)
	}
	// set deleted
	ap._deleted = true
	return nil
}

// ArchetypePropertyByArchetypeIDPropertyID retrieves a row from 'public.archetype_property' as a ArchetypeProperty.
//
// Generated from index 'archetype_property_pkey'.
func ArchetypePropertyByArchetypeIDPropertyID(ctx context.Context, db DB, archetypeID, propertyID int64) (*ArchetypeProperty, error) {
	// query
	const sqlstr = `SELECT ` +
		`archetype_id, property_id ` +
		`FROM public.archetype_property ` +
		`WHERE archetype_id = $1 AND property_id = $2`
	// run
	logf(sqlstr, archetypeID, propertyID)
	ap := ArchetypeProperty{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, archetypeID, propertyID).Scan(&ap.ArchetypeID, &ap.PropertyID); err != nil {
		return nil, logerror(err)
	}
	return &ap, nil
}

// Archetype returns the Archetype associated with the ArchetypeProperty's (ArchetypeID).
//
// Generated from foreign key 'archetype_property_archetype_id_fkey'.
func (ap *ArchetypeProperty) Archetype(ctx context.Context, db DB) (*Archetype, error) {
	return ArchetypeByID(ctx, db, ap.ArchetypeID)
}

// Property returns the Property associated with the ArchetypeProperty's (PropertyID).
//
// Generated from foreign key 'archetype_property_property_id_fkey'.
func (ap *ArchetypeProperty) Property(ctx context.Context, db DB) (*Property, error) {
	return PropertyByID(ctx, db, ap.PropertyID)
}