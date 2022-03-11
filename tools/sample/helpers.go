package main

import (
	"database/sql"
)

func toNullString(in string) sql.NullString {
	out := sql.NullString{}
	out.String = in
	out.Valid = true
	return out
}

func toNullInt64(in int64) sql.NullInt64 {
	out := sql.NullInt64{}
	out.Int64 = in
	out.Valid = true
	return out
}

func toNullBool(in bool) sql.NullBool {
	out := sql.NullBool{}
	out.Bool = in
	out.Valid = true
	return out
}

func toNullFloat64(in float64) sql.NullFloat64 {
	out := sql.NullFloat64{}
	out.Float64 = in
	out.Valid = true
	return out
}
