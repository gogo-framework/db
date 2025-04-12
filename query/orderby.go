package query

import (
	"context"
	"fmt"
	"io"

	"github.com/gogo-framework/db/dialect"
)

// OrderByClause represents an ORDER BY clause
type OrderByClause struct {
	Columns []SqlWriter
}

func (o *OrderByClause) ApplySelect(stmt *SelectStmt) {
	stmt.orderBy = o
}

func (o *OrderByClause) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	if len(o.Columns) == 0 {
		return nil, nil
	}

	var allArgs []any
	for i, col := range o.Columns {
		if i > 0 {
			w.Write([]byte(", "))
		}
		args, err := col.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing order by column: %w", err)
		}
		allArgs = append(allArgs, args...)
		argPos += len(args)
	}
	return allArgs, nil
}

// OrderBy creates an ORDER BY clause
func OrderBy(columns ...SqlWriter) SelectPart {
	return &OrderByClause{
		Columns: columns,
	}
}
