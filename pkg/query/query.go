package query

import (
	"context"
	"io"

	"github.com/gogo-framework/db/pkg/dialect"
)

type QueryWriter interface {
	WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error)
}
