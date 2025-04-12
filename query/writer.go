package query

import (
	"context"
	"io"

	"github.com/gogo-framework/db/dialect"
)

// SqlWriter defines the interface for SQL generation
type SqlWriter interface {
	WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error)
}
