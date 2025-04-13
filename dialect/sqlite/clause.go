package sqlite

import (
	"context"
	"fmt"
	"io"

	"github.com/gogo-framework/db/dialect"
	"github.com/gogo-framework/db/query"
	"github.com/gogo-framework/db/schema"
)

// SelectClause represents a SELECT clause in SQLite
type SelectClause struct {
	*query.SelectClause
}

func (s *SelectClause) ApplySelect(stmt *SelectStmt) {
	stmt.columns = s
}

// Select creates a new SQLite SELECT statement
func Select(parts ...SelectPart) *SelectStmt {
	stmt := &SelectStmt{
		dialect: &SqliteDialect{},
	}
	for _, part := range parts {
		if part != nil {
			part.ApplySelect(stmt)
		}
	}
	return stmt
}

// sqlWriterTabler is a wrapper that implements both Tabler and SqlWriter
type sqlWriterTabler struct {
	writer query.SqlWriter
}

func (s *sqlWriterTabler) Table() *schema.Table {
	// For SqlWriter, we create a dummy table since we don't need it for SQL generation
	return &schema.Table{Name: "dummy"}
}

func (s *sqlWriterTabler) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	return s.writer.WriteSql(ctx, w, d, argPos)
}

// FromClause represents a FROM clause in SQLite
type FromClause struct {
	*query.FromClause
	invalidSource bool
}

func (f *FromClause) ApplySelect(stmt *SelectStmt) {
	stmt.from = f
}

func (f *FromClause) WriteSql(ctx context.Context, writer io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	if f.invalidSource {
		return nil, fmt.Errorf("invalid source type for FROM clause")
	}
	return f.FromClause.WriteSql(ctx, writer, d, argPos)
}

func (f *FromClause) As(alias string) SelectPart {
	f.FromClause.As(alias)
	return f
}

// From creates a FROM clause
func From(source any) *FromClause {
	if tabler, ok := source.(schema.Tabler); ok {
		return &FromClause{
			FromClause: &query.FromClause{
				Source: tabler,
			},
		}
	}

	if writer, ok := source.(query.SqlWriter); ok {
		wrapper := &sqlWriterTabler{writer: writer}
		return &FromClause{
			FromClause: &query.FromClause{
				Source: wrapper,
			},
		}
	}

	return &FromClause{
		FromClause:    &query.FromClause{},
		invalidSource: true,
	}
}

// WhereClause represents a WHERE clause in SQLite
type WhereClause struct {
	*query.WhereClause
}

func (w *WhereClause) ApplySelect(stmt *SelectStmt) {
	stmt.where = w
}

// Where creates a WHERE clause
func Where(conditions ...query.Condition) *WhereClause {
	return &WhereClause{
		WhereClause: &query.WhereClause{
			Conditions: conditions,
		},
	}
}

// And adds additional conditions to an existing WHERE clause
func (w *WhereClause) And(conditions ...query.Condition) *WhereClause {
	w.Conditions = append(w.Conditions, conditions...)
	return w
}

// OrderByClause represents an ORDER BY clause in SQLite
type OrderByClause struct {
	*query.OrderByClause
}

func (o *OrderByClause) ApplySelect(stmt *SelectStmt) {
	stmt.orderBy = o
}

// OrderBy creates an ORDER BY clause
func OrderBy(columns ...query.SqlWriter) *OrderByClause {
	return &OrderByClause{
		OrderByClause: &query.OrderByClause{
			Columns: columns,
		},
	}
}

// LimitOffsetClause represents a LIMIT and OFFSET clause in SQLite
type LimitOffsetClause struct {
	*query.LimitOffsetClause
}

func (l *LimitOffsetClause) ApplySelect(stmt *SelectStmt) {
	stmt.limitOffset = l.LimitOffsetClause
}

// LimitOffset creates a LIMIT and OFFSET clause
func LimitOffset(limit *int, offset *int) *LimitOffsetClause {
	return &LimitOffsetClause{
		LimitOffsetClause: &query.LimitOffsetClause{
			Limit:  limit,
			Offset: offset,
		},
	}
}

// Limit creates a LIMIT clause
func Limit(limit int) *LimitOffsetClause {
	return &LimitOffsetClause{
		LimitOffsetClause: &query.LimitOffsetClause{
			Limit: &limit,
		},
	}
}

// Offset creates an OFFSET clause
func Offset(offset int) *LimitOffsetClause {
	return &LimitOffsetClause{
		LimitOffsetClause: &query.LimitOffsetClause{
			Offset: &offset,
		},
	}
}

// DistinctClause represents a DISTINCT clause in SQLite
type DistinctClause struct {
	*query.DistinctClause
}

func (d *DistinctClause) ApplySelect(stmt *SelectStmt) {
	stmt.distinct = d
}

// Distinct creates a DISTINCT clause
func Distinct() *DistinctClause {
	return &DistinctClause{
		DistinctClause: &query.DistinctClause{},
	}
}

// GroupByClause represents a GROUP BY clause in SQLite
type GroupByClause struct {
	*query.GroupByClause
}

func (g *GroupByClause) ApplySelect(stmt *SelectStmt) {
	stmt.groupBy = g
}

// GroupBy creates a GROUP BY clause
func GroupBy(columns ...query.SqlWriter) *GroupByClause {
	return &GroupByClause{
		GroupByClause: &query.GroupByClause{
			Columns: columns,
		},
	}
}

// HavingClause represents a HAVING clause in SQLite
type HavingClause struct {
	*query.HavingClause
}

func (h *HavingClause) ApplySelect(stmt *SelectStmt) {
	stmt.having = h
}

// Having creates a HAVING clause
func Having(conditions ...query.Condition) *HavingClause {
	return &HavingClause{
		HavingClause: &query.HavingClause{
			Conditions: conditions,
		},
	}
}
