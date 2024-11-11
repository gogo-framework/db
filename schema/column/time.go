package column

import (
	"database/sql"
	"time"

	"github.com/gogo-framework/db/query/expression"
)

type Time struct {
	Name     string
	value    sql.NullTime
	modified bool
}

// Value operations
func (c *Time) Set(value time.Time) {
	c.value = sql.NullTime{Time: value, Valid: true}
	c.modified = true
}

func (c *Time) SetNull() {
	c.value = sql.NullTime{Valid: false}
	c.modified = true
}

func (c Time) Get() (time.Time, bool) {
	return c.value.Time, c.value.Valid
}

func (c Time) IsModified() bool {
	return c.modified
}

// Time comparison operations
func (c Time) Eq(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "=",
		Right: value,
	}
}

func (c Time) NotEq(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "!=",
		Right: value,
	}
}

func (c Time) Gt(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    ">",
		Right: value,
	}
}

func (c Time) Gte(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    ">=",
		Right: value,
	}
}

func (c Time) Lt(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "<",
		Right: value,
	}
}

func (c Time) Lte(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "<=",
		Right: value,
	}
}

func (c Time) Between(start, end time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "BETWEEN",
		Right: []time.Time{start, end},
	}
}

func (c Time) IsNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NULL",
		Right: nil,
	}
}

func (c Time) IsNotNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NOT NULL",
		Right: nil,
	}
}
