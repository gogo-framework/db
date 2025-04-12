package query

import (
	"context"
	"io"

	"github.com/gogo-framework/db/dialect"
)

type Operator string

const (
	OpEqual              Operator = "="
	OpNotEqual           Operator = "!="
	OpGreaterThan        Operator = ">"
	OpGreaterThanOrEqual Operator = ">="
	OpLessThan           Operator = "<"
	OpLessThanOrEqual    Operator = "<="
	OpLike               Operator = "LIKE"
	OpNotLike            Operator = "NOT LIKE"
	OpIn                 Operator = "IN"
	OpNotIn              Operator = "NOT IN"
	OpGlob               Operator = "GLOB"
	OpMatch              Operator = "MATCH"
	OpRegexp             Operator = "REGEXP"
)

// Condition is an interface for SQL conditions
type Condition interface {
	SqlWriter
}

// BinaryCondition represents a binary operation (e.g., =, >, <)
type BinaryCondition struct {
	Left  SqlWriter
	Op    Operator
	Right SqlWriter
}

func (c *BinaryCondition) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	var args []any

	leftArgs, err := c.Left.WriteSql(ctx, w, d, argPos)
	if err != nil {
		return nil, err
	}
	args = append(args, leftArgs...)

	w.Write([]byte(" " + string(c.Op) + " "))

	rightArgs, err := c.Right.WriteSql(ctx, w, d, argPos+len(args))
	if err != nil {
		return nil, err
	}
	args = append(args, rightArgs...)

	return args, nil
}

// OrCondition represents a set of conditions joined by OR
type OrCondition struct {
	Conditions []Condition
}

func (c *OrCondition) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	if len(c.Conditions) == 0 {
		return nil, nil
	}

	var args []any
	w.Write([]byte("("))

	for i, condition := range c.Conditions {
		if i > 0 {
			w.Write([]byte(" OR "))
		}

		conditionArgs, err := condition.WriteSql(ctx, w, d, argPos+len(args))
		if err != nil {
			return nil, err
		}
		args = append(args, conditionArgs...)
	}

	w.Write([]byte(")"))
	return args, nil
}

// InCondition represents an IN clause
type InCondition struct {
	Column SqlWriter
	Values []SqlWriter
}

func (c *InCondition) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	var args []any

	// Write the column
	columnArgs, err := c.Column.WriteSql(ctx, w, d, argPos)
	if err != nil {
		return nil, err
	}
	args = append(args, columnArgs...)

	// Write IN and the values
	w.Write([]byte(" IN ("))
	for i, value := range c.Values {
		valueArgs, err := value.WriteSql(ctx, w, d, argPos+len(args))
		if err != nil {
			return nil, err
		}
		args = append(args, valueArgs...)
		if i < len(c.Values)-1 {
			w.Write([]byte(", "))
		}
	}
	w.Write([]byte(")"))

	return args, nil
}

// NotLikeCondition represents a NOT LIKE clause
type NotLikeCondition struct {
	Column  SqlWriter
	Pattern SqlWriter
}

func (c *NotLikeCondition) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	var args []any

	columnArgs, err := c.Column.WriteSql(ctx, w, d, argPos)
	if err != nil {
		return nil, err
	}
	args = append(args, columnArgs...)

	w.Write([]byte(" NOT LIKE "))

	patternArgs, err := c.Pattern.WriteSql(ctx, w, d, argPos+len(args))
	if err != nil {
		return nil, err
	}
	args = append(args, patternArgs...)

	return args, nil
}

// NotInCondition represents a NOT IN clause
type NotInCondition struct {
	Column SqlWriter
	Values []SqlWriter
}

func (c *NotInCondition) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	var args []any

	columnArgs, err := c.Column.WriteSql(ctx, w, d, argPos)
	if err != nil {
		return nil, err
	}
	args = append(args, columnArgs...)

	w.Write([]byte(" NOT IN ("))
	for i, value := range c.Values {
		valueArgs, err := value.WriteSql(ctx, w, d, argPos+len(args))
		if err != nil {
			return nil, err
		}
		args = append(args, valueArgs...)
		if i < len(c.Values)-1 {
			w.Write([]byte(", "))
		}
	}
	w.Write([]byte(")"))

	return args, nil
}

// IsNullCondition represents an IS NULL clause
type IsNullCondition struct {
	Column SqlWriter
}

func (c *IsNullCondition) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	args, err := c.Column.WriteSql(ctx, w, d, argPos)
	if err != nil {
		return nil, err
	}
	w.Write([]byte(" IS NULL"))
	return args, nil
}

// IsNotNullCondition represents an IS NOT NULL clause
type IsNotNullCondition struct {
	Column SqlWriter
}

func (c *IsNotNullCondition) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	args, err := c.Column.WriteSql(ctx, w, d, argPos)
	if err != nil {
		return nil, err
	}
	w.Write([]byte(" IS NOT NULL"))
	return args, nil
}

// Condition builder functions
func Eq[T any](column SqlWriter, value T) Condition {
	return &BinaryCondition{
		Left:  column,
		Op:    OpEqual,
		Right: NewLiteral(value),
	}
}

func Neq[T any](column SqlWriter, value T) Condition {
	return &BinaryCondition{
		Left:  column,
		Op:    OpNotEqual,
		Right: NewLiteral(value),
	}
}

func Gt[T any](column SqlWriter, value T) Condition {
	return &BinaryCondition{
		Left:  column,
		Op:    OpGreaterThan,
		Right: NewLiteral(value),
	}
}

func Gte[T any](column SqlWriter, value T) Condition {
	return &BinaryCondition{
		Left:  column,
		Op:    OpGreaterThanOrEqual,
		Right: NewLiteral(value),
	}
}

func Lt[T any](column SqlWriter, value T) Condition {
	return &BinaryCondition{
		Left:  column,
		Op:    OpLessThan,
		Right: NewLiteral(value),
	}
}

func Lte[T any](column SqlWriter, value T) Condition {
	return &BinaryCondition{
		Left:  column,
		Op:    OpLessThanOrEqual,
		Right: NewLiteral(value),
	}
}

func Like(column SqlWriter, pattern string) Condition {
	return &BinaryCondition{
		Left:  column,
		Op:    OpLike,
		Right: NewLiteral(pattern),
	}
}

func In[T any](column SqlWriter, values ...T) Condition {
	literals := make([]SqlWriter, len(values))
	for i, v := range values {
		literals[i] = NewLiteral(v)
	}
	return &InCondition{
		Column: column,
		Values: literals,
	}
}

// NotLike creates a NOT LIKE condition
func NotLike(column SqlWriter, pattern string) Condition {
	return &NotLikeCondition{
		Column:  column,
		Pattern: NewLiteral(pattern),
	}
}

// NotIn creates a NOT IN condition
func NotIn[T any](column SqlWriter, values ...T) Condition {
	literals := make([]SqlWriter, len(values))
	for i, v := range values {
		literals[i] = NewLiteral(v)
	}
	return &NotInCondition{
		Column: column,
		Values: literals,
	}
}

// IsNull creates an IS NULL condition
func IsNull(column SqlWriter) Condition {
	return &IsNullCondition{
		Column: column,
	}
}

// IsNotNull creates an IS NOT NULL condition
func IsNotNull(column SqlWriter) Condition {
	return &IsNotNullCondition{
		Column: column,
	}
}
