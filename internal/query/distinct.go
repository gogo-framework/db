package query

import (
	"context"
	"io"

	"github.com/gogo-framework/db/dialect"
)

// DistinctClause represents a DISTINCT clause in a SELECT statement
type DistinctClause struct {
	OnExpr Expression
}

func (dc *DistinctClause) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	w.Write([]byte("DISTINCT "))
	return nil, nil
}

func (dc *DistinctClause) On(expr Expression) *DistinctClause {
	dc.OnExpr = expr
	return dc
}
