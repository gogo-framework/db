package schema

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"io"
	"log/slog"
	"strings"

	"github.com/gogo-framework/db/dialect"
)

// ColumnSchema represents the schema of a column in the database.
// As of now, it only contains the name of the column. However, in the future,
// it may contain additional information such as data type, constraints, etc
// so that it can also be used for automatic schema migrations.
type ColumnSchema struct {
	name string
}

func (cs *ColumnSchema) GetName() string {
	return cs.name
}

// The column interface is used within queries and is also the return type when mapping a row.
type Column interface {
	// sql.Scanner must be implemented so that results can be scanned back into the column.
	sql.Scanner
	// driver.Valuer must be implemented so that the column can be used as a parameter in a query.
	driver.Valuer
	// GetColumnSchema returns the schema of the column.
	GetColumnSchema() *ColumnSchema
	// SetColumnSchema sets the schema of the column.
	SetColumnSchema(*ColumnSchema)
	// GetTableSchema returns the schema of the table that the column belongs to.
	GetTableSchema() *TableSchema
	// SetTableSchema sets the schema of the table that the column belongs to.
	SetTableSchema(*TableSchema)
	// GetTable returns the table that the column belongs to. This is needed for queries and mapping back.
	GetTable() Table
	// SetTable sets the table that the column belongs to.
	SetTable(Table)
	// WriteSql writes to the io.Writer of the query builder.
	// It must be implemented so that columns can directly be part of a query.
	WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error)
	// GetAlias returns the alias of the column.
	GetAlias() string
	// SetAlias sets the alias of the column.
	SetAlias(string)
}

// BaseColumn is a base implementation of the Column interface.
type BaseColumn[T any] struct {
	// tableSchema is the schema of the table that the column belongs to.
	tableSchema *TableSchema
	// table is the table type that the column belongs to.
	table Table
	// columnSchema is the schema of the column.
	columnSchema *ColumnSchema
	// alias is the alias of the column.
	alias string
	// value is used to store the value of column from the row.
	value sql.Null[T]
	// scanned is used to track if the column has been scanned from the row.
	scanned bool
}

func (bc *BaseColumn[T]) Scan(value any) error {
	return bc.value.Scan(value)
}

func (bc *BaseColumn[T]) Value() (driver.Value, error) {
	return bc.value.Value()
}

func (bc *BaseColumn[T]) GetColumnSchema() *ColumnSchema {
	return bc.columnSchema
}

func (bc *BaseColumn[T]) SetColumnSchema(cs *ColumnSchema) {
	bc.columnSchema = cs
}

func (bc *BaseColumn[T]) GetTableSchema() *TableSchema {
	return bc.tableSchema
}

func (bc *BaseColumn[T]) SetTableSchema(ts *TableSchema) {
	bc.tableSchema = ts
}

func (bc *BaseColumn[T]) GetTable() Table {
	return bc.table
}

func (bc *BaseColumn[T]) SetTable(t Table) {
	bc.table = t
}

func (bc *BaseColumn[T]) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	var sql strings.Builder

	if bc.table != nil {
		prefix := bc.tableSchema.name
		if bc.table.GetAlias() != "" {
			prefix = bc.table.GetAlias()
		}
		sql.WriteString(d.QuoteIdentifier(prefix) + ".")
	}

	sql.WriteString(d.QuoteIdentifier(bc.columnSchema.name))

	if bc.alias != "" {
		sql.WriteString(" AS " + d.QuoteIdentifier(bc.alias))
	}

	_, err := io.WriteString(w, sql.String())
	return nil, err
}

func (bc *BaseColumn[T]) GetAlias() string {
	return bc.alias
}

func (bc *BaseColumn[T]) SetAlias(alias string) {
	bc.alias = alias
}

func (bc *BaseColumn[T]) IsScanned() bool {
	return bc.scanned
}

func (bc *BaseColumn[T]) SetScanned(scanned bool) {
	bc.scanned = scanned
}

func (bc *BaseColumn[T]) Get() T {
	if !bc.scanned {
		slog.Warn("The column has not been scanned, and thus the value might not be accurate.", "column", bc.columnSchema.name)
	}
	return bc.value.V
}

func (bc *BaseColumn[T]) Set(value T) {
	bc.value.V = value
	bc.scanned = true
}

func (bc *BaseColumn[T]) IsValid() bool {
	return bc.value.Valid
}
