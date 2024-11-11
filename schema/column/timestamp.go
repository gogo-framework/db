package column

import (
	"database/sql"
	"time"

	"github.com/gogo-framework/db/query/expression"
)

type Timestamp struct {
	Name     string
	value    sql.NullTime
	modified bool
}

// Value operations
func (c *Timestamp) Set(value time.Time) {
	c.value = sql.NullTime{Time: value, Valid: true}
	c.modified = true
}

func (c *Timestamp) SetNull() {
	c.value = sql.NullTime{Valid: false}
	c.modified = true
}

func (c Timestamp) Get() (time.Time, bool) {
	return c.value.Time, c.value.Valid
}

func (c Timestamp) IsModified() bool {
	return c.modified
}

// Timestamp comparison operations
func (c Timestamp) Eq(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "=",
		Right: value,
	}
}

func (c Timestamp) NotEq(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "!=",
		Right: value,
	}
}

func (c Timestamp) Gt(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    ">",
		Right: value,
	}
}

func (c Timestamp) Gte(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    ">=",
		Right: value,
	}
}

func (c Timestamp) Lt(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "<",
		Right: value,
	}
}

func (c Timestamp) Lte(value time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "<=",
		Right: value,
	}
}

func (c Timestamp) Between(start, end time.Time) expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "BETWEEN",
		Right: []time.Time{start, end},
	}
}

func (c Timestamp) IsNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NULL",
		Right: nil,
	}
}

func (c Timestamp) IsNotNull() expression.Comparison {
	return expression.Comparison{
		Left:  c.Name,
		Op:    "IS NOT NULL",
		Right: nil,
	}
}
