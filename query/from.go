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

func (f *FromClause) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	var args []any

	if f.Source != nil {
		if f.Alias != "" {
			w.Write([]byte(f.Source.Table().Name + " AS " + f.Alias))
		} else {
			w.Write([]byte(f.Source.Table().Name))
		}
	}

	return args, nil
}
