package query

import (
	"context"
	"io"

	"github.com/gogo-framework/db/dialect"
)

// LimitOffsetClause represents a LIMIT and OFFSET clause
type LimitOffsetClause struct {
	Limit  *int
	Offset *int
}

func (l *LimitOffsetClause) ApplySelect(stmt *SelectStmt) {
	stmt.limitOffset = l
}

func (l *LimitOffsetClause) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	if l.Limit == nil && l.Offset == nil {
		return nil, nil
	}
	if _, err := w.Write([]byte(d.LimitOffset(l.Limit, l.Offset))); err != nil {
		return nil, err
	}
	return nil, nil
}

// LimitOffset creates a LIMIT and OFFSET clause
func LimitOffset(limit *int, offset *int) SelectPart {
	return &LimitOffsetClause{
		Limit:  limit,
		Offset: offset,
	}
}

// Limit creates a LIMIT clause
func Limit(limit int) SelectPart {
	return &LimitOffsetClause{
		Limit: &limit,
	}
}

// Offset creates an OFFSET clause
func Offset(offset int) SelectPart {
	return &LimitOffsetClause{
		Offset: &offset,
	}
}
