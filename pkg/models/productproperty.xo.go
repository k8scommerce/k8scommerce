package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// ProductProperty represents a row from 'public.product_property'.
type ProductProperty struct {
	ProductID  int64 `json:"product_id" db:"product_id"`   // product_id
	PropertyID int64 `json:"property_id" db:"property_id"` // property_id
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the ProductProperty exists in the database.
func (pp *ProductProperty) Exists() bool {
	return pp._exists
}

// Deleted returns true when the ProductProperty has been marked for deletion from
// the database.
func (pp *ProductProperty) Deleted() bool {
	return pp._deleted
}

// Insert inserts the ProductProperty to the database.
func (pp *ProductProperty) Insert(ctx context.Context, db DB) error {
	switch {
	case pp._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case pp._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.product_property (` +
		`product_id, property_id` +
		`) VALUES (` +
		`$1, $2` +
		`)`
	// run
	logf(sqlstr, pp.ProductID, pp.PropertyID)
	if _, err := db.ExecContext(ctx, sqlstr, pp.ProductID, pp.PropertyID); err != nil {
		return logerror(err)
	}
	// set exists
	pp._exists = true
	return nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the ProductProperty from the database.
func (pp *ProductProperty) Delete(ctx context.Context, db DB) error {
	switch {
	case !pp._exists: // doesn't exist
		return nil
	case pp._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	const sqlstr = `DELETE FROM public.product_property ` +
		`WHERE product_id = $1 AND property_id = $2`
	// run
	logf(sqlstr, pp.ProductID, pp.PropertyID)
	if _, err := db.ExecContext(ctx, sqlstr, pp.ProductID, pp.PropertyID); err != nil {
		return logerror(err)
	}
	// set deleted
	pp._deleted = true
	return nil
}

// ProductPropertyByProductIDPropertyID retrieves a row from 'public.product_property' as a ProductProperty.
//
// Generated from index 'product_property_pkey'.
func ProductPropertyByProductIDPropertyID(ctx context.Context, db DB, productID, propertyID int64) (*ProductProperty, error) {
	// query
	const sqlstr = `SELECT ` +
		`product_id, property_id ` +
		`FROM public.product_property ` +
		`WHERE product_id = $1 AND property_id = $2`
	// run
	logf(sqlstr, productID, propertyID)
	pp := ProductProperty{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, productID, propertyID).Scan(&pp.ProductID, &pp.PropertyID); err != nil {
		return nil, logerror(err)
	}
	return &pp, nil
}

// Product returns the Product associated with the ProductProperty's (ProductID).
//
// Generated from foreign key 'product_property_product_id_fkey'.
func (pp *ProductProperty) Product(ctx context.Context, db DB) (*Product, error) {
	return ProductByID(ctx, db, pp.ProductID)
}

// Property returns the Property associated with the ProductProperty's (PropertyID).
//
// Generated from foreign key 'product_property_property_id_fkey'.
func (pp *ProductProperty) Property(ctx context.Context, db DB) (*Property, error) {
	return PropertyByID(ctx, db, pp.PropertyID)
}
