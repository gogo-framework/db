package schema

import "fmt"

type Table struct {
	Name    string
	Schema  string
	columns map[ColumnValuer]*ColumnDefinition
}

func (t *Table) RegisterColumn(valuePtr ColumnValuer, colDef *ColumnDefinition) {
	colDef.SetTable(t)
	if t.columns == nil {
		t.columns = make(map[ColumnValuer]*ColumnDefinition)
	}
	colDef.SetValuePtr(valuePtr)
	t.columns[valuePtr] = colDef
}

func (t *Table) GetColumnFromValuePtr(valuePtr ColumnValuer) (*ColumnDefinition, error) {
	config, ok := t.columns[valuePtr]
	if !ok {
		return nil, fmt.Errorf("column not found")
	}
	return config, nil
}
