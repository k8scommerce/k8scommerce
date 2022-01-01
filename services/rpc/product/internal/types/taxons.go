package types

import (
	"time"
)

type Taxon struct {
	Id              int64      `db:"id" json:"id"`
	ParentId        int64      `db:"parent_id" json:"parent_id"`
	TaxonomyId      int64      `db:"taxonomy_id" json:"taxonomy_id"`
	Name            string     `db:"name" json:"name"`
	Description     string     `db:"description" json:"description"`
	Permalink       NullString `db:"permalink" json:"permalink"`
	MetaTitle       NullString `db:"meta_title" json:"meta_title"`
	MetaDescription NullString `db:"meta_description" json:"meta_description"`
	MetaKeywords    NullString `db:"meta_keywords" json:"meta_keywords"`
	HideFromNav     bool       `db:"hide_from_nav" json:"hide_from_nav"`
	Lft             NullInt64  `db:"lft" json:"lft"`
	Rgt             NullInt64  `db:"rgt" json:"rgt"`
	Depth           NullInt64  `db:"depth" json:"depth"`
	Position        NullInt32  `db:"position" json:"position"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt       NullTime   `db:"deleted_at" json:"deleted_at"`
}
