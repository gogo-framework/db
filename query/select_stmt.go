package query

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/gogo-framework/db/dialect"
)

// SelectPart represents a part of a SELECT statement that can be applied to a SelectStmt
type SelectPart interface {
	ApplySelect(stmt *SelectStmt)
}

// SelectStmt represents a SELECT statement in the query package.
// This is the base implementation that dialect packages can embed or use as a reference.
type SelectStmt struct {
	columns *SelectClause
	from    *FromClause
	where   *WhereClause
	groupBy *GroupByClause
	having  *HavingClause
	orderBy *OrderByClause
	limit   *LimitClause
	offset  *OffsetClause
}

// WriteSql implements the SqlWriter interface
func (s *SelectStmt) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	if s.columns == nil || len(s.columns.Columns) == 0 {
		return nil, fmt.Errorf("no columns selected")
	}

	var args []any

	// Write SELECT
	w.Write([]byte("SELECT "))
	columnArgs, err := s.columns.WriteSql(ctx, w, d, argPos)
	if err != nil {
		return nil, fmt.Errorf("error writing columns: %w", err)
	}
	args = append(args, columnArgs...)
	argPos += len(columnArgs)

	// Write FROM
	if s.from != nil {
		w.Write([]byte(" FROM "))
		fromArgs, err := s.from.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing from: %w", err)
		}
		args = append(args, fromArgs...)
		argPos += len(fromArgs)
	}

	// Write WHERE
	if s.where != nil {
		w.Write([]byte(" WHERE "))
		whereArgs, err := s.where.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing where: %w", err)
		}
		args = append(args, whereArgs...)
		argPos += len(whereArgs)
	}

	// Write GROUP BY
	if s.groupBy != nil {
		w.Write([]byte(" GROUP BY "))
		groupByArgs, err := s.groupBy.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing group by: %w", err)
		}
		args = append(args, groupByArgs...)
		argPos += len(groupByArgs)
	}

	// Write HAVING
	if s.having != nil {
		w.Write([]byte(" HAVING "))
		havingArgs, err := s.having.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing having: %w", err)
		}
		args = append(args, havingArgs...)
		argPos += len(havingArgs)
	}

	// Write ORDER BY
	if s.orderBy != nil {
		w.Write([]byte(" ORDER BY "))
		orderByArgs, err := s.orderBy.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing order by: %w", err)
		}
		args = append(args, orderByArgs...)
		argPos += len(orderByArgs)
	}

	// Write LIMIT
	if s.limit != nil {
		w.Write([]byte(" LIMIT "))
		limitArgs, err := s.limit.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing limit: %w", err)
		}
		args = append(args, limitArgs...)
		argPos += len(limitArgs)
	}

	// Write OFFSET
	if s.offset != nil {
		w.Write([]byte(" OFFSET "))
		offsetArgs, err := s.offset.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing offset: %w", err)
		}
		args = append(args, offsetArgs...)
		argPos += len(offsetArgs)
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
