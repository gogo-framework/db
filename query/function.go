package query

import (
	"context"
	"fmt"
	"io"

	"github.com/gogo-framework/db/dialect"
	"github.com/gogo-framework/db/schema"
)

// Function represents a SQL function call
type Function struct {
	Name      string
	Arguments []schema.Column
	Result    schema.Column
}

// WriteSql implements the SqlWriter interface
func (f *Function) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	// Write the function name and opening parenthesis
	w.Write([]byte(f.Name))
	w.Write([]byte("("))

	// Write the arguments
	var allArgs []any
	for i, arg := range f.Arguments {
		if i > 0 {
			w.Write([]byte(", "))
		}
		args, err := arg.WriteSql(ctx, w, d, argPos+len(allArgs))
		if err != nil {
			return nil, fmt.Errorf("error writing function argument: %w", err)
		}
		allArgs = append(allArgs, args...)
	}

	// Write closing parenthesis
	w.Write([]byte(")"))

	// Write alias if we have a result column
	if f.Result != nil {
		w.Write([]byte(" AS "))
		w.Write([]byte(d.QuoteIdentifier(f.Result.GetName())))
	}

	return allArgs, nil
}

// String functions

// Upper creates an UPPER function
func Upper(column schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "UPPER",
		Arguments: []schema.Column{column},
		Result:    resultPtr,
	}
}

// Lower creates a LOWER function
func Lower(column schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "LOWER",
		Arguments: []schema.Column{column},
		Result:    resultPtr,
	}
}

// Trim creates a TRIM function
func Trim(column schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "TRIM",
		Arguments: []schema.Column{column},
		Result:    resultPtr,
	}
}

// Substr creates a SUBSTR function
func Substr(column schema.Column, start, length schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "SUBSTR",
		Arguments: []schema.Column{column, start, length},
		Result:    resultPtr,
	}
}

// Numeric functions

// Abs creates an ABS function
func Abs(column schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "ABS",
		Arguments: []schema.Column{column},
		Result:    resultPtr,
	}
}

// Round creates a ROUND function
func Round(column schema.Column, decimals schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "ROUND",
		Arguments: []schema.Column{column, decimals},
		Result:    resultPtr,
	}
}

// Date/Time functions

// Date creates a DATE function
func Date(column schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "DATE",
		Arguments: []schema.Column{column},
		Result:    resultPtr,
	}
}

// Time creates a TIME function
func Time(column schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "TIME",
		Arguments: []schema.Column{column},
		Result:    resultPtr,
	}
}

// Datetime creates a DATETIME function
func Datetime(column schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "DATETIME",
		Arguments: []schema.Column{column},
		Result:    resultPtr,
	}
}

// Additional string functions

// Instr creates an INSTR function
func Instr(haystack, needle schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "INSTR",
		Arguments: []schema.Column{haystack, needle},
		Result:    resultPtr,
	}
}

// Hex creates a HEX function
func Hex(column schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "HEX",
		Arguments: []schema.Column{column},
		Result:    resultPtr,
	}
}

// Quote creates a QUOTE function
func Quote(column schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "QUOTE",
		Arguments: []schema.Column{column},
		Result:    resultPtr,
	}
}

// Additional numeric functions

// Ceil creates a CEIL function
func Ceil(column schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "CEIL",
		Arguments: []schema.Column{column},
		Result:    resultPtr,
	}
}

// Floor creates a FLOOR function
func Floor(column schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "FLOOR",
		Arguments: []schema.Column{column},
		Result:    resultPtr,
	}
}

// Mod creates a MOD function
func Mod(dividend, divisor schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "MOD",
		Arguments: []schema.Column{dividend, divisor},
		Result:    resultPtr,
	}
}

// Additional date/time functions

// JulianDay creates a JULIANDAY function
func JulianDay(column schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "JULIANDAY",
		Arguments: []schema.Column{column},
		Result:    resultPtr,
	}
}

// Strftime creates a STRFTIME function
func Strftime(format, column schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "STRFTIME",
		Arguments: []schema.Column{format, column},
		Result:    resultPtr,
	}
}

// Type conversion functions

// Cast creates a CAST function
func Cast(column schema.Column, typeName string, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "CAST",
		Arguments: []schema.Column{column},
		Result:    resultPtr,
	}
}

// Typeof creates a TYPEOF function
func Typeof(column schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "TYPEOF",
		Arguments: []schema.Column{column},
		Result:    resultPtr,
	}
}

// JSON functions

// JsonExtract creates a JSON_EXTRACT function
func JsonExtract(column, path schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "JSON_EXTRACT",
		Arguments: []schema.Column{column, path},
		Result:    resultPtr,
	}
}

// JsonArray creates a JSON_ARRAY function
func JsonArray(resultPtr schema.Column, columns ...schema.Column) *Function {
	return &Function{
		Name:      "JSON_ARRAY",
		Arguments: columns,
		Result:    resultPtr,
	}
}

// JsonObject creates a JSON_OBJECT function
func JsonObject(resultPtr schema.Column, columns ...schema.Column) *Function {
	return &Function{
		Name:      "JSON_OBJECT",
		Arguments: columns,
		Result:    resultPtr,
	}
}

// Utility functions

// Coalesce creates a COALESCE function
func Coalesce(resultPtr schema.Column, columns ...schema.Column) *Function {
	return &Function{
		Name:      "COALESCE",
		Arguments: columns,
		Result:    resultPtr,
	}
}

// IfNull creates an IFNULL function
func IfNull(column, defaultValue schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "IFNULL",
		Arguments: []schema.Column{column, defaultValue},
		Result:    resultPtr,
	}
}

// NullIf creates a NULLIF function
func NullIf(column1, column2 schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "NULLIF",
		Arguments: []schema.Column{column1, column2},
		Result:    resultPtr,
	}
}

// Random creates a RANDOM function
func Random(resultPtr schema.Column) *Function {
	return &Function{
		Name:      "RANDOM",
		Arguments: []schema.Column{},
		Result:    resultPtr,
	}
}

// Length creates a LENGTH function
func Length(column schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "LENGTH",
		Arguments: []schema.Column{column},
		Result:    resultPtr,
	}
}

// Replace creates a REPLACE function
func Replace(column, search, replace schema.Column, resultPtr schema.Column) *Function {
	return &Function{
		Name:      "REPLACE",
		Arguments: []schema.Column{column, search, replace},
		Result:    resultPtr,
	}
}
