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
	Movies MovieModel
}

func NewModles(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
