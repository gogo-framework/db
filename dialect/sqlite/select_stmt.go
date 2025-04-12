package sqlite

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/gogo-framework/db/dialect"
	"github.com/gogo-framework/db/query"
)

// SelectPart represents a part of a SELECT statement that can be applied to a SelectStmt
type SelectPart interface {
	ApplySelect(*SelectStmt)
}

// SelectStmt represents a SQLite SELECT statement
type SelectStmt struct {
	*query.SelectStmt
	columns *SelectClause
	from    *FromClause
	where   *WhereClause
	orderBy *OrderByClause
	limit   *LimitClause
	offset  *OffsetClause
}

// WriteSql generates the SQL for the SELECT statement
func (s *SelectStmt) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	var args []any

	// Write SELECT
	w.Write([]byte("SELECT "))
	if s.columns != nil {
		columnArgs, err := s.columns.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing SELECT: %w", err)
		}
		args = append(args, columnArgs...)
	} else {
		if _, err := w.Write([]byte("*")); err != nil {
			return nil, fmt.Errorf("error writing SELECT *: %w", err)
		}
	}

	// Write FROM
	if s.from != nil {
		if _, err := w.Write([]byte(" FROM ")); err != nil {
			return nil, fmt.Errorf("error writing FROM: %w", err)
		}
		fromArgs, err := s.from.WriteSql(ctx, w, d, argPos+len(args))
		if err != nil {
			return nil, fmt.Errorf("error writing FROM clause: %w", err)
		}
		args = append(args, fromArgs...)
	}

	// Write WHERE
	if s.where != nil {
		if _, err := w.Write([]byte(" WHERE ")); err != nil {
			return nil, fmt.Errorf("error writing WHERE: %w", err)
		}
		whereArgs, err := s.where.WriteSql(ctx, w, d, argPos+len(args))
		if err != nil {
			return nil, fmt.Errorf("error writing WHERE clause: %w", err)
		}
		args = append(args, whereArgs...)
	}

	// Write ORDER BY
	if s.orderBy != nil {
		if _, err := w.Write([]byte(" ORDER BY ")); err != nil {
			return nil, fmt.Errorf("error writing ORDER BY: %w", err)
		}
		orderArgs, err := s.orderBy.WriteSql(ctx, w, d, argPos+len(args))
		if err != nil {
			return nil, fmt.Errorf("error writing ORDER BY clause: %w", err)
		}
		args = append(args, orderArgs...)
	}

	// Write LIMIT
	if s.limit != nil {
		if _, err := w.Write([]byte(" LIMIT ")); err != nil {
			return nil, fmt.Errorf("error writing LIMIT: %w", err)
		}
		limitArgs, err := s.limit.WriteSql(ctx, w, d, argPos+len(args))
		if err != nil {
			return nil, fmt.Errorf("error writing LIMIT clause: %w", err)
		}
		args = append(args, limitArgs...)
	}

	// Write OFFSET
	if s.offset != nil {
		if _, err := w.Write([]byte(" OFFSET ")); err != nil {
			return nil, fmt.Errorf("error writing OFFSET: %w", err)
		}
		offsetArgs, err := s.offset.WriteSql(ctx, w, d, argPos+len(args))
		if err != nil {
			return nil, fmt.Errorf("error writing OFFSET clause: %w", err)
		}
		args = append(args, offsetArgs...)
	}

	return args, nil
}

// ToSql converts the statement to SQL string with arguments (for testing)
func (s *SelectStmt) ToSql() (string, []any) {
	ctx := context.Background()
	w := &bytes.Buffer{}
	args, _ := s.WriteSql(ctx, w, nil, 1)
	return w.String(), args
}

// OrderBy adds an ORDER BY clause to the statement
func (s *SelectStmt) OrderBy(columns ...query.SqlWriter) *SelectStmt {
	s.orderBy = &OrderByClause{
		OrderByClause: &query.OrderByClause{
			Columns: columns,
		},
	}
	return s
}

// Limit adds a LIMIT clause to the statement
func (s *SelectStmt) Limit(limit int) *SelectStmt {
	s.limit = &LimitClause{
		LimitClause: &query.LimitClause{
			Limit: limit,
		},
	}
	return s
}

// Offset adds an OFFSET clause to the statement
func (s *SelectStmt) Offset(offset int) *SelectStmt {
	s.offset = &OffsetClause{
		OffsetClause: &query.OffsetClause{
			Offset: offset,
		},
	}
	return s
}
