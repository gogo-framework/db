package sqlite

import (
	"context"
	"fmt"
	"io"

	"github.com/gogo-framework/db/dialect"
	"github.com/gogo-framework/db/query"
	"github.com/gogo-framework/db/schema"
)

// SelectClause represents a SELECT clause
type SelectClause struct {
	Columns []query.SqlWriter
}

func (s *SelectClause) ApplySelect(stmt *SelectStmt) {
	stmt.columns = s
}

func (s *SelectClause) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	if len(s.Columns) == 0 {
		return nil, fmt.Errorf("no columns selected")
	}

	var allArgs []any
	for i, col := range s.Columns {
		if i > 0 {
			w.Write([]byte(", "))
		}
		args, err := col.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, fmt.Errorf("error writing column: %w", err)
		}
		allArgs = append(allArgs, args...)
		argPos += len(args)
	}
	return allArgs, nil
}

// Select creates a new SQLite SELECT statement
func Select(parts ...query.SqlWriter) *SelectStmt {
	stmt := &SelectStmt{}
	var columns []query.SqlWriter

	for _, part := range parts {
		if selectPart, ok := part.(SelectPart); ok {
			selectPart.ApplySelect(stmt)
		} else if part != nil {
			columns = append(columns, part)
		}
	}

	if len(columns) > 0 {
		stmt.columns = &SelectClause{
			Columns: columns,
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

func (f *FromClause) As(alias string) *FromClause {
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

func (w *WhereClause) WriteSql(ctx context.Context, writer io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	return w.WhereClause.WriteSql(ctx, writer, d, argPos)
}

// Where creates a WHERE clause
func Where(conditions ...query.Condition) SelectPart {
	return &WhereClause{
		WhereClause: &query.WhereClause{
			Conditions: conditions,
		},
	}
}

// OrderByClause represents an ORDER BY clause
type OrderByClause struct {
	Columns []query.SqlWriter
}

func (o *OrderByClause) ApplySelect(stmt *SelectStmt) {
	stmt.orderBy = o
}

func (o *OrderByClause) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	if len(o.Columns) == 0 {
		return nil, nil
	}

	var allArgs []any
	for i, col := range o.Columns {
		if i > 0 {
			w.Write([]byte(", "))
		}
		args, err := col.WriteSql(ctx, w, d, argPos)
		if err != nil {
			return nil, err
		}
		allArgs = append(allArgs, args...)
		argPos += len(args)
	}
	return allArgs, nil
}

// OrderBy creates an ORDER BY clause
func OrderBy(columns ...query.SqlWriter) SelectPart {
	return &OrderByClause{
		Columns: columns,
	}
}

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
