package query

import (
	"context"
	"fmt"
	"io"

	"github.com/gogo-framework/db/dialect"
	"github.com/gogo-framework/db/schema"
)

// SelectClause represents a SELECT clause
type SelectClause struct {
	Columns []schema.Column
}

func (s *SelectClause) ApplySelect(stmt *SelectStmt) {
	stmt.columns = s
}

func (s *SelectClause) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	if len(s.Columns) == 0 {
		return nil, fmt.Errorf("no columns selected")
	}

	var allArgs []any
	for i, col := range s.Columns {
		if i > 0 {
			w.Write([]byte(", "))
		}
		args, err := col.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing column: %w", err)
		}
		allArgs = append(allArgs, args...)
		argPos += len(args)
	}
	return allArgs, nil
}

// NewSelectClause creates a new SELECT clause
func NewSelectClause(columns ...schema.Column) *SelectClause {
	return &SelectClause{
		Columns: columns,
	}
}
