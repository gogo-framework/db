package sqlite

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/gogo-framework/db/dialect"
	"github.com/gogo-framework/db/query"
)

type SelectPart interface {
	query.SqlWriter
	ApplySelect(stmt *SelectStmt)
}

type SelectStmt struct {
	columns *SelectClause
	from    *FromClause
	where   *WhereClause
	orderBy *OrderByClause
	limit   *LimitClause
	offset  *OffsetClause
}

// WriteSql implements the SqlWriter interface
func (s *SelectStmt) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	if s.columns == nil || len(s.columns.Columns) == 0 {
		return nil, fmt.Errorf("no columns selected")
	}

	var qArgs []any

	// Write SELECT
	w.Write([]byte("SELECT "))
	args, err := s.columns.WriteSql(ctx, w, d, argPos)
	if err != nil {
		return nil, fmt.Errorf("error writing columns: %w", err)
	}
	qArgs = append(qArgs, args...)
	argPos += len(args)

	// Write FROM
	if s.from != nil {
		w.Write([]byte(" FROM "))
		args, err := s.from.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing from: %w", err)
		}
		qArgs = append(qArgs, args...)
		argPos += len(args)
	}

	// Write WHERE
	if s.where != nil {
		w.Write([]byte(" WHERE "))
		args, err := s.where.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing where: %w", err)
		}
		qArgs = append(qArgs, args...)
		argPos += len(args)
	}

	// Write ORDER BY
	if s.orderBy != nil {
		w.Write([]byte(" ORDER BY "))
		args, err := s.orderBy.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing order by: %w", err)
		}
		qArgs = append(qArgs, args...)
		argPos += len(args)
	}

	// Write LIMIT
	if s.limit != nil {
		w.Write([]byte(" LIMIT "))
		args, err := s.limit.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing limit: %w", err)
		}
		qArgs = append(qArgs, args...)
		argPos += len(args)
	}

	// Write OFFSET
	if s.offset != nil {
		w.Write([]byte(" OFFSET "))
		args, err := s.offset.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing offset: %w", err)
		}
		qArgs = append(qArgs, args...)
		argPos += len(args)
	}

	return qArgs, nil
}

// ToSql converts the statement to SQL string with arguments (for testing)
func (s *SelectStmt) ToSql() (string, []any) {
	ctx := context.Background()
	w := &bytes.Buffer{}
	args, _ := s.WriteSql(ctx, w, nil, 1)
	return w.String(), args
}
