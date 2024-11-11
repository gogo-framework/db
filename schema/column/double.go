package column

import (
	"database/sql"

	"github.com/gogo-framework/db/query/expression"
)

// Double is similar to Float but communicates to the DB that we want double precision
type Double struct {
	Name     string
	value    sql.NullFloat64 // Go doesn't differentiate between float and double internally
	modified bool
}

// Value operations
func (c *Double) Set(value float64) {
	c.value = sql.NullFloat64{Float64: value, Valid: true}
	c.modified = true
}

func (c *Double) SetNull() {
	c.value = sql.NullFloat64{Valid: false}
	c.modified = true
}

func (c Double) Get() (float64, bool) {
	return c.value.Float64, c.value.Valid
}

func (c Double) IsModified() bool {
	return c.modified
}

// Comparison operations
func (c Double) Eq(value float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "=",
		Right: value,
	}
}

func (c Double) NotEq(value float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "!=",
		Right: value,
	}
}

func (c Double) Gt(value float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    ">",
		Right: value,
	}
}

func (c Double) Gte(value float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    ">=",
		Right: value,
	}
}

func (c Double) Lt(value float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "<",
		Right: value,
	}
}

func (c Double) Lte(value float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "<=",
		Right: value,
	}
}

func (c Double) Between(start, end float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "BETWEEN",
		Right: []float64{start, end},
	}
}

func (c Double) In(values ...float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IN",
		Right: values,
	}
}

func (c Double) NotIn(values ...float64) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "NOT IN",
		Right: values,
	}
}

func (c Double) IsNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NULL",
		Right: nil,
	}
}

func (c Double) IsNotNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NOT NULL",
		Right: nil,
	}
}
