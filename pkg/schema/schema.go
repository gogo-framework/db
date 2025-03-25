package schema

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type Column interface {
	sql.Scanner
	driver.Valuer
	GetTable() *Table
	SetTable(*Table)
	GetName() string
	SetName(string)
	GetType() string
}

type Tabler interface {
	Table() *Table
}

type Table struct {
	Name    string
	Schema  string
	columns map[string]Column
}

// A helper method to correctly initialize a table.
func NewTable[T any]() *T {
	table := new(T)
	if t, ok := any(table).(Tabler); ok {
		t.Table()
	} else {
		panic(fmt.Sprintf("Type %T does not implement Tabler interface", table))
	}
	return table
}

func (t *Table) RegisterColumn(name string, col Column) {
	col.SetName(name)
	col.SetTable(t)
	if t.columns == nil {
		t.columns = make(map[string]Column)
	}
	t.columns[col.GetName()] = col
}

func (t *Table) GetColumn(name string) (Column, bool) {
	col, ok := t.columns[name]
	return col, ok
}

func (t *Table) GetColumns() map[string]Column {
	return t.columns
}
