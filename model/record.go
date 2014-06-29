package model

import (
	"database/sql"
)

type Record interface {
	InsertSql() string
	UpdateSql() string
	DeleteSql() string
	Scan(*sql.Rows) error
}

