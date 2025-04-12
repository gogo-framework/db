package query

import (
	"context"
	"fmt"
	"io"

	"github.com/gogo-framework/db/dialect"
)

// LimitClause represents a LIMIT clause
type LimitClause struct {
	Limit int
}

func (l *LimitClause) ApplySelect(stmt *SelectStmt) {
	stmt.limit = l
}

func (l *LimitClause) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	w.Write([]byte(fmt.Sprintf("%d", l.Limit)))
	return nil, nil
}

// Limit creates a LIMIT clause
func Limit(limit int) SelectPart {
	return &LimitClause{
		Limit: limit,
	}
}

// OffsetClause represents an OFFSET clause
type OffsetClause struct {
	Offset int
}

func (o *OffsetClause) ApplySelect(stmt *SelectStmt) {
	stmt.offset = o
}

func (o *OffsetClause) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	w.Write([]byte(fmt.Sprintf("%d", o.Offset)))
	return nil, nil
}

// Offset creates an OFFSET clause
func Offset(offset int) SelectPart {
	return &OffsetClause{
		Offset: offset,
	}
}
