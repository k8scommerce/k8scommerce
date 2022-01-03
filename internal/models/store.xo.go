package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Store represents a row from 'public.store'.
type Store struct {
	ID                         int64          `json:"id" db:"id"`                                                       // id
	IsDefault                  bool           `json:"is_default" db:"is_default"`                                       // is_default
	Name                       string         `json:"name" db:"name"`                                                   // name
	Description                sql.NullString `json:"description" db:"description"`                                     // description
	URL                        string         `json:"url" db:"url"`                                                     // url
	SeoTitle                   sql.NullString `json:"seo_title" db:"seo_title"`                                         // seo_title
	SeoRobots                  sql.NullString `json:"seo_robots" db:"seo_robots"`                                       // seo_robots
	MetaDescription            sql.NullString `json:"meta_description" db:"meta_description"`                           // meta_description
	MetaKeywords               sql.NullString `json:"meta_keywords" db:"meta_keywords"`                                 // meta_keywords
	Facebook                   sql.NullString `json:"facebook" db:"facebook"`                                           // facebook
	Twitter                    sql.NullString `json:"twitter" db:"twitter"`                                             // twitter
	Instagram                  sql.NullString `json:"instagram" db:"instagram"`                                         // instagram
	Code                       sql.NullString `json:"code" db:"code"`                                                   // code
	DefaultCurrency            string         `json:"default_currency" db:"default_currency"`                           // default_currency
	SupportedCurrencies        sql.NullString `json:"supported_currencies" db:"supported_currencies"`                   // supported_currencies
	DefaultLocale              string         `json:"default_locale" db:"default_locale"`                               // default_locale
	SupportedLocales           sql.NullString `json:"supported_locales" db:"supported_locales"`                         // supported_locales
	DefaultCountryID           int64          `json:"default_country_id" db:"default_country_id"`                       // default_country_id
	Address                    sql.NullString `json:"address" db:"address"`                                             // address
	ContactPhone               sql.NullString `json:"contact_phone" db:"contact_phone"`                                 // contact_phone
	MailFromAddress            sql.NullString `json:"mail_from_address" db:"mail_from_address"`                         // mail_from_address
	CustomerSupportEmail       sql.NullString `json:"customer_support_email" db:"customer_support_email"`               // customer_support_email
	NewOrderNotificationsEmail sql.NullString `json:"new_order_notifications_email" db:"new_order_notifications_email"` // new_order_notifications_email
	CheckoutZoneID             sql.NullInt64  `json:"checkout_zone_id" db:"checkout_zone_id"`                           // checkout_zone_id
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Store exists in the database.
func (s *Store) Exists() bool {
	return s._exists
}

// Deleted returns true when the Store has been marked for deletion from
// the database.
func (s *Store) Deleted() bool {
	return s._deleted
}

// Insert inserts the Store to the database.
func (s *Store) Insert(ctx context.Context, db DB) error {
	switch {
	case s._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case s._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.store (` +
		`is_default, name, description, url, seo_title, seo_robots, meta_description, meta_keywords, facebook, twitter, instagram, code, default_currency, supported_currencies, default_locale, supported_locales, default_country_id, address, contact_phone, mail_from_address, customer_support_email, new_order_notifications_email, checkout_zone_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23` +
		`) RETURNING id`
	// run
	logf(sqlstr, s.IsDefault, s.Name, s.Description, s.URL, s.SeoTitle, s.SeoRobots, s.MetaDescription, s.MetaKeywords, s.Facebook, s.Twitter, s.Instagram, s.Code, s.DefaultCurrency, s.SupportedCurrencies, s.DefaultLocale, s.SupportedLocales, s.DefaultCountryID, s.Address, s.ContactPhone, s.MailFromAddress, s.CustomerSupportEmail, s.NewOrderNotificationsEmail, s.CheckoutZoneID)
	if err := db.QueryRowContext(ctx, sqlstr, s.IsDefault, s.Name, s.Description, s.URL, s.SeoTitle, s.SeoRobots, s.MetaDescription, s.MetaKeywords, s.Facebook, s.Twitter, s.Instagram, s.Code, s.DefaultCurrency, s.SupportedCurrencies, s.DefaultLocale, s.SupportedLocales, s.DefaultCountryID, s.Address, s.ContactPhone, s.MailFromAddress, s.CustomerSupportEmail, s.NewOrderNotificationsEmail, s.CheckoutZoneID).Scan(&s.ID); err != nil {
		return logerror(err)
	}
	// set exists
	s._exists = true
	return nil
}

// Update updates a Store in the database.
func (s *Store) Update(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case s._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.store SET ` +
		`is_default = $1, name = $2, description = $3, url = $4, seo_title = $5, seo_robots = $6, meta_description = $7, meta_keywords = $8, facebook = $9, twitter = $10, instagram = $11, code = $12, default_currency = $13, supported_currencies = $14, default_locale = $15, supported_locales = $16, default_country_id = $17, address = $18, contact_phone = $19, mail_from_address = $20, customer_support_email = $21, new_order_notifications_email = $22, checkout_zone_id = $23 ` +
		`WHERE id = $24`
	// run
	logf(sqlstr, s.IsDefault, s.Name, s.Description, s.URL, s.SeoTitle, s.SeoRobots, s.MetaDescription, s.MetaKeywords, s.Facebook, s.Twitter, s.Instagram, s.Code, s.DefaultCurrency, s.SupportedCurrencies, s.DefaultLocale, s.SupportedLocales, s.DefaultCountryID, s.Address, s.ContactPhone, s.MailFromAddress, s.CustomerSupportEmail, s.NewOrderNotificationsEmail, s.CheckoutZoneID, s.ID)
	if _, err := db.ExecContext(ctx, sqlstr, s.IsDefault, s.Name, s.Description, s.URL, s.SeoTitle, s.SeoRobots, s.MetaDescription, s.MetaKeywords, s.Facebook, s.Twitter, s.Instagram, s.Code, s.DefaultCurrency, s.SupportedCurrencies, s.DefaultLocale, s.SupportedLocales, s.DefaultCountryID, s.Address, s.ContactPhone, s.MailFromAddress, s.CustomerSupportEmail, s.NewOrderNotificationsEmail, s.CheckoutZoneID, s.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Store to the database.
func (s *Store) Save(ctx context.Context, db DB) error {
	if s.Exists() {
		return s.Update(ctx, db)
	}
	return s.Insert(ctx, db)
}

// Upsert performs an upsert for Store.
func (s *Store) Upsert(ctx context.Context, db DB) error {
	switch {
	case s._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.store (` +
		`id, is_default, name, description, url, seo_title, seo_robots, meta_description, meta_keywords, facebook, twitter, instagram, code, default_currency, supported_currencies, default_locale, supported_locales, default_country_id, address, contact_phone, mail_from_address, customer_support_email, new_order_notifications_email, checkout_zone_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`is_default = EXCLUDED.is_default, name = EXCLUDED.name, description = EXCLUDED.description, url = EXCLUDED.url, seo_title = EXCLUDED.seo_title, seo_robots = EXCLUDED.seo_robots, meta_description = EXCLUDED.meta_description, meta_keywords = EXCLUDED.meta_keywords, facebook = EXCLUDED.facebook, twitter = EXCLUDED.twitter, instagram = EXCLUDED.instagram, code = EXCLUDED.code, default_currency = EXCLUDED.default_currency, supported_currencies = EXCLUDED.supported_currencies, default_locale = EXCLUDED.default_locale, supported_locales = EXCLUDED.supported_locales, default_country_id = EXCLUDED.default_country_id, address = EXCLUDED.address, contact_phone = EXCLUDED.contact_phone, mail_from_address = EXCLUDED.mail_from_address, customer_support_email = EXCLUDED.customer_support_email, new_order_notifications_email = EXCLUDED.new_order_notifications_email, checkout_zone_id = EXCLUDED.checkout_zone_id `
	// run
	logf(sqlstr, s.ID, s.IsDefault, s.Name, s.Description, s.URL, s.SeoTitle, s.SeoRobots, s.MetaDescription, s.MetaKeywords, s.Facebook, s.Twitter, s.Instagram, s.Code, s.DefaultCurrency, s.SupportedCurrencies, s.DefaultLocale, s.SupportedLocales, s.DefaultCountryID, s.Address, s.ContactPhone, s.MailFromAddress, s.CustomerSupportEmail, s.NewOrderNotificationsEmail, s.CheckoutZoneID)
	if _, err := db.ExecContext(ctx, sqlstr, s.ID, s.IsDefault, s.Name, s.Description, s.URL, s.SeoTitle, s.SeoRobots, s.MetaDescription, s.MetaKeywords, s.Facebook, s.Twitter, s.Instagram, s.Code, s.DefaultCurrency, s.SupportedCurrencies, s.DefaultLocale, s.SupportedLocales, s.DefaultCountryID, s.Address, s.ContactPhone, s.MailFromAddress, s.CustomerSupportEmail, s.NewOrderNotificationsEmail, s.CheckoutZoneID); err != nil {
		return logerror(err)
	}
	// set exists
	s._exists = true
	return nil
}

// Delete deletes the Store from the database.
func (s *Store) Delete(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return nil
	case s._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.store ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, s.ID)
	if _, err := db.ExecContext(ctx, sqlstr, s.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	s._deleted = true
	return nil
}

// StoreByNameURL retrieves a row from 'public.store' as a Store.
//
// Generated from index 'store_name_url_key'.
func StoreByNameURL(ctx context.Context, db DB, name, url string) (*Store, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, is_default, name, description, url, seo_title, seo_robots, meta_description, meta_keywords, facebook, twitter, instagram, code, default_currency, supported_currencies, default_locale, supported_locales, default_country_id, address, contact_phone, mail_from_address, customer_support_email, new_order_notifications_email, checkout_zone_id ` +
		`FROM public.store ` +
		`WHERE name = $1 AND url = $2`
	// run
	logf(sqlstr, name, url)
	s := Store{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, name, url).Scan(&s.ID, &s.IsDefault, &s.Name, &s.Description, &s.URL, &s.SeoTitle, &s.SeoRobots, &s.MetaDescription, &s.MetaKeywords, &s.Facebook, &s.Twitter, &s.Instagram, &s.Code, &s.DefaultCurrency, &s.SupportedCurrencies, &s.DefaultLocale, &s.SupportedLocales, &s.DefaultCountryID, &s.Address, &s.ContactPhone, &s.MailFromAddress, &s.CustomerSupportEmail, &s.NewOrderNotificationsEmail, &s.CheckoutZoneID); err != nil {
		return nil, logerror(err)
	}
	return &s, nil
}

// StoreByID retrieves a row from 'public.store' as a Store.
//
// Generated from index 'store_pkey'.
func StoreByID(ctx context.Context, db DB, id int64) (*Store, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, is_default, name, description, url, seo_title, seo_robots, meta_description, meta_keywords, facebook, twitter, instagram, code, default_currency, supported_currencies, default_locale, supported_locales, default_country_id, address, contact_phone, mail_from_address, customer_support_email, new_order_notifications_email, checkout_zone_id ` +
		`FROM public.store ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	s := Store{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&s.ID, &s.IsDefault, &s.Name, &s.Description, &s.URL, &s.SeoTitle, &s.SeoRobots, &s.MetaDescription, &s.MetaKeywords, &s.Facebook, &s.Twitter, &s.Instagram, &s.Code, &s.DefaultCurrency, &s.SupportedCurrencies, &s.DefaultLocale, &s.SupportedLocales, &s.DefaultCountryID, &s.Address, &s.ContactPhone, &s.MailFromAddress, &s.CustomerSupportEmail, &s.NewOrderNotificationsEmail, &s.CheckoutZoneID); err != nil {
		return nil, logerror(err)
	}
	return &s, nil
}