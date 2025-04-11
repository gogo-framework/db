package query

import (
	"context"
	"io"

	"github.com/gogo-framework/db/dialect"
)

type SelectPart interface {
	ApplySelect(*SelectStmt)
}

type SelectStmt struct {
	Columns []SqlWriter
	from    SqlWriter
	where   WhereClause
	orderBy []SqlWriter
	limit   *int
	offset  *int
}

func (s *SelectStmt) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	var args []any

	// Write SELECT
	w.Write([]byte("SELECT "))

	// Write columns
	columnArgs, err := WriteSqlWithSeparator(ctx, w, d, argPos, ", ", s.Columns...)
	if err != nil {
		return nil, err
	}
	args = append(args, columnArgs...)

	// Write FROM
	if s.from != nil {
		w.Write([]byte(" FROM "))
		fromArgs, err := s.from.WriteSql(ctx, w, d, argPos+len(args))
		if err != nil {
			return nil, err
		}
		args = append(args, fromArgs...)
	}

	// Write WHERE
	if len(s.where.Conditions) > 0 {
		whereArgs, err := s.where.WriteSql(ctx, w, d, argPos+len(args))
		if err != nil {
			return nil, err
		}
		args = append(args, whereArgs...)
	}

	// Write ORDER BY
	if len(s.orderBy) > 0 {
		w.Write([]byte(" ORDER BY "))
		orderByArgs, err := WriteSqlWithSeparator(ctx, w, d, argPos+len(args), ", ", s.orderBy...)
		if err != nil {
			return nil, err
		}
		args = append(args, orderByArgs...)
	}

	// Write LIMIT
	if s.limit != nil {
		w.Write([]byte(" LIMIT "))
		limitArgs, err := NewLiteral(*s.limit).WriteSql(ctx, w, d, argPos+len(args))
		if err != nil {
			return nil, err
		}
		args = append(args, limitArgs...)
	}

	// Write OFFSET
	if s.offset != nil {
		w.Write([]byte(" OFFSET "))
		offsetArgs, err := NewLiteral(*s.offset).WriteSql(ctx, w, d, argPos+len(args))
		if err != nil {
			return nil, err
		}
		args = append(args, offsetArgs...)
	}

	return args, nil
}

func Select(columns ...SqlWriter) *SelectStmt {
	return &SelectStmt{
		Columns: columns,
	}
}

func (s *SelectStmt) From(table SqlWriter) *SelectStmt {
	s.from = table
	return s
}

func (s *SelectStmt) Where(conditions ...Condition) *SelectStmt {
	s.where = Where(conditions...)
	return s
}

func (s *SelectStmt) OrderBy(columns ...SqlWriter) *SelectStmt {
	s.orderBy = columns
	return s
}

func (s *SelectStmt) Limit(limit int) *SelectStmt {
	s.limit = &limit
	return s
}

func (s *SelectStmt) Offset(offset int) *SelectStmt {
	s.offset = &offset
	return s
}
