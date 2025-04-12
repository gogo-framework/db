package sqlite

import (
	"github.com/gogo-framework/db/query"
	"github.com/gogo-framework/db/schema/columns"
)

// Text represents a SQLite TEXT column
type Text struct {
	columns.Text[string]
}

// GetType returns the SQL type for this column
func (t *Text) GetType() string {
	return "TEXT"
}

// ApplySelect implements the SelectPart interface
func (t *Text) ApplySelect(stmt *SelectStmt) {
	if stmt.columns == nil {
		stmt.columns = &SelectClause{
			SelectClause: &query.SelectClause{},
		}
	}
	stmt.columns.Columns = append(stmt.columns.Columns, t)
}

// Integer represents a SQLite INTEGER column
type Integer struct {
	columns.Numeric[int64]
}

// GetType returns the SQL type for this column
func (i *Integer) GetType() string {
	return "INTEGER"
}

// ApplySelect implements the SelectPart interface
func (i *Integer) ApplySelect(stmt *SelectStmt) {
	if stmt.columns == nil {
		stmt.columns = &SelectClause{
			SelectClause: &query.SelectClause{},
		}
	}
	stmt.columns.Columns = append(stmt.columns.Columns, i)
}

// Float represents a SQLite REAL column
type Float struct {
	columns.Numeric[float64]
}

// GetType returns the SQL type for this column
func (f *Float) GetType() string {
	return "REAL"
}

// ApplySelect implements the SelectPart interface
func (f *Float) ApplySelect(stmt *SelectStmt) {
	if stmt.columns == nil {
		stmt.columns = &SelectClause{
			SelectClause: &query.SelectClause{},
		}
	}
	stmt.columns.Columns = append(stmt.columns.Columns, f)
}

// Blob represents a SQLite BLOB column
type Blob struct {
	columns.Binary[[]byte]
}

// GetType returns the SQL type for this column
func (b *Blob) GetType() string {
	return "BLOB"
}

// ApplySelect implements the SelectPart interface
func (b *Blob) ApplySelect(stmt *SelectStmt) {
	if stmt.columns == nil {
		stmt.columns = &SelectClause{
			SelectClause: &query.SelectClause{},
		}
	}
	stmt.columns.Columns = append(stmt.columns.Columns, b)
}
