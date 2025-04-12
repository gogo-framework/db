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
	stmt := &SelectStmt{}
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

// Or creates an OR condition
func Or(conditions ...query.Condition) query.Condition {
	return &query.OrCondition{Conditions: conditions}
}

// Condition builders that wrap query package functions
func Equal[T any](column schema.Column, value T) query.Condition {
	return query.Eq(column, value)
}

func NotEqual[T any](column schema.Column, value T) query.Condition {
	return query.Neq(column, value)
}

func GreaterThan[T any](column schema.Column, value T) query.Condition {
	return query.Gt(column, value)
}

func GreaterThanOrEqual[T any](column schema.Column, value T) query.Condition {
	return query.Gte(column, value)
}

func LessThan[T any](column schema.Column, value T) query.Condition {
	return query.Lt(column, value)
}

func LessThanOrEqual[T any](column schema.Column, value T) query.Condition {
	return query.Lte(column, value)
}

func Like(column schema.Column, pattern string) query.Condition {
	return query.Like(column, pattern)
}

func In[T any](column schema.Column, values ...T) query.Condition {
	return query.In(column, values...)
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

// LimitClause represents a LIMIT clause in SQLite
type LimitClause struct {
	*query.LimitClause
}

func (l *LimitClause) ApplySelect(stmt *SelectStmt) {
	stmt.limit = l
}

// Limit creates a LIMIT clause
func Limit(limit int) *LimitClause {
	return &LimitClause{
		LimitClause: &query.LimitClause{
			Limit: limit,
		},
	}
}

// OffsetClause represents an OFFSET clause in SQLite
type OffsetClause struct {
	*query.OffsetClause
}

func (o *OffsetClause) ApplySelect(stmt *SelectStmt) {
	stmt.offset = o
}

// Offset creates an OFFSET clause
func Offset(offset int) *OffsetClause {
	return &OffsetClause{
		OffsetClause: &query.OffsetClause{
			Offset: offset,
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
