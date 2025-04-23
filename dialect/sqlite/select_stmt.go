package sqlite

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/gogo-framework/db/dialect"
	"github.com/gogo-framework/db/internal/query"
)

// SelectPart represents a part of a SELECT statement that can be applied to a SelectStmt
type SelectPart interface {
	ApplySelect(*SelectStmt)
}

// SelectStmt represents a SQLite SELECT statement
type SelectStmt struct {
	// can be removed later, right now it's used in the toSql method but this is only for testing
	dialect     dialect.Dialect
	Columns     *SelectClause
	distinct    *DistinctClause
	from        *FromClause
	where       *WhereClause
	groupBy     *GroupByClause
	having      *HavingClause
	orderBy     *OrderByClause
	limitOffset *query.LimitOffsetClause
}

// WriteSql generates the SQL for the SELECT statement
func (s *SelectStmt) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	var args []any

	// Write SELECT
	w.Write([]byte("SELECT "))
	if s.distinct != nil {
		distinctArgs, err := s.distinct.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing DISTINCT: %w", err)
		}
		args = append(args, distinctArgs...)
	}
	if s.Columns != nil {
		columnArgs, err := s.Columns.WriteSql(ctx, w, d, argPos+len(args))
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

	// Write GROUP BY
	if s.groupBy != nil {
		if _, err := w.Write([]byte(" GROUP BY ")); err != nil {
			return nil, fmt.Errorf("error writing GROUP BY: %w", err)
		}
		groupByArgs, err := s.groupBy.WriteSql(ctx, w, d, argPos+len(args))
		if err != nil {
			return nil, fmt.Errorf("error writing GROUP BY clause: %w", err)
		}
		args = append(args, groupByArgs...)
	}

	// Write HAVING
	if s.having != nil {
		if _, err := w.Write([]byte(" HAVING ")); err != nil {
			return nil, fmt.Errorf("error writing HAVING: %w", err)
		}
		havingArgs, err := s.having.WriteSql(ctx, w, d, argPos+len(args))
		if err != nil {
			return nil, fmt.Errorf("error writing HAVING clause: %w", err)
		}
		args = append(args, havingArgs...)
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

	// Write LIMIT and OFFSET
	if s.limitOffset != nil {
		if _, err := w.Write([]byte(" ")); err != nil {
			return nil, fmt.Errorf("error writing LIMIT/OFFSET: %w", err)
		}
		limitOffsetArgs, err := s.limitOffset.WriteSql(ctx, w, d, argPos+len(args))
		if err != nil {
			return nil, fmt.Errorf("error writing LIMIT/OFFSET clause: %w", err)
		}
		args = append(args, limitOffsetArgs...)
	}

	return args, nil
}

// ToSql converts the statement to SQL string with arguments (for testing)
func (s *SelectStmt) ToSql() (string, []any) {
	ctx := context.Background()
	w := &bytes.Buffer{}
	args, _ := s.WriteSql(ctx, w, s.dialect, 1)
	return w.String(), args
}
