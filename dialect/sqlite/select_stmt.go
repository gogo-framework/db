package sqlite

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/gogo-framework/db/dialect"
)

type SelectPart interface {
	ApplySelect(*SelectStmt)
}

// https://www.sqlite.org/lang_select.html
type SelectStmt struct {
	with         any
	selectClause SelectClause
	distinct     bool
	from         FromClause
	where        any
	groupBy      any
	having       any
	windows      any
	orderBy      any
	limit        any
	offset       any
}

func Select(parts ...SelectPart) *SelectStmt {
	stmt := &SelectStmt{}
	for _, part := range parts {
		part.ApplySelect(stmt)
	}
	return stmt
}

// Implement the WriteSql interface
func (stmt *SelectStmt) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	var args []any

	w.Write([]byte("SELECT "))
	stmt.selectClause.WriteSql(ctx, w, d, argPos)

	// From clause
	if stmt.from.Source != nil {
		w.Write([]byte(" FROM "))
		stmt.from.WriteSql(ctx, w, d, argPos)
	} else {
		return nil, errors.New("FROM clause is required for a SELECT statement")
	}

	return args, nil
}

// implement a tosql function for testing purposes. Will be removed later.
func (stmt *SelectStmt) ToSql() string {
	ctx := context.Background()
	w := &bytes.Buffer{}
	stmt.WriteSql(ctx, w, nil, 0)
	return w.String()
}
