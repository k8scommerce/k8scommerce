// Package models contains generated code for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql/driver"
	"fmt"
)

// AddressKind is the 'address_kind' enum type from schema 'public'.
type AddressKind uint16

// AddressKind values.
const (
	// AddressKindBilling is the 'billing' address_kind.
	AddressKindBilling AddressKind = 1
	// AddressKindShipping is the 'shipping' address_kind.
	AddressKindShipping AddressKind = 2
	// AddressKindMailing is the 'mailing' address_kind.
	AddressKindMailing AddressKind = 3
)

// String satisfies the fmt.Stringer interface.
func (ak AddressKind) String() string {
	switch ak {
	case AddressKindBilling:
		return "billing"
	case AddressKindShipping:
		return "shipping"
	case AddressKindMailing:
		return "mailing"
	}
	return fmt.Sprintf("AddressKind(%d)", ak)
}

// MarshalText marshals AddressKind into text.
func (ak AddressKind) MarshalText() ([]byte, error) {
	return []byte(ak.String()), nil
}

// UnmarshalText unmarshals AddressKind from text.
func (ak *AddressKind) UnmarshalText(buf []byte) error {
	switch str := string(buf); str {
	case "billing":
		*ak = AddressKindBilling
	case "shipping":
		*ak = AddressKindShipping
	case "mailing":
		*ak = AddressKindMailing
	default:
		return ErrInvalidAddressKind(str)
	}
	return nil
}

// Value satisfies the driver.Valuer interface.
func (ak AddressKind) Value() (driver.Value, error) {
	return ak.String(), nil
}

// Scan satisfies the sql.Scanner interface.
func (ak *AddressKind) Scan(v interface{}) error {
	if buf, ok := v.([]byte); ok {
		return ak.UnmarshalText(buf)
	}
	if buf, ok := v.(string); ok {
		return ak.UnmarshalText([]byte(buf))
	}
	return ErrInvalidAddressKind(fmt.Sprintf("%T", v))
}

// NullAddressKind represents a null 'address_kind' enum for schema 'public'.
type NullAddressKind struct {
	AddressKind AddressKind
	// Valid is true if AddressKind is not null.
	Valid bool
}

// Value satisfies the driver.Valuer interface.
func (nak NullAddressKind) Value() (driver.Value, error) {
	if !nak.Valid {
		return nil, nil
	}
	return nak.AddressKind.Value()
}

// Scan satisfies the sql.Scanner interface.
func (nak *NullAddressKind) Scan(v interface{}) error {
	if v == nil {
		nak.AddressKind, nak.Valid = 0, false
		return nil
	}
	err := nak.AddressKind.Scan(v)
	nak.Valid = err == nil
	return err
}

// ErrInvalidAddressKind is the invalid AddressKind error.
type ErrInvalidAddressKind string

// Error satisfies the error interface.
func (err ErrInvalidAddressKind) Error() string {
	return fmt.Sprintf("invalid AddressKind(%s)", string(err))
}
