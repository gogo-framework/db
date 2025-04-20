package query

import (
	"context"
	"database/sql/driver"
	"fmt"
	"io"

	"github.com/gogo-framework/db/dialect"
	"github.com/gogo-framework/db/schema"
)

// Aggregation represents a SQL aggregation function
type Aggregation struct {
	Function  string
	Column    schema.Column
	Alias     string
	table     *schema.Table
	name      string
	resultPtr schema.Column
}

// GetTable returns the table this column belongs to
func (a *Aggregation) GetTable() *schema.Table {
	return a.table
}

// SetTable sets the table this column belongs to
func (a *Aggregation) SetTable(table *schema.Table) {
	a.table = table
}

// GetName returns the name of this column
func (a *Aggregation) GetName() string {
	if a.Alias != "" {
		return a.Alias
	}
	return a.name
}

// SetName sets the name of this column
func (a *Aggregation) SetName(name string) {
	a.name = name
}

// GetType returns the SQL type of this column
func (a *Aggregation) GetType() string {
	if a.resultPtr != nil {
		return a.resultPtr.GetType()
	}
	return a.Column.GetType()
}

// Scan implements the sql.Scanner interface
func (a *Aggregation) Scan(value any) error {
	if a.resultPtr != nil {
		return a.resultPtr.Scan(value)
	}
	return a.Column.Scan(value)
}

// Value implements the driver.Valuer interface
func (a *Aggregation) Value() (driver.Value, error) {
	if a.resultPtr != nil {
		return a.resultPtr.Value()
	}
	return a.Column.Value()
}

// WriteSql implements the Expression interface
func (a *Aggregation) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	// Write the function name and opening parenthesis
	w.Write([]byte(a.Function))
	w.Write([]byte("("))

	// Write the column expression
	args, err := a.Column.WriteSql(ctx, w, d, argPos)
	if err != nil {
		return nil, fmt.Errorf("error writing aggregation column: %w", err)
	}

	// Write closing parenthesis
	w.Write([]byte(")"))

	// Write alias if provided
	if a.Alias != "" {
		w.Write([]byte(" AS "))
		w.Write([]byte(d.QuoteIdentifier(a.Alias)))
	}

	return args, nil
}

// Avg creates an AVG aggregation
func Avg(column schema.Column, resultPtr schema.Column) *Aggregation {
	return &Aggregation{
		Function:  "AVG",
		Column:    column,
		Alias:     fmt.Sprintf("avg_%s", column.GetName()),
		resultPtr: resultPtr,
	}
}

// Count creates a COUNT aggregation
func Count(column schema.Column, resultPtr schema.Column) *Aggregation {
	return &Aggregation{
		Function:  "COUNT",
		Column:    column,
		Alias:     fmt.Sprintf("count_%s", column.GetName()),
		resultPtr: resultPtr,
	}
}

// CountDistinct creates a COUNT(DISTINCT) aggregation
func CountDistinct(column schema.Column, resultPtr schema.Column) *Aggregation {
	return &Aggregation{
		Function:  "COUNT(DISTINCT ",
		Column:    column,
		Alias:     fmt.Sprintf("count_distinct_%s", column.GetName()),
		resultPtr: resultPtr,
	}
}

// Sum creates a SUM aggregation
func Sum(column schema.Column, resultPtr schema.Column) *Aggregation {
	return &Aggregation{
		Function:  "SUM",
		Column:    column,
		Alias:     fmt.Sprintf("sum_%s", column.GetName()),
		resultPtr: resultPtr,
	}
}

// Min creates a MIN aggregation
func Min(column schema.Column, resultPtr schema.Column) *Aggregation {
	return &Aggregation{
		Function:  "MIN",
		Column:    column,
		Alias:     fmt.Sprintf("min_%s", column.GetName()),
		resultPtr: resultPtr,
	}
}

// Max creates a MAX aggregation
func Max(column schema.Column, resultPtr schema.Column) *Aggregation {
	return &Aggregation{
		Function:  "MAX",
		Column:    column,
		Alias:     fmt.Sprintf("max_%s", column.GetName()),
		resultPtr: resultPtr,
	}
}

// GroupConcat creates a GROUP_CONCAT aggregation
func GroupConcat(column schema.Column, resultPtr schema.Column) *Aggregation {
	return &Aggregation{
		Function:  "GROUP_CONCAT",
		Column:    column,
		Alias:     fmt.Sprintf("group_concat_%s", column.GetName()),
		resultPtr: resultPtr,
	}
}

// Total creates a TOTAL aggregation (SQLite-specific)
func Total(column schema.Column, resultPtr schema.Column) *Aggregation {
	return &Aggregation{
		Function:  "TOTAL",
		Column:    column,
		Alias:     fmt.Sprintf("total_%s", column.GetName()),
		resultPtr: resultPtr,
	}
}

// StarColumn represents a * column in SQL
type StarColumn struct{}

func (s *StarColumn) GetTable() *schema.Table {
	return nil
}

func (s *StarColumn) SetTable(*schema.Table) {}

func (s *StarColumn) GetName() string {
	return "*"
}

func (s *StarColumn) SetName(string) {}

func (s *StarColumn) GetType() string {
	return ""
}

func (s *StarColumn) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	w.Write([]byte("*"))
	return nil, nil
}

func (s *StarColumn) Scan(value any) error {
	return nil
}

func (s *StarColumn) Value() (driver.Value, error) {
	return nil, nil
}

// CountAll creates a COUNT(*) aggregation
func CountAll(resultPtr schema.Column) *Aggregation {
	return &Aggregation{
		Function:  "COUNT",
		Column:    &StarColumn{},
		Alias:     "count_all",
		resultPtr: resultPtr,
	}
}
