package schema

import (
	"context"
	"io"

	"github.com/gogo-framework/db/dialect"
)

// TableSchema represents the schema of a table in the database.
type TableSchema struct {
	// The name of the table.
	name string
	// The schema of the table.
	schema string
	// The columns of the table.
	columns []Column
}

func (ts *TableSchema) GetName() string {
	return ts.name
}

func (ts *TableSchema) GetSchema() string {
	return ts.schema
}

func (ts *TableSchema) GetColumns() []Column {
	return ts.columns
}

// The Table interface is used within queries and is also the return type when mapping a row.
type Table interface {
	// GetTableSchema returns the schema of the table, the implementor should ensure that this happens only once.
	// This means storing the TableSchema in the table struct, and return this if it's not nil.
	GetTableSchema() *TableSchema
	// GetAlias returns the alias of the table if it was aliased.
	GetAlias() string
	// SetAlias sets the alias of the table.
	SetAlias(alias string)
	// WriteSql writes to the io.Writer of the query builder.
	// It must be implemented so that tables can directly be part of a query.
	WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error)
}

type TableConfigurer interface {
	ConfigureSchema(schema *TableSchema)
}

// BaseTable is an incomplete implementation of the Table interface.
// It implements most of the Table interface, except the ConfigureSchema method.
// This type can (should) be embedded in a table type to avoid a lot of boilerplate.
type BaseTable struct {
	TableConfigurer
	schema *TableSchema
	alias  string
}

func (t *BaseTable) GetTableSchema() *TableSchema {
	if t.schema == nil {
		ts := new(TableSchema)
		t.ConfigureSchema(ts)
		t.schema = ts
	}
	return t.schema
}

func (t *BaseTable) GetAlias() string {
	return t.alias
}

func (t *BaseTable) SetAlias(alias string) {
	t.alias = alias
}
