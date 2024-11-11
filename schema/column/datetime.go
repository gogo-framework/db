package column

import (
	"database/sql"
	"time"

	"github.com/gogo-framework/db/query/expression"
)

type DateTime struct {
	Name     string
	value    sql.NullTime
	modified bool
}

// Value operations
func (c *DateTime) Set(value time.Time) {
	c.value = sql.NullTime{Time: value, Valid: true}
	c.modified = true
}

func (c *DateTime) SetNull() {
	c.value = sql.NullTime{Valid: false}
	c.modified = true
}

func (c DateTime) Get() (time.Time, bool) {
	return c.value.Time, c.value.Valid
}

func (c DateTime) IsModified() bool {
	return c.modified
}

// DateTime comparison operations
func (c DateTime) Eq(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "=",
		Right: value,
	}
}

func (c DateTime) NotEq(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "!=",
		Right: value,
	}
}

func (c DateTime) Gt(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    ">",
		Right: value,
	}
}

func (c DateTime) Gte(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    ">=",
		Right: value,
	}
}

func (c DateTime) Lt(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "<",
		Right: value,
	}
}

func (c DateTime) Lte(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "<=",
		Right: value,
	}
}

func (c DateTime) Between(start, end time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "BETWEEN",
		Right: []time.Time{start, end},
	}
}

func (c DateTime) IsNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NULL",
		Right: nil,
	}
}

func (c DateTime) IsNotNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NOT NULL",
		Right: nil,
	}
}
