package column

import (
	"database/sql"
	"time"

	"github.com/gogo-framework/db/query/expression"
)

type Date struct {
	Name     string
	value    sql.NullTime
	modified bool
}

// Value operations
func (c *Date) Set(value time.Time) {
	c.value = sql.NullTime{Time: value, Valid: true}
	c.modified = true
}

func (c *Date) SetNull() {
	c.value = sql.NullTime{Valid: false}
	c.modified = true
}

func (c Date) Get() (time.Time, bool) {
	return c.value.Time, c.value.Valid
}

func (c Date) IsModified() bool {
	return c.modified
}

// Date comparison operations
func (c Date) Eq(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "=",
		Right: value,
	}
}

func (c Date) NotEq(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "!=",
		Right: value,
	}
}

func (c Date) Gt(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    ">",
		Right: value,
	}
}

func (c Date) Gte(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    ">=",
		Right: value,
	}
}

func (c Date) Lt(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "<",
		Right: value,
	}
}

func (c Date) Lte(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "<=",
		Right: value,
	}
}

func (c Date) Between(start, end time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "BETWEEN",
		Right: []time.Time{start, end},
	}
}

func (c Date) IsNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NULL",
		Right: nil,
	}
}

func (c Date) IsNotNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NOT NULL",
		Right: nil,
	}
}
