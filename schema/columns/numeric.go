package columns

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"io"

	"github.com/gogo-framework/db/dialect"
	"github.com/gogo-framework/db/query"
	"github.com/gogo-framework/db/schema"
	"golang.org/x/exp/constraints"
)

// Numeric represents a generic numeric column that can store different numeric values
type Numeric[T constraints.Ordered] struct {
	table *schema.Table
	name  string
	value sql.Null[T]
}

// GetTable returns the table this column belongs to
func (n *Numeric[T]) GetTable() *schema.Table {
	return n.table
}

// SetTable sets the table this column belongs to
func (n *Numeric[T]) SetTable(table *schema.Table) {
	n.table = table
}

// GetName returns the name of this column
func (n *Numeric[T]) GetName() string {
	return n.name
}

// SetName sets the name of this column
func (n *Numeric[T]) SetName(name string) {
	n.name = name
}

// GetType returns the SQL type of this column
// This should be implemented by the dialect-specific type
func (n *Numeric[T]) GetType() string {
	panic("GetType must be implemented by dialect-specific type")
}

// Scan implements the sql.Scanner interface
func (n *Numeric[T]) Scan(value any) error {
	return n.value.Scan(value)
}

// Value implements the driver.Valuer interface
func (n *Numeric[T]) Value() (driver.Value, error) {
	return n.value.Value()
}

// Get returns the current value of the column
func (n *Numeric[T]) Get() T {
	return n.value.V
}

// Set sets the value of the column
func (n *Numeric[T]) Set(value T) {
	n.value.V = value
	n.value.Valid = true
}

// Valid returns whether the value is valid (not NULL)
func (n *Numeric[T]) Valid() bool {
	return n.value.Valid
}

// WriteSql writes the SQL representation of this column
func (n *Numeric[T]) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	if n.table != nil {
		prefix := n.table.Name
		if n.table.Alias != "" {
			prefix = n.table.Alias
		}
		w.Write([]byte(prefix + "."))
	}
	w.Write([]byte(n.name))
	return nil, nil
}

// Eq creates an equality condition
func (n *Numeric[T]) Eq(value T) query.Condition {
	return query.Eq(n, value)
}

// Neq creates an inequality condition
func (n *Numeric[T]) Neq(value T) query.Condition {
	return query.Neq(n, value)
}

// Gt creates a greater than condition
func (n *Numeric[T]) Gt(value T) query.Condition {
	return query.Gt(n, value)
}

// Gte creates a greater than or equal condition
func (n *Numeric[T]) Gte(value T) query.Condition {
	return query.Gte(n, value)
}

// Lt creates a less than condition
func (n *Numeric[T]) Lt(value T) query.Condition {
	return query.Lt(n, value)
}

// Lte creates a less than or equal condition
func (n *Numeric[T]) Lte(value T) query.Condition {
	return query.Lte(n, value)
}

// In creates an IN condition
func (n *Numeric[T]) In(values ...T) query.Condition {
	return query.In(n, values)
}

// NotIn creates a NOT IN condition
func (n *Numeric[T]) NotIn(values ...T) query.Condition {
	return query.NotIn(n, values)
}

// IsNull creates an IS NULL condition
func (n *Numeric[T]) IsNull() query.Condition {
	return query.IsNull(n)
}

// IsNotNull creates an IS NOT NULL condition
func (n *Numeric[T]) IsNotNull() query.Condition {
	return query.IsNotNull(n)
}
