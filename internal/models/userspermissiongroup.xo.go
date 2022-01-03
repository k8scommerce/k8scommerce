package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// UsersPermissionGroup represents a row from 'public.users_permission_group'.
type UsersPermissionGroup struct {
	UserID            int64 `json:"user_id" db:"user_id"`                         // user_id
	PermissionGroupID int64 `json:"permission_group_id" db:"permission_group_id"` // permission_group_id
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the UsersPermissionGroup exists in the database.
func (upg *UsersPermissionGroup) Exists() bool {
	return upg._exists
}

// Deleted returns true when the UsersPermissionGroup has been marked for deletion from
// the database.
func (upg *UsersPermissionGroup) Deleted() bool {
	return upg._deleted
}

// Insert inserts the UsersPermissionGroup to the database.
func (upg *UsersPermissionGroup) Insert(ctx context.Context, db DB) error {
	switch {
	case upg._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case upg._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.users_permission_group (` +
		`` +
		`) VALUES (` +
		`` +
		`) RETURNING permission_group_id`
	// run
	logf(sqlstr)
	if err := db.QueryRowContext(ctx, sqlstr).Scan(&upg.UserID); err != nil {
		return logerror(err)
	}
	// set exists
	upg._exists = true
	return nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the UsersPermissionGroup from the database.
func (upg *UsersPermissionGroup) Delete(ctx context.Context, db DB) error {
	switch {
	case !upg._exists: // doesn't exist
		return nil
	case upg._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	const sqlstr = `DELETE FROM public.users_permission_group ` +
		`WHERE user_id = $1 AND permission_group_id = $2`
	// run
	logf(sqlstr, upg.UserID, upg.PermissionGroupID)
	if _, err := db.ExecContext(ctx, sqlstr, upg.UserID, upg.PermissionGroupID); err != nil {
		return logerror(err)
	}
	// set deleted
	upg._deleted = true
	return nil
}

// UsersPermissionGroupByUserIDPermissionGroupID retrieves a row from 'public.users_permission_group' as a UsersPermissionGroup.
//
// Generated from index 'users_permission_group_pkey'.
func UsersPermissionGroupByUserIDPermissionGroupID(ctx context.Context, db DB, userID, permissionGroupID int64) (*UsersPermissionGroup, error) {
	// query
	const sqlstr = `SELECT ` +
		`user_id, permission_group_id ` +
		`FROM public.users_permission_group ` +
		`WHERE user_id = $1 AND permission_group_id = $2`
	// run
	logf(sqlstr, userID, permissionGroupID)
	upg := UsersPermissionGroup{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, userID, permissionGroupID).Scan(&upg.UserID, &upg.PermissionGroupID); err != nil {
		return nil, logerror(err)
	}
	return &upg, nil
}

// PermissionGroup returns the PermissionGroup associated with the UsersPermissionGroup's (PermissionGroupID).
//
// Generated from foreign key 'users_permission_group_permission_group_id_fkey'.
func (upg *UsersPermissionGroup) PermissionGroup(ctx context.Context, db DB) (*PermissionGroup, error) {
	return PermissionGroupByID(ctx, db, upg.PermissionGroupID)
}

// User returns the User associated with the UsersPermissionGroup's (UserID).
//
// Generated from foreign key 'users_permission_group_user_id_fkey'.
func (upg *UsersPermissionGroup) User(ctx context.Context, db DB) (*User, error) {
	return UserByID(ctx, db, upg.UserID)
}