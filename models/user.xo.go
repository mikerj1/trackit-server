// Package models contains the types for schema 'trackit'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// User represents a row from 'trackit.user'.
type User struct {
	ID                     int            `json:"id"`                       // id
	Email                  string         `json:"email"`                    // email
	Auth                   string         `json:"auth"`                     // auth
	NextExternal           sql.NullString `json:"next_external"`            // next_external
	ParentUserID           sql.NullInt64  `json:"parent_user_id"`           // parent_user_id
	AwsCustomerIdentifier  string         `json:"aws_customer_identifier"`  // aws_customer_identifier
	AwsCustomerEntitlement bool           `json:"aws_customer_entitlement"` // aws_customer_entitlement

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted provides information if the User has been deleted from the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO trackit.user (` +
		`email, auth, next_external, parent_user_id, aws_customer_identifier, aws_customer_entitlement` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, u.Email, u.Auth, u.NextExternal, u.ParentUserID, u.AwsCustomerIdentifier, u.AwsCustomerEntitlement)
	res, err := db.Exec(sqlstr, u.Email, u.Auth, u.NextExternal, u.ParentUserID, u.AwsCustomerIdentifier, u.AwsCustomerEntitlement)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	u.ID = int(id)
	u._exists = true

	return nil
}

// Update updates the User in the database.
func (u *User) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if u._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE trackit.user SET ` +
		`email = ?, auth = ?, next_external = ?, parent_user_id = ?, aws_customer_identifier = ?, aws_customer_entitlement = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, u.Email, u.Auth, u.NextExternal, u.ParentUserID, u.AwsCustomerIdentifier, u.AwsCustomerEntitlement, u.ID)
	_, err = db.Exec(sqlstr, u.Email, u.Auth, u.NextExternal, u.ParentUserID, u.AwsCustomerIdentifier, u.AwsCustomerEntitlement, u.ID)
	return err
}

// Save saves the User to the database.
func (u *User) Save(db XODB) error {
	if u.Exists() {
		return u.Update(db)
	}

	return u.Insert(db)
}

// Delete deletes the User from the database.
func (u *User) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return nil
	}

	// if deleted, bail
	if u._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM trackit.user WHERE id = ?`

	// run query
	XOLog(sqlstr, u.ID)
	_, err = db.Exec(sqlstr, u.ID)
	if err != nil {
		return err
	}

	// set deleted
	u._deleted = true

	return nil
}

// User returns the User associated with the User's ParentUserID (parent_user_id).
//
// Generated from foreign key 'parent_user'.
func (u *User) User(db XODB) (*User, error) {
	return UserByID(db, int(u.ParentUserID.Int64))
}

// UsersByParentUserID retrieves a row from 'trackit.user' as a User.
//
// Generated from index 'parent_user'.
func UsersByParentUserID(db XODB, parentUserID sql.NullInt64) ([]*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, email, auth, next_external, parent_user_id, aws_customer_identifier, aws_customer_entitlement ` +
		`FROM trackit.user ` +
		`WHERE parent_user_id = ?`

	// run query
	XOLog(sqlstr, parentUserID)
	q, err := db.Query(sqlstr, parentUserID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*User{}
	for q.Next() {
		u := User{
			_exists: true,
		}

		// scan
		err = q.Scan(&u.ID, &u.Email, &u.Auth, &u.NextExternal, &u.ParentUserID, &u.AwsCustomerIdentifier, &u.AwsCustomerEntitlement)
		if err != nil {
			return nil, err
		}

		res = append(res, &u)
	}

	return res, nil
}

// UserByEmail retrieves a row from 'trackit.user' as a User.
//
// Generated from index 'unique_email'.
func UserByEmail(db XODB, email string) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, email, auth, next_external, parent_user_id, aws_customer_identifier, aws_customer_entitlement ` +
		`FROM trackit.user ` +
		`WHERE email = ?`

	// run query
	XOLog(sqlstr, email)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, email).Scan(&u.ID, &u.Email, &u.Auth, &u.NextExternal, &u.ParentUserID, &u.AwsCustomerIdentifier, &u.AwsCustomerEntitlement)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// UserByID retrieves a row from 'trackit.user' as a User.
//
// Generated from index 'user_id_pkey'.
func UserByID(db XODB, id int) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, email, auth, next_external, parent_user_id, aws_customer_identifier, aws_customer_entitlement ` +
		`FROM trackit.user ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&u.ID, &u.Email, &u.Auth, &u.NextExternal, &u.ParentUserID, &u.AwsCustomerIdentifier, &u.AwsCustomerEntitlement)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
