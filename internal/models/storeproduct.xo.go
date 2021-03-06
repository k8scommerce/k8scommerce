package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// StoreProduct represents a row from 'public.store_product'.
type StoreProduct struct {
	StoreID   int64 `json:"store_id" db:"store_id"`     // store_id
	ProductID int64 `json:"product_id" db:"product_id"` // product_id
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the StoreProduct exists in the database.
func (sp *StoreProduct) Exists() bool {
	return sp._exists
}

// Deleted returns true when the StoreProduct has been marked for deletion from
// the database.
func (sp *StoreProduct) Deleted() bool {
	return sp._deleted
}

// Insert inserts the StoreProduct to the database.
func (sp *StoreProduct) Insert(ctx context.Context, db DB) error {
	switch {
	case sp._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case sp._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.store_product (` +
		`store_id, product_id` +
		`) VALUES (` +
		`$1, $2` +
		`)`
	// run
	logf(sqlstr, sp.StoreID, sp.ProductID)
	if _, err := db.ExecContext(ctx, sqlstr, sp.StoreID, sp.ProductID); err != nil {
		return logerror(err)
	}
	// set exists
	sp._exists = true
	return nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the StoreProduct from the database.
func (sp *StoreProduct) Delete(ctx context.Context, db DB) error {
	switch {
	case !sp._exists: // doesn't exist
		return nil
	case sp._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	const sqlstr = `DELETE FROM public.store_product ` +
		`WHERE store_id = $1 AND product_id = $2`
	// run
	logf(sqlstr, sp.StoreID, sp.ProductID)
	if _, err := db.ExecContext(ctx, sqlstr, sp.StoreID, sp.ProductID); err != nil {
		return logerror(err)
	}
	// set deleted
	sp._deleted = true
	return nil
}

// StoreProductByStoreIDProductID retrieves a row from 'public.store_product' as a StoreProduct.
//
// Generated from index 'store_product_pkey'.
func StoreProductByStoreIDProductID(ctx context.Context, db DB, storeID, productID int64) (*StoreProduct, error) {
	// query
	const sqlstr = `SELECT ` +
		`store_id, product_id ` +
		`FROM public.store_product ` +
		`WHERE store_id = $1 AND product_id = $2`
	// run
	logf(sqlstr, storeID, productID)
	sp := StoreProduct{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, storeID, productID).Scan(&sp.StoreID, &sp.ProductID); err != nil {
		return nil, logerror(err)
	}
	return &sp, nil
}

// Product returns the Product associated with the StoreProduct's (ProductID).
//
// Generated from foreign key 'store_product_product_id_fkey'.
func (sp *StoreProduct) Product(ctx context.Context, db DB) (*Product, error) {
	return ProductByID(ctx, db, sp.ProductID)
}

// Store returns the Store associated with the StoreProduct's (StoreID).
//
// Generated from foreign key 'store_product_store_id_fkey'.
func (sp *StoreProduct) Store(ctx context.Context, db DB) (*Store, error) {
	return StoreByID(ctx, db, sp.StoreID)
}
