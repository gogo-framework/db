package column

import (
	"database/sql"

	"github.com/gogo-framework/db/query/expression"
)

type Integer struct {
	Name     string
	value    sql.NullInt64
	modified bool
}

// Value operations
func (c *Integer) Set(value int64) {
	c.value = sql.NullInt64{Int64: value, Valid: true}
	c.modified = true
}

func (c *Integer) SetNull() {
	c.value = sql.NullInt64{Valid: false}
	c.modified = true
}

func (c Integer) Get() (int64, bool) {
	return c.value.Int64, c.value.Valid
}

func (c Integer) IsModified() bool {
	return c.modified
}

// Comparison operations
func (c Integer) Eq(value int64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "=",
		Right: value,
	}
}

func (c Integer) NotEq(value int64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "!=",
		Right: value,
	}
}

func (c Integer) Gt(value int64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    ">",
		Right: value,
	}
}

func (c Integer) Gte(value int64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    ">=",
		Right: value,
	}
}

func (c Integer) Lt(value int64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "<",
		Right: value,
	}
}

func (c Integer) Lte(value int64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "<=",
		Right: value,
	}
}

func (c Integer) Between(start, end int64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "BETWEEN",
		Right: []int64{start, end},
	}
}

func (c Integer) In(values ...int64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IN",
		Right: values,
	}
}

func (c Integer) NotIn(values ...int64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "NOT IN",
		Right: values,
	}
}

func (c Integer) IsNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NULL",
		Right: nil,
	}
}

func (c Integer) IsNotNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NOT NULL",
		Right: nil,
	}
}
