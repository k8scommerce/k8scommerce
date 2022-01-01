package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Product represents a row from 'public.product'.
type Product struct {
	ID               int64          `json:"id" db:"id"`                               // id
	Slug             string         `json:"slug" db:"slug"`                           // slug
	Name             string         `json:"name" db:"name"`                           // name
	ShortDescription sql.NullString `json:"short_description" db:"short_description"` // short_description
	Description      sql.NullString `json:"description" db:"description"`             // description
	MetaTitle        sql.NullString `json:"meta_title" db:"meta_title"`               // meta_title
	MetaDescription  sql.NullString `json:"meta_description" db:"meta_description"`   // meta_description
	MetaKeywords     sql.NullString `json:"meta_keywords" db:"meta_keywords"`         // meta_keywords
	Promotionable    bool           `json:"promotionable" db:"promotionable"`         // promotionable
	Featured         bool           `json:"featured" db:"featured"`                   // featured
	AvailableOn      sql.NullTime   `json:"available_on" db:"available_on"`           // available_on
	DiscontinueOn    sql.NullTime   `json:"discontinue_on" db:"discontinue_on"`       // discontinue_on
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Product exists in the database.
func (p *Product) Exists() bool {
	return p._exists
}

// Deleted returns true when the Product has been marked for deletion from
// the database.
func (p *Product) Deleted() bool {
	return p._deleted
}

// Insert inserts the Product to the database.
func (p *Product) Insert(ctx context.Context, db DB) error {
	switch {
	case p._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case p._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.product (` +
		`slug, name, short_description, description, meta_title, meta_description, meta_keywords, promotionable, featured, available_on, discontinue_on` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) RETURNING id`
	// run
	logf(sqlstr, p.Slug, p.Name, p.ShortDescription, p.Description, p.MetaTitle, p.MetaDescription, p.MetaKeywords, p.Promotionable, p.Featured, p.AvailableOn, p.DiscontinueOn)
	if err := db.QueryRowContext(ctx, sqlstr, p.Slug, p.Name, p.ShortDescription, p.Description, p.MetaTitle, p.MetaDescription, p.MetaKeywords, p.Promotionable, p.Featured, p.AvailableOn, p.DiscontinueOn).Scan(&p.ID); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Update updates a Product in the database.
func (p *Product) Update(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case p._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.product SET ` +
		`slug = $1, name = $2, short_description = $3, description = $4, meta_title = $5, meta_description = $6, meta_keywords = $7, promotionable = $8, featured = $9, available_on = $10, discontinue_on = $11 ` +
		`WHERE id = $12`
	// run
	logf(sqlstr, p.Slug, p.Name, p.ShortDescription, p.Description, p.MetaTitle, p.MetaDescription, p.MetaKeywords, p.Promotionable, p.Featured, p.AvailableOn, p.DiscontinueOn, p.ID)
	if _, err := db.ExecContext(ctx, sqlstr, p.Slug, p.Name, p.ShortDescription, p.Description, p.MetaTitle, p.MetaDescription, p.MetaKeywords, p.Promotionable, p.Featured, p.AvailableOn, p.DiscontinueOn, p.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Product to the database.
func (p *Product) Save(ctx context.Context, db DB) error {
	if p.Exists() {
		return p.Update(ctx, db)
	}
	return p.Insert(ctx, db)
}

// Upsert performs an upsert for Product.
func (p *Product) Upsert(ctx context.Context, db DB) error {
	switch {
	case p._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.product (` +
		`id, slug, name, short_description, description, meta_title, meta_description, meta_keywords, promotionable, featured, available_on, discontinue_on` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`slug = EXCLUDED.slug, name = EXCLUDED.name, short_description = EXCLUDED.short_description, description = EXCLUDED.description, meta_title = EXCLUDED.meta_title, meta_description = EXCLUDED.meta_description, meta_keywords = EXCLUDED.meta_keywords, promotionable = EXCLUDED.promotionable, featured = EXCLUDED.featured, available_on = EXCLUDED.available_on, discontinue_on = EXCLUDED.discontinue_on `
	// run
	logf(sqlstr, p.ID, p.Slug, p.Name, p.ShortDescription, p.Description, p.MetaTitle, p.MetaDescription, p.MetaKeywords, p.Promotionable, p.Featured, p.AvailableOn, p.DiscontinueOn)
	if _, err := db.ExecContext(ctx, sqlstr, p.ID, p.Slug, p.Name, p.ShortDescription, p.Description, p.MetaTitle, p.MetaDescription, p.MetaKeywords, p.Promotionable, p.Featured, p.AvailableOn, p.DiscontinueOn); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Delete deletes the Product from the database.
func (p *Product) Delete(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return nil
	case p._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.product ` +
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

// ProductByID retrieves a row from 'public.product' as a Product.
//
// Generated from index 'product_pkey'.
func ProductByID(ctx context.Context, db DB, id int64) (*Product, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, slug, name, short_description, description, meta_title, meta_description, meta_keywords, promotionable, featured, available_on, discontinue_on ` +
		`FROM public.product ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	p := Product{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&p.ID, &p.Slug, &p.Name, &p.ShortDescription, &p.Description, &p.MetaTitle, &p.MetaDescription, &p.MetaKeywords, &p.Promotionable, &p.Featured, &p.AvailableOn, &p.DiscontinueOn); err != nil {
		return nil, logerror(err)
	}
	return &p, nil
}

// ProductBySlug retrieves a row from 'public.product' as a Product.
//
// Generated from index 'product_slug_key'.
func ProductBySlug(ctx context.Context, db DB, slug string) (*Product, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, slug, name, short_description, description, meta_title, meta_description, meta_keywords, promotionable, featured, available_on, discontinue_on ` +
		`FROM public.product ` +
		`WHERE slug = $1`
	// run
	logf(sqlstr, slug)
	p := Product{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, slug).Scan(&p.ID, &p.Slug, &p.Name, &p.ShortDescription, &p.Description, &p.MetaTitle, &p.MetaDescription, &p.MetaKeywords, &p.Promotionable, &p.Featured, &p.AvailableOn, &p.DiscontinueOn); err != nil {
		return nil, logerror(err)
	}
	return &p, nil
}
