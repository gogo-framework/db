package query

import (
	"context"
	"database/sql/driver"
	"fmt"
	"io"

	"github.com/gogo-framework/db/dialect"
)

// Literal represents a SQL literal value
type Literal[T any] struct {
	val T
}

// NewLiteral creates a new Literal with the given value
func NewLiteral[T any](value T) *Literal[T] {
	return &Literal[T]{val: value}
}

// WriteSql implements the SqlWriter interface
func (l *Literal[T]) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	w.Write([]byte(d.Placeholder(argPos)))
	return []any{l.val}, nil
}

// Value implements the driver.Valuer interface
func (l *Literal[T]) Value() (driver.Value, error) {
	return l.val, nil
}

// Scan implements the sql.Scanner interface
func (l *Literal[T]) Scan(value any) error {
	if value == nil {
		return fmt.Errorf("cannot scan nil into Literal")
	}

	// Type assertion to handle different types
	switch v := value.(type) {
	case T:
		l.val = v
	default:
		return fmt.Errorf("cannot scan %T into Literal[%T]", value, l.val)
	}

	return nil
}
