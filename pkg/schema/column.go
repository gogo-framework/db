package schema

import (
	"database/sql"
	"database/sql/driver"
)

type ColumnValuer interface {
	sql.Scanner
	driver.Valuer
}

type ColumnDefinition struct {
	Name     string
	table    *Table
	valuePtr ColumnValuer
}

func NewColumnDefinition(name string) *ColumnDefinition {
	return &ColumnDefinition{Name: name}
}

func (c *ColumnDefinition) GetTable() *Table {
	return c.table
}

func (c *ColumnDefinition) SetTable(table *Table) {
	c.table = table
}

func (c *ColumnDefinition) GetValuePtr() ColumnValuer {
	return c.valuePtr
}

func (c *ColumnDefinition) SetValuePtr(valuePtr ColumnValuer) {
	c.valuePtr = valuePtr
}
