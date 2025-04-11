package schema

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"io"

	"github.com/gogo-framework/db/dialect"
)

type Column interface {
	sql.Scanner
	driver.Valuer
	GetTable() *Table
	SetTable(*Table)
	GetName() string
	SetName(string)
	GetType() string
	WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error)
}
