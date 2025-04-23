package query

import (
	"context"
	"io"

	"github.com/gogo-framework/db/dialect"
)

// GroupByClause represents a GROUP BY clause
type GroupByClause struct {
	Columns []Expression
}

func (g *GroupByClause) ApplySelect(stmt *SelectStmt) {
	stmt.groupBy = g
}

func (g *GroupByClause) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	if len(g.Columns) == 0 {
		return nil, nil
	}

	var allArgs []any
	for i, col := range g.Columns {
		if i > 0 {
			w.Write([]byte(", "))
		}
		args, err := col.WriteSql(ctx, w, d, argPos+len(allArgs))
		if err != nil {
			return nil, err
		}
		allArgs = append(allArgs, args...)
	}
	return allArgs, nil
}

// GroupBy creates a GROUP BY clause
func GroupBy(columns ...Expression) SelectPart {
	return &GroupByClause{
		Columns: columns,
	}
}
