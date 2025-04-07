package sqlite

import (
	"database/sql"
	"database/sql/driver"

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

// Implements the SelectPart interface.
// This way columns can directly be part of the select statement.
// And you won't need to wrap it into a Col function call.
func (t *Text) ApplySelect(stmt *SelectStmt) {
	stmt.selectClause.AppendColumns(t)
}
