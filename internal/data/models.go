package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflicts")
)

type Models struct {
	Movies      MovieModel
	Users       UserModel // Add a new Users field
	Tokens      TokenModel
	Permissions PermissionModel
}

func NewModles(db *sql.DB) Models {
	return Models{
		Movies:      MovieModel{DB: db},
		Users:       UserModel{DB: db}, // Add a new Users field
		Tokens:      TokenModel{DB: db},
		Permissions: PermissionModel{DB: db},
	}
}
