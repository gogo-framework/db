package sqlite

import (
	"database/sql"
	"database/sql/driver"

	"github.com/gogo-framework/db/pkg/schema"
)

type Text struct {
	table *schema.Table
	name  string
	value sql.NullString
}

func (t *Text) GetTable() *schema.Table      { return t.table }
func (t *Text) SetTable(table *schema.Table) { t.table = table }
func (t *Text) GetName() string              { return t.name }
func (t *Text) SetName(name string)          { t.name = name }
func (t *Text) GetType() string              { return "TEXT" }
func (t *Text) Scan(value any) error         { return t.value.Scan(value) }
func (t *Text) Value() (driver.Value, error) { return t.value.Value() }
func (t *Text) String() string               { return t.value.String }
