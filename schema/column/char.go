package column

import (
	"database/sql"

	"github.com/gogo-framework/db/query/expression"
)

// Char is similar to Text but communicates to the DB that we want a fixed-length char field
type Char struct {
	Name     string
	value    sql.NullString
	modified bool
}

// Value operations
func (c *Char) Set(value string) {
	c.value = sql.NullString{String: value, Valid: true}
	c.modified = true
}

func (c *Char) SetNull() {
	c.value = sql.NullString{Valid: false}
	c.modified = true
}

func (c Char) Get() (string, bool) {
	return c.value.String, c.value.Valid
}

func (c Char) IsModified() bool {
	return c.modified
}

// Text-specific comparison operations
func (c Char) Eq(value string) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "=",
		Right: value,
	}
}

func (c Char) NotEq(value string) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "!=",
		Right: value,
	}
}

func (c Char) Like(pattern string) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "LIKE",
		Right: pattern,
	}
}

func (c Char) ILike(pattern string) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "ILIKE",
		Right: pattern,
	}
}

func (c Char) In(values ...string) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IN",
		Right: values,
	}
}

func (c Char) NotIn(values ...string) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "NOT IN",
		Right: values,
	}
}

func (c Char) IsNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NULL",
		Right: nil,
	}
}

func (c Char) IsNotNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NOT NULL",
		Right: nil,
	}
}
