package query

import (
	"context"
	"io"

	"github.com/gogo-framework/db/dialect"
	"github.com/gogo-framework/db/schema"
)

// FromClause represents a FROM clause
type FromClause struct {
	Source schema.Tabler
	Alias  string
	Joins  []any // We have no type yet for joins.
}

func (f *FromClause) ApplySelect(stmt *SelectStmt) {
	stmt.from = f
}

func (f *FromClause) As(alias string) {
	f.Alias = alias
	if f.Source != nil {
		table := f.Source.Table()
		table.Alias = alias
	}
}

func (f *FromClause) AppendJoins(joins ...any) {
	f.Joins = append(f.Joins, joins...)
}

// WriteSql writes the FROM clause to the given writer.
func (f *FromClause) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	var args []any

	// Write the table name with proper quoting
	if _, err := w.Write([]byte(d.QuoteIdentifier(f.Source.Table().Name))); err != nil {
		return nil, err
	}

	// Write the alias if it exists
	if f.Alias != "" {
		if _, err := w.Write([]byte(" AS " + d.QuoteIdentifier(f.Alias))); err != nil {
			return nil, err
		}
	}

	return args, nil
}
