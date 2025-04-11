package sqlite

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"io"

	"github.com/gogo-framework/db/dialect"
	"github.com/gogo-framework/db/query"
	"github.com/gogo-framework/db/schema"
)

type Column[T any] struct {
	table *schema.Table
	name  string
	value T
}

type Text struct {
	Column[sql.NullString]
}

// Implement Column interface
func (t *Text) GetTable() *schema.Table      { return t.table }
func (t *Text) SetTable(table *schema.Table) { t.table = table }
func (t *Text) GetName() string              { return t.name }
func (t *Text) SetName(name string)          { t.name = name }
func (t *Text) GetType() string              { return "TEXT" }

// Implement sql.Scanner and driver.Valuer interfaces
func (t *Text) Scan(value any) error         { return t.value.Scan(value) }
func (t *Text) Value() (driver.Value, error) { return t.value.Value() }

// Value accessors
func (t *Text) Get() string { return t.value.String }
func (t *Text) Valid() bool { return t.value.Valid }

// Implement SqlWriter interface for conditions
func (t *Text) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	if t.table != nil {
		prefix := t.table.Name
		if t.table.Alias != "" {
			prefix = t.table.Alias
		}
		w.Write([]byte(prefix + "."))
	}
	w.Write([]byte(t.name))
	return nil, nil
}

// Condition methods
func (t *Text) Eq(value any) query.Condition {
	return Equal(t, value)
}

func (t *Text) Neq(value any) query.Condition {
	return NotEqual(t, value)
}

func (t *Text) Gt(value any) query.Condition {
	return GreaterThan(t, value)
}

func (t *Text) Gte(value any) query.Condition {
	return GreaterThanOrEqual(t, value)
}

func (t *Text) Lt(value any) query.Condition {
	return LessThan(t, value)
}

func (t *Text) Lte(value any) query.Condition {
	return LessThanOrEqual(t, value)
}

func (t *Text) Like(pattern string) query.Condition {
	return Like(t, pattern)
}

func (t *Text) In(values ...any) query.Condition {
	return In(t, values...)
}

// Implement SelectPart interface
func (t *Text) ApplySelect(stmt *query.SelectStmt) {
	stmt.Columns = append(stmt.Columns, t)
}
