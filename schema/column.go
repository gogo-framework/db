package schema

import (
	"database/sql"
	"database/sql/driver"
)

type Column interface {
	sql.Scanner
	driver.Valuer
	GetTable() *Table
	SetTable(*Table)
	GetName() string
	SetName(string)
	GetType() string
}
