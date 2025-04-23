package sqlite

import (
	"github.com/gogo-framework/db/internal/query"
	"github.com/gogo-framework/db/internal/schema"
)

// Aggregation represents a SQLite-specific aggregation
type Aggregation struct {
	*query.Aggregation
}

// ApplySelect implements the SelectPart interface
func (a *Aggregation) ApplySelect(stmt *SelectStmt) {
	if stmt.Columns == nil {
		stmt.Columns = &SelectClause{
			SelectClause: &query.SelectClause{},
		}
	}
	stmt.Columns.Columns = append(stmt.Columns.Columns, a.Aggregation)
}

// As creates an alias for the aggregation
func (a *Aggregation) As(name string) *Aggregation {
	a.Alias = name
	return a
}

// Avg creates an AVG aggregation for SQLite
// This is a wrapper around query.Avg that ensures type safety for SQLite
func Avg(column schema.Column, resultPtr *Float) *Aggregation {
	return &Aggregation{
		Aggregation: query.Avg(column, resultPtr),
	}
}

// Count creates a COUNT aggregation for SQLite
func Count(column schema.Column, resultPtr *Integer) *Aggregation {
	return &Aggregation{
		Aggregation: query.Count(column, resultPtr),
	}
}

// CountDistinct creates a COUNT(DISTINCT) aggregation for SQLite
func CountDistinct(column schema.Column, resultPtr *Integer) *Aggregation {
	return &Aggregation{
		Aggregation: query.CountDistinct(column, resultPtr),
	}
}

// CountAll creates a COUNT(*) aggregation for SQLite
func CountAll(resultPtr *Integer) *Aggregation {
	return &Aggregation{
		Aggregation: query.CountAll(resultPtr),
	}
}

// Sum creates a SUM aggregation for SQLite
func Sum(column schema.Column, resultPtr *Float) *Aggregation {
	return &Aggregation{
		Aggregation: query.Sum(column, resultPtr),
	}
}

// Min creates a MIN aggregation for SQLite
func Min(column schema.Column, resultPtr schema.Column) *Aggregation {
	return &Aggregation{
		Aggregation: query.Min(column, resultPtr),
	}
}

// Max creates a MAX aggregation for SQLite
func Max(column schema.Column, resultPtr schema.Column) *Aggregation {
	return &Aggregation{
		Aggregation: query.Max(column, resultPtr),
	}
}

// GroupConcat creates a GROUP_CONCAT aggregation for SQLite
func GroupConcat(column schema.Column, resultPtr *Text) *Aggregation {
	return &Aggregation{
		Aggregation: query.GroupConcat(column, resultPtr),
	}
}

// Total creates a TOTAL aggregation for SQLite (SQLite-specific)
func Total(column schema.Column, resultPtr *Float) *Aggregation {
	return &Aggregation{
		Aggregation: query.Total(column, resultPtr),
	}
}
