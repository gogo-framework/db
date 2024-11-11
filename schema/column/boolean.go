package column

import (
	"database/sql"

	"github.com/gogo-framework/db/query/expression"
)

type Boolean struct {
	Name     string
	value    sql.NullBool
	modified bool
}

// Value operations
func (c *Boolean) Set(value bool) {
	c.value = sql.NullBool{Bool: value, Valid: true}
	c.modified = true
}

func (c *Boolean) SetNull() {
	c.value = sql.NullBool{Valid: false}
	c.modified = true
}

func (c Boolean) Get() (bool, bool) {
	return c.value.Bool, c.value.Valid
}

func (c Boolean) IsModified() bool {
	return c.modified
}

// Boolean comparison operations
func (c Boolean) Eq(value bool) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "=",
		Right: value,
	}
}

func (c Boolean) NotEq(value bool) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "!=",
		Right: value,
	}
}

func (c Boolean) IsTrue() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "=",
		Right: true,
	}
}

func (c Boolean) IsFalse() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "=",
		Right: false,
	}
}

func (c Boolean) IsNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NULL",
		Right: nil,
	}
}

func (c Boolean) IsNotNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NOT NULL",
		Right: nil,
	}
}
