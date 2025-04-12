package columns

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"io"

	"github.com/gogo-framework/db/dialect"
	"github.com/gogo-framework/db/query"
	"github.com/gogo-framework/db/schema"
)

// TextConstraint defines the constraints for text types
type TextConstraint interface {
	~string | ~[]byte
}

// Text represents a generic text column that can store different text values
type Text[T TextConstraint] struct {
	table *schema.Table
	name  string
	value sql.Null[T]
}

// GetTable returns the table this column belongs to
func (t *Text[T]) GetTable() *schema.Table {
	return t.table
}

// SetTable sets the table this column belongs to
func (t *Text[T]) SetTable(table *schema.Table) {
	t.table = table
}

// GetName returns the name of this column
func (t *Text[T]) GetName() string {
	return t.name
}

// SetName sets the name of this column
func (t *Text[T]) SetName(name string) {
	t.name = name
}

// Scan implements the sql.Scanner interface
func (t *Text[T]) Scan(value any) error {
	return t.value.Scan(value)
}

// Value implements the driver.Valuer interface
func (t *Text[T]) Value() (driver.Value, error) {
	return t.value.Value()
}

// Get returns the current value of the column
func (t *Text[T]) Get() T {
	return t.value.V
}

// Set sets the value of the column
func (t *Text[T]) Set(value T) {
	t.value.V = value
	t.value.Valid = true
}

// Valid returns whether the value is valid (not NULL)
func (t *Text[T]) Valid() bool {
	return t.value.Valid
}

// WriteSql writes the SQL representation of this column
func (t *Text[T]) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	if t.table != nil {
		prefix := t.table.Name
		if t.table.Alias != "" {
			prefix = t.table.Alias
		}
		if _, err := w.Write([]byte(d.QuoteIdentifier(prefix) + ".")); err != nil {
			return nil, err
		}
	}
	if _, err := w.Write([]byte(d.QuoteIdentifier(t.name))); err != nil {
		return nil, err
	}
	return nil, nil
}

// Eq creates an equality condition
func (t *Text[T]) Eq(value T) query.Condition {
	return query.Eq(t, value)
}

// Neq creates an inequality condition
func (t *Text[T]) Neq(value T) query.Condition {
	return query.Neq(t, value)
}

// Like creates a LIKE condition
func (t *Text[T]) Like(pattern string) query.Condition {
	return query.Like(t, pattern)
}

// NotLike creates a NOT LIKE condition
func (t *Text[T]) NotLike(pattern string) query.Condition {
	return query.NotLike(t, pattern)
}

// In creates an IN condition
func (t *Text[T]) In(values ...T) query.Condition {
	return query.In(t, values)
}

// NotIn creates a NOT IN condition
func (t *Text[T]) NotIn(values ...T) query.Condition {
	return query.NotIn(t, values)
}

// IsNull creates an IS NULL condition
func (t *Text[T]) IsNull() query.Condition {
	return query.IsNull(t)
}

// IsNotNull creates an IS NOT NULL condition
func (t *Text[T]) IsNotNull() query.Condition {
	return query.IsNotNull(t)
}
