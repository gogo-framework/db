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

// BinaryConstraint defines the constraints for binary types
type BinaryConstraint interface {
	~[]byte
}

// Binary represents a generic binary column that can store different binary values
type Binary[T BinaryConstraint] struct {
	table *schema.Table
	name  string
	value sql.Null[T]
}

// GetTable returns the table this column belongs to
func (b *Binary[T]) GetTable() *schema.Table {
	return b.table
}

// SetTable sets the table this column belongs to
func (b *Binary[T]) SetTable(table *schema.Table) {
	b.table = table
}

// GetName returns the name of this column
func (b *Binary[T]) GetName() string {
	return b.name
}

// SetName sets the name of this column
func (b *Binary[T]) SetName(name string) {
	b.name = name
}

// Scan implements the sql.Scanner interface
func (b *Binary[T]) Scan(value any) error {
	return b.value.Scan(value)
}

// Value implements the driver.Valuer interface
func (b *Binary[T]) Value() (driver.Value, error) {
	return b.value.Value()
}

// Get returns the current value of the column
func (b *Binary[T]) Get() T {
	return b.value.V
}

// Set sets the value of the column
func (b *Binary[T]) Set(value T) {
	b.value.V = value
	b.value.Valid = true
}

// Valid returns whether the value is valid (not NULL)
func (b *Binary[T]) Valid() bool {
	return b.value.Valid
}

// WriteSql writes the SQL representation of this column
func (b *Binary[T]) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	if b.table != nil {
		prefix := b.table.Name
		if b.table.Alias != "" {
			prefix = b.table.Alias
		}
		if _, err := w.Write([]byte(d.QuoteIdentifier(prefix) + ".")); err != nil {
			return nil, err
		}
	}
	if _, err := w.Write([]byte(d.QuoteIdentifier(b.name))); err != nil {
		return nil, err
	}
	return nil, nil
}

// Eq creates an equality condition
func (b *Binary[T]) Eq(value T) query.Condition {
	return query.Eq(b, value)
}

// Neq creates an inequality condition
func (b *Binary[T]) Neq(value T) query.Condition {
	return query.Neq(b, value)
}

// In creates an IN condition
func (b *Binary[T]) In(values ...T) query.Condition {
	return query.In(b, values)
}

// NotIn creates a NOT IN condition
func (b *Binary[T]) NotIn(values ...T) query.Condition {
	return query.NotIn(b, values)
}

// IsNull creates an IS NULL condition
func (b *Binary[T]) IsNull() query.Condition {
	return query.IsNull(b)
}

// IsNotNull creates an IS NOT NULL condition
func (b *Binary[T]) IsNotNull() query.Condition {
	return query.IsNotNull(b)
}
