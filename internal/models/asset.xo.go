package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Asset represents a row from 'public.asset'.
type Asset struct {
	ID          int64          `json:"id" db:"id"`                     // id
	StoreID     int64          `json:"store_id" db:"store_id"`         // store_id
	ProductID   int64          `json:"product_id" db:"product_id"`     // product_id
	VariantID   int64          `json:"variant_id" db:"variant_id"`     // variant_id
	Name        string         `json:"name" db:"name"`                 // name
	URL         string         `json:"url" db:"url"`                   // url
	DisplayName sql.NullString `json:"display_name" db:"display_name"` // display_name
	Kind        string         `json:"kind" db:"kind"`                 // kind
	ContentType string         `json:"content_type" db:"content_type"` // content_type
	SortOrder   sql.NullInt64  `json:"sort_order" db:"sort_order"`     // sort_order
	Sizes       []byte         `json:"sizes" db:"sizes"`               // sizes
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Asset exists in the database.
func (a *Asset) Exists() bool {
	return a._exists
}

// Deleted returns true when the Asset has been marked for deletion from
// the database.
func (a *Asset) Deleted() bool {
	return a._deleted
}

// Insert inserts the Asset to the database.
func (a *Asset) Insert(ctx context.Context, db DB) error {
	switch {
	case a._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case a._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.asset (` +
		`store_id, product_id, variant_id, name, url, display_name, kind, content_type, sort_order, sizes` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10` +
		`) RETURNING id`
	// run
	logf(sqlstr, a.StoreID, a.ProductID, a.VariantID, a.Name, a.URL, a.DisplayName, a.Kind, a.ContentType, a.SortOrder, a.Sizes)
	if err := db.QueryRowContext(ctx, sqlstr, a.StoreID, a.ProductID, a.VariantID, a.Name, a.URL, a.DisplayName, a.Kind, a.ContentType, a.SortOrder, a.Sizes).Scan(&a.ID); err != nil {
		return logerror(err)
	}
	// set exists
	a._exists = true
	return nil
}

// Update updates a Asset in the database.
func (a *Asset) Update(ctx context.Context, db DB) error {
	switch {
	case !a._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case a._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.asset SET ` +
		`store_id = $1, product_id = $2, variant_id = $3, name = $4, url = $5, display_name = $6, kind = $7, content_type = $8, sort_order = $9, sizes = $10 ` +
		`WHERE id = $11`
	// run
	logf(sqlstr, a.StoreID, a.ProductID, a.VariantID, a.Name, a.URL, a.DisplayName, a.Kind, a.ContentType, a.SortOrder, a.Sizes, a.ID)
	if _, err := db.ExecContext(ctx, sqlstr, a.StoreID, a.ProductID, a.VariantID, a.Name, a.URL, a.DisplayName, a.Kind, a.ContentType, a.SortOrder, a.Sizes, a.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Asset to the database.
func (a *Asset) Save(ctx context.Context, db DB) error {
	if a.Exists() {
		return a.Update(ctx, db)
	}
	return a.Insert(ctx, db)
}

// Upsert performs an upsert for Asset.
func (a *Asset) Upsert(ctx context.Context, db DB) error {
	switch {
	case a._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.asset (` +
		`id, store_id, product_id, variant_id, name, url, display_name, kind, content_type, sort_order, sizes` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`store_id = EXCLUDED.store_id, product_id = EXCLUDED.product_id, variant_id = EXCLUDED.variant_id, name = EXCLUDED.name, url = EXCLUDED.url, display_name = EXCLUDED.display_name, kind = EXCLUDED.kind, content_type = EXCLUDED.content_type, sort_order = EXCLUDED.sort_order, sizes = EXCLUDED.sizes `
	// run
	logf(sqlstr, a.ID, a.StoreID, a.ProductID, a.VariantID, a.Name, a.URL, a.DisplayName, a.Kind, a.ContentType, a.SortOrder, a.Sizes)
	if _, err := db.ExecContext(ctx, sqlstr, a.ID, a.StoreID, a.ProductID, a.VariantID, a.Name, a.URL, a.DisplayName, a.Kind, a.ContentType, a.SortOrder, a.Sizes); err != nil {
		return logerror(err)
	}
	// set exists
	a._exists = true
	return nil
}

// Delete deletes the Asset from the database.
func (a *Asset) Delete(ctx context.Context, db DB) error {
	switch {
	case !a._exists: // doesn't exist
		return nil
	case a._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.asset ` +
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

// AssetByID retrieves a row from 'public.asset' as a Asset.
//
// Generated from index 'asset_pkey'.
func AssetByID(ctx context.Context, db DB, id int64) (*Asset, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, store_id, product_id, variant_id, name, url, display_name, kind, content_type, sort_order, sizes ` +
		`FROM public.asset ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	a := Asset{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&a.ID, &a.StoreID, &a.ProductID, &a.VariantID, &a.Name, &a.URL, &a.DisplayName, &a.Kind, &a.ContentType, &a.SortOrder, &a.Sizes); err != nil {
		return nil, logerror(err)
	}
	return &a, nil
}

// AssetByStoreIDName retrieves a row from 'public.asset' as a Asset.
//
// Generated from index 'asset_store_id_name_key'.
func AssetByStoreIDName(ctx context.Context, db DB, storeID int64, name string) (*Asset, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, store_id, product_id, variant_id, name, url, display_name, kind, content_type, sort_order, sizes ` +
		`FROM public.asset ` +
		`WHERE store_id = $1 AND name = $2`
	// run
	logf(sqlstr, storeID, name)
	a := Asset{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, storeID, name).Scan(&a.ID, &a.StoreID, &a.ProductID, &a.VariantID, &a.Name, &a.URL, &a.DisplayName, &a.Kind, &a.ContentType, &a.SortOrder, &a.Sizes); err != nil {
		return nil, logerror(err)
	}
	return &a, nil
}

// AssetByProductID retrieves a row from 'public.asset' as a Asset.
//
// Generated from index 'idx_asset_product_id'.
func AssetByProductID(ctx context.Context, db DB, productID int64) ([]*Asset, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, store_id, product_id, variant_id, name, url, display_name, kind, content_type, sort_order, sizes ` +
		`FROM public.asset ` +
		`WHERE product_id = $1`
	// run
	logf(sqlstr, productID)
	rows, err := db.QueryContext(ctx, sqlstr, productID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Asset
	for rows.Next() {
		a := Asset{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&a.ID, &a.StoreID, &a.ProductID, &a.VariantID, &a.Name, &a.URL, &a.DisplayName, &a.Kind, &a.ContentType, &a.SortOrder, &a.Sizes); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &a)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// AssetBySizes retrieves a row from 'public.asset' as a Asset.
//
// Generated from index 'idx_asset_sizes'.
func AssetBySizes(ctx context.Context, db DB, sizes []byte) ([]*Asset, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, store_id, product_id, variant_id, name, url, display_name, kind, content_type, sort_order, sizes ` +
		`FROM public.asset ` +
		`WHERE sizes = $1`
	// run
	logf(sqlstr, sizes)
	rows, err := db.QueryContext(ctx, sqlstr, sizes)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Asset
	for rows.Next() {
		a := Asset{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&a.ID, &a.StoreID, &a.ProductID, &a.VariantID, &a.Name, &a.URL, &a.DisplayName, &a.Kind, &a.ContentType, &a.SortOrder, &a.Sizes); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &a)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// AssetByVariantID retrieves a row from 'public.asset' as a Asset.
//
// Generated from index 'idx_asset_variant_id'.
func AssetByVariantID(ctx context.Context, db DB, variantID int64) ([]*Asset, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, store_id, product_id, variant_id, name, url, display_name, kind, content_type, sort_order, sizes ` +
		`FROM public.asset ` +
		`WHERE variant_id = $1`
	// run
	logf(sqlstr, variantID)
	rows, err := db.QueryContext(ctx, sqlstr, variantID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Asset
	for rows.Next() {
		a := Asset{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&a.ID, &a.StoreID, &a.ProductID, &a.VariantID, &a.Name, &a.URL, &a.DisplayName, &a.Kind, &a.ContentType, &a.SortOrder, &a.Sizes); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &a)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// AssetByVariantIDKind retrieves a row from 'public.asset' as a Asset.
//
// Generated from index 'idx_asset_variant_id_kind'.
func AssetByVariantIDKind(ctx context.Context, db DB, variantID int64, kind string) ([]*Asset, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, store_id, product_id, variant_id, name, url, display_name, kind, content_type, sort_order, sizes ` +
		`FROM public.asset ` +
		`WHERE variant_id = $1 AND kind = $2`
	// run
	logf(sqlstr, variantID, kind)
	rows, err := db.QueryContext(ctx, sqlstr, variantID, kind)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Asset
	for rows.Next() {
		a := Asset{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&a.ID, &a.StoreID, &a.ProductID, &a.VariantID, &a.Name, &a.URL, &a.DisplayName, &a.Kind, &a.ContentType, &a.SortOrder, &a.Sizes); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &a)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
