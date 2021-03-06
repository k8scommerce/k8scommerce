package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Category represents a row from 'public.category'.
type Category struct {
	ID              int64          `json:"id" db:"id"`                             // id
	ParentID        sql.NullInt64  `json:"parent_id" db:"parent_id"`               // parent_id
	StoreID         int64          `json:"store_id" db:"store_id"`                 // store_id
	Slug            string         `json:"slug" db:"slug"`                         // slug
	Name            string         `json:"name" db:"name"`                         // name
	Description     sql.NullString `json:"description" db:"description"`           // description
	MetaTitle       sql.NullString `json:"meta_title" db:"meta_title"`             // meta_title
	MetaDescription sql.NullString `json:"meta_description" db:"meta_description"` // meta_description
	MetaKeywords    sql.NullString `json:"meta_keywords" db:"meta_keywords"`       // meta_keywords
	HideFromNav     sql.NullBool   `json:"hide_from_nav" db:"hide_from_nav"`       // hide_from_nav
	Lft             sql.NullInt64  `json:"lft" db:"lft"`                           // lft
	Rgt             sql.NullInt64  `json:"rgt" db:"rgt"`                           // rgt
	Depth           sql.NullInt64  `json:"depth" db:"depth"`                       // depth
	SortOrder       sql.NullInt64  `json:"sort_order" db:"sort_order"`             // sort_order
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Category exists in the database.
func (c *Category) Exists() bool {
	return c._exists
}

// Deleted returns true when the Category has been marked for deletion from
// the database.
func (c *Category) Deleted() bool {
	return c._deleted
}

// Insert inserts the Category to the database.
func (c *Category) Insert(ctx context.Context, db DB) error {
	switch {
	case c._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case c._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.category (` +
		`parent_id, store_id, slug, name, description, meta_title, meta_description, meta_keywords, hide_from_nav, lft, rgt, depth, sort_order` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`) RETURNING id`
	// run
	logf(sqlstr, c.ParentID, c.StoreID, c.Slug, c.Name, c.Description, c.MetaTitle, c.MetaDescription, c.MetaKeywords, c.HideFromNav, c.Lft, c.Rgt, c.Depth, c.SortOrder)
	if err := db.QueryRowContext(ctx, sqlstr, c.ParentID, c.StoreID, c.Slug, c.Name, c.Description, c.MetaTitle, c.MetaDescription, c.MetaKeywords, c.HideFromNav, c.Lft, c.Rgt, c.Depth, c.SortOrder).Scan(&c.ID); err != nil {
		return logerror(err)
	}
	// set exists
	c._exists = true
	return nil
}

// Update updates a Category in the database.
func (c *Category) Update(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case c._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.category SET ` +
		`parent_id = $1, store_id = $2, slug = $3, name = $4, description = $5, meta_title = $6, meta_description = $7, meta_keywords = $8, hide_from_nav = $9, lft = $10, rgt = $11, depth = $12, sort_order = $13 ` +
		`WHERE id = $14`
	// run
	logf(sqlstr, c.ParentID, c.StoreID, c.Slug, c.Name, c.Description, c.MetaTitle, c.MetaDescription, c.MetaKeywords, c.HideFromNav, c.Lft, c.Rgt, c.Depth, c.SortOrder, c.ID)
	if _, err := db.ExecContext(ctx, sqlstr, c.ParentID, c.StoreID, c.Slug, c.Name, c.Description, c.MetaTitle, c.MetaDescription, c.MetaKeywords, c.HideFromNav, c.Lft, c.Rgt, c.Depth, c.SortOrder, c.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Category to the database.
func (c *Category) Save(ctx context.Context, db DB) error {
	if c.Exists() {
		return c.Update(ctx, db)
	}
	return c.Insert(ctx, db)
}

// Upsert performs an upsert for Category.
func (c *Category) Upsert(ctx context.Context, db DB) error {
	switch {
	case c._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.category (` +
		`id, parent_id, store_id, slug, name, description, meta_title, meta_description, meta_keywords, hide_from_nav, lft, rgt, depth, sort_order` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`parent_id = EXCLUDED.parent_id, store_id = EXCLUDED.store_id, slug = EXCLUDED.slug, name = EXCLUDED.name, description = EXCLUDED.description, meta_title = EXCLUDED.meta_title, meta_description = EXCLUDED.meta_description, meta_keywords = EXCLUDED.meta_keywords, hide_from_nav = EXCLUDED.hide_from_nav, lft = EXCLUDED.lft, rgt = EXCLUDED.rgt, depth = EXCLUDED.depth, sort_order = EXCLUDED.sort_order `
	// run
	logf(sqlstr, c.ID, c.ParentID, c.StoreID, c.Slug, c.Name, c.Description, c.MetaTitle, c.MetaDescription, c.MetaKeywords, c.HideFromNav, c.Lft, c.Rgt, c.Depth, c.SortOrder)
	if _, err := db.ExecContext(ctx, sqlstr, c.ID, c.ParentID, c.StoreID, c.Slug, c.Name, c.Description, c.MetaTitle, c.MetaDescription, c.MetaKeywords, c.HideFromNav, c.Lft, c.Rgt, c.Depth, c.SortOrder); err != nil {
		return logerror(err)
	}
	// set exists
	c._exists = true
	return nil
}

// Delete deletes the Category from the database.
func (c *Category) Delete(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return nil
	case c._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.category ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, c.ID)
	if _, err := db.ExecContext(ctx, sqlstr, c.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	c._deleted = true
	return nil
}

// CategoryByID retrieves a row from 'public.category' as a Category.
//
// Generated from index 'category_pkey'.
func CategoryByID(ctx context.Context, db DB, id int64) (*Category, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, parent_id, store_id, slug, name, description, meta_title, meta_description, meta_keywords, hide_from_nav, lft, rgt, depth, sort_order ` +
		`FROM public.category ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	c := Category{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&c.ID, &c.ParentID, &c.StoreID, &c.Slug, &c.Name, &c.Description, &c.MetaTitle, &c.MetaDescription, &c.MetaKeywords, &c.HideFromNav, &c.Lft, &c.Rgt, &c.Depth, &c.SortOrder); err != nil {
		return nil, logerror(err)
	}
	return &c, nil
}

// CategoryByStoreIDSlug retrieves a row from 'public.category' as a Category.
//
// Generated from index 'category_store_id_slug_key'.
func CategoryByStoreIDSlug(ctx context.Context, db DB, storeID int64, slug string) (*Category, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, parent_id, store_id, slug, name, description, meta_title, meta_description, meta_keywords, hide_from_nav, lft, rgt, depth, sort_order ` +
		`FROM public.category ` +
		`WHERE store_id = $1 AND slug = $2`
	// run
	logf(sqlstr, storeID, slug)
	c := Category{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, storeID, slug).Scan(&c.ID, &c.ParentID, &c.StoreID, &c.Slug, &c.Name, &c.Description, &c.MetaTitle, &c.MetaDescription, &c.MetaKeywords, &c.HideFromNav, &c.Lft, &c.Rgt, &c.Depth, &c.SortOrder); err != nil {
		return nil, logerror(err)
	}
	return &c, nil
}

// CategoryByDepth retrieves a row from 'public.category' as a Category.
//
// Generated from index 'idx_category_depth'.
func CategoryByDepth(ctx context.Context, db DB, depth sql.NullInt64) ([]*Category, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, parent_id, store_id, slug, name, description, meta_title, meta_description, meta_keywords, hide_from_nav, lft, rgt, depth, sort_order ` +
		`FROM public.category ` +
		`WHERE depth = $1`
	// run
	logf(sqlstr, depth)
	rows, err := db.QueryContext(ctx, sqlstr, depth)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Category
	for rows.Next() {
		c := Category{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&c.ID, &c.ParentID, &c.StoreID, &c.Slug, &c.Name, &c.Description, &c.MetaTitle, &c.MetaDescription, &c.MetaKeywords, &c.HideFromNav, &c.Lft, &c.Rgt, &c.Depth, &c.SortOrder); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// CategoryByLft retrieves a row from 'public.category' as a Category.
//
// Generated from index 'idx_category_lft'.
func CategoryByLft(ctx context.Context, db DB, lft sql.NullInt64) ([]*Category, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, parent_id, store_id, slug, name, description, meta_title, meta_description, meta_keywords, hide_from_nav, lft, rgt, depth, sort_order ` +
		`FROM public.category ` +
		`WHERE lft = $1`
	// run
	logf(sqlstr, lft)
	rows, err := db.QueryContext(ctx, sqlstr, lft)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Category
	for rows.Next() {
		c := Category{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&c.ID, &c.ParentID, &c.StoreID, &c.Slug, &c.Name, &c.Description, &c.MetaTitle, &c.MetaDescription, &c.MetaKeywords, &c.HideFromNav, &c.Lft, &c.Rgt, &c.Depth, &c.SortOrder); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// CategoryByParentID retrieves a row from 'public.category' as a Category.
//
// Generated from index 'idx_category_parent_id'.
func CategoryByParentID(ctx context.Context, db DB, parentID sql.NullInt64) ([]*Category, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, parent_id, store_id, slug, name, description, meta_title, meta_description, meta_keywords, hide_from_nav, lft, rgt, depth, sort_order ` +
		`FROM public.category ` +
		`WHERE parent_id = $1`
	// run
	logf(sqlstr, parentID)
	rows, err := db.QueryContext(ctx, sqlstr, parentID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Category
	for rows.Next() {
		c := Category{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&c.ID, &c.ParentID, &c.StoreID, &c.Slug, &c.Name, &c.Description, &c.MetaTitle, &c.MetaDescription, &c.MetaKeywords, &c.HideFromNav, &c.Lft, &c.Rgt, &c.Depth, &c.SortOrder); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// CategoryByRgt retrieves a row from 'public.category' as a Category.
//
// Generated from index 'idx_category_rgt'.
func CategoryByRgt(ctx context.Context, db DB, rgt sql.NullInt64) ([]*Category, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, parent_id, store_id, slug, name, description, meta_title, meta_description, meta_keywords, hide_from_nav, lft, rgt, depth, sort_order ` +
		`FROM public.category ` +
		`WHERE rgt = $1`
	// run
	logf(sqlstr, rgt)
	rows, err := db.QueryContext(ctx, sqlstr, rgt)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Category
	for rows.Next() {
		c := Category{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&c.ID, &c.ParentID, &c.StoreID, &c.Slug, &c.Name, &c.Description, &c.MetaTitle, &c.MetaDescription, &c.MetaKeywords, &c.HideFromNav, &c.Lft, &c.Rgt, &c.Depth, &c.SortOrder); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// CategoryByStoreID retrieves a row from 'public.category' as a Category.
//
// Generated from index 'idx_category_store_id'.
func CategoryByStoreID(ctx context.Context, db DB, storeID int64) ([]*Category, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, parent_id, store_id, slug, name, description, meta_title, meta_description, meta_keywords, hide_from_nav, lft, rgt, depth, sort_order ` +
		`FROM public.category ` +
		`WHERE store_id = $1`
	// run
	logf(sqlstr, storeID)
	rows, err := db.QueryContext(ctx, sqlstr, storeID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Category
	for rows.Next() {
		c := Category{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&c.ID, &c.ParentID, &c.StoreID, &c.Slug, &c.Name, &c.Description, &c.MetaTitle, &c.MetaDescription, &c.MetaKeywords, &c.HideFromNav, &c.Lft, &c.Rgt, &c.Depth, &c.SortOrder); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Category returns the Category associated with the Category's (ParentID).
//
// Generated from foreign key 'category_parent_id_fkey'.
func (c *Category) Category(ctx context.Context, db DB) (*Category, error) {
	return CategoryByID(ctx, db, c.ParentID.Int64)
}
