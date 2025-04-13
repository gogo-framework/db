package sqlite

import (
	"github.com/gogo-framework/db/query"
	"github.com/gogo-framework/db/schema"
)

// Function represents a SQLite-specific function
type Function struct {
	*query.Function
}

// ApplySelect implements the SelectPart interface
func (f *Function) ApplySelect(stmt *SelectStmt) {
	if stmt.Columns == nil {
		stmt.Columns = &SelectClause{
			SelectClause: &query.SelectClause{},
		}
	}
	stmt.Columns.Columns = append(stmt.Columns.Columns, f.Function)
}

// String functions

// Upper creates an UPPER function for SQLite
func Upper(column schema.Column, resultPtr *Text) *Function {
	return &Function{
		Function: query.Upper(column, resultPtr),
	}
}

// Lower creates a LOWER function for SQLite
func Lower(column schema.Column, resultPtr *Text) *Function {
	return &Function{
		Function: query.Lower(column, resultPtr),
	}
}

// Trim creates a TRIM function for SQLite
func Trim(column schema.Column, resultPtr *Text) *Function {
	return &Function{
		Function: query.Trim(column, resultPtr),
	}
}

// Substr creates a SUBSTR function for SQLite
func Substr(column schema.Column, start, length int, resultPtr *Text) *Function {
	return &Function{
		Function: query.Substr(column, start, length, resultPtr),
	}
}

// Instr creates an INSTR function for SQLite
func Instr(haystack schema.Column, needle string, resultPtr *Text) *Function {
	return &Function{
		Function: query.Instr(haystack, needle, resultPtr),
	}
}

// Hex creates a HEX function for SQLite
func Hex(column schema.Column, resultPtr *Text) *Function {
	return &Function{
		Function: query.Hex(column, resultPtr),
	}
}

// Quote creates a QUOTE function for SQLite
func Quote(column schema.Column, resultPtr *Text) *Function {
	return &Function{
		Function: query.Quote(column, resultPtr),
	}
}

// Length creates a LENGTH function for SQLite
func Length(column schema.Column, resultPtr *Integer) *Function {
	return &Function{
		Function: query.Length(column, resultPtr),
	}
}

// Replace creates a REPLACE function for SQLite
func Replace(column schema.Column, search, replace string, resultPtr *Text) *Function {
	return &Function{
		Function: query.Replace(column, search, replace, resultPtr),
	}
}

// Numeric functions

// Abs creates an ABS function for SQLite
func Abs(column schema.Column, resultPtr *Float) *Function {
	return &Function{
		Function: query.Abs(column, resultPtr),
	}
}

// Round creates a ROUND function for SQLite
func Round(column schema.Column, decimals int, resultPtr *Float) *Function {
	return &Function{
		Function: query.Round(column, decimals, resultPtr),
	}
}

// Ceil creates a CEIL function for SQLite
func Ceil(column schema.Column, resultPtr *Float) *Function {
	return &Function{
		Function: query.Ceil(column, resultPtr),
	}
}

// Floor creates a FLOOR function for SQLite
func Floor(column schema.Column, resultPtr *Float) *Function {
	return &Function{
		Function: query.Floor(column, resultPtr),
	}
}

// Mod creates a MOD function for SQLite
func Mod(dividend schema.Column, divisor any, resultPtr *Float) *Function {
	return &Function{
		Function: query.Mod(dividend, divisor, resultPtr),
	}
}

// Random creates a RANDOM function for SQLite
func Random(resultPtr *Float) *Function {
	return &Function{
		Function: query.Random(resultPtr),
	}
}

// Date/Time functions

// Date creates a DATE function for SQLite
func Date(column schema.Column, resultPtr *Text) *Function {
	return &Function{
		Function: query.Date(column, resultPtr),
	}
}

// Time creates a TIME function for SQLite
func Time(column schema.Column, resultPtr *Text) *Function {
	return &Function{
		Function: query.Time(column, resultPtr),
	}
}

// Datetime creates a DATETIME function for SQLite
func Datetime(column schema.Column, resultPtr *Text) *Function {
	return &Function{
		Function: query.Datetime(column, resultPtr),
	}
}

// JulianDay creates a JULIANDAY function for SQLite
func JulianDay(column schema.Column, resultPtr *Float) *Function {
	return &Function{
		Function: query.JulianDay(column, resultPtr),
	}
}

// Strftime creates a STRFTIME function for SQLite
func Strftime(format string, column schema.Column, resultPtr *Text) *Function {
	return &Function{
		Function: query.Strftime(format, column, resultPtr),
	}
}

// Type conversion functions

// Cast creates a CAST function for SQLite
func Cast(column schema.Column, typeName string, resultPtr schema.Column) *Function {
	return &Function{
		Function: query.Cast(column, typeName, resultPtr),
	}
}

// Typeof creates a TYPEOF function for SQLite
func Typeof(column schema.Column, resultPtr *Text) *Function {
	return &Function{
		Function: query.Typeof(column, resultPtr),
	}
}

// JSON functions

// JsonExtract creates a JSON_EXTRACT function for SQLite
func JsonExtract(column schema.Column, path string, resultPtr *Text) *Function {
	return &Function{
		Function: query.JsonExtract(column, path, resultPtr),
	}
}

// JsonArray creates a JSON_ARRAY function for SQLite
func JsonArray(resultPtr *Text, columns ...schema.Column) *Function {
	return &Function{
		Function: query.JsonArray(resultPtr, columns...),
	}
}

// JsonObject creates a JSON_OBJECT function for SQLite
func JsonObject(resultPtr *Text, columns ...schema.Column) *Function {
	return &Function{
		Function: query.JsonObject(resultPtr, columns...),
	}
}

// Utility functions

// Coalesce creates a COALESCE function for SQLite
func Coalesce(resultPtr schema.Column, columns ...schema.Column) *Function {
	return &Function{
		Function: query.Coalesce(resultPtr, columns...),
	}
}

// IfNull creates an IFNULL function for SQLite
func IfNull(column, defaultValue schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Function: query.IfNull(column, defaultValue, resultPtr),
	}
}

// NullIf creates a NULLIF function for SQLite
func NullIf(column1, column2 schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Function: query.NullIf(column1, column2, resultPtr),
	}
}
