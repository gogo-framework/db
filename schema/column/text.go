package column

import (
	"database/sql"

	"github.com/gogo-framework/db/query/expression"
)

type Text struct {
	Name     string
	value    sql.NullString
	modified bool
}

// Value operations
func (c *Text) Set(value string) {
	c.value = sql.NullString{String: value, Valid: true}
	c.modified = true
}

func (c *Text) SetNull() {
	c.value = sql.NullString{Valid: false}
	c.modified = true
}

func (c Text) Get() (string, bool) {
	return c.value.String, c.value.Valid
}

func (c Text) IsModified() bool {
	return c.modified
}

// Text-specific comparison operations
func (c Text) Eq(value string) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "=",
		Right: value,
	}
}

func (c Text) NotEq(value string) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "!=",
		Right: value,
	}
}

func (c Text) Like(pattern string) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "LIKE",
		Right: pattern,
	}
}

func (c Text) ILike(pattern string) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "ILIKE",
		Right: pattern,
	}
}

func (c Text) In(values ...string) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IN",
		Right: values,
	}
}

func (c Text) NotIn(values ...string) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "NOT IN",
		Right: values,
	}
}

func (c Text) IsNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NULL",
		Right: nil,
	}
}

func (c Text) IsNotNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NOT NULL",
		Right: nil,
	}
}
