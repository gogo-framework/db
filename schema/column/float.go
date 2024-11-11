package column

import (
	"database/sql"

	"github.com/gogo-framework/db/query/expression"
)

type Float struct {
	Name     string
	value    sql.NullFloat64
	modified bool
}

// Value operations
func (c *Float) Set(value float64) {
	c.value = sql.NullFloat64{Float64: value, Valid: true}
	c.modified = true
}

func (c *Float) SetNull() {
	c.value = sql.NullFloat64{Valid: false}
	c.modified = true
}

func (c Float) Get() (float64, bool) {
	return c.value.Float64, c.value.Valid
}

func (c Float) IsModified() bool {
	return c.modified
}

// Comparison operations
func (c Float) Eq(value float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "=",
		Right: value,
	}
}

func (c Float) NotEq(value float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "!=",
		Right: value,
	}
}

func (c Float) Gt(value float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    ">",
		Right: value,
	}
}

func (c Float) Gte(value float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    ">=",
		Right: value,
	}
}

func (c Float) Lt(value float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "<",
		Right: value,
	}
}

func (c Float) Lte(value float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "<=",
		Right: value,
	}
}

func (c Float) Between(start, end float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "BETWEEN",
		Right: []float64{start, end},
	}
}

func (c Float) In(values ...float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IN",
		Right: values,
	}
}

func (c Float) NotIn(values ...float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "NOT IN",
		Right: values,
	}
}

func (c Float) IsNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NULL",
		Right: nil,
	}
}

func (c Float) IsNotNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NOT NULL",
		Right: nil,
	}
}
