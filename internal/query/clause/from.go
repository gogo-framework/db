package clause

import (
	"context"
	"io"

	"github.com/gogo-framework/db/dialect"
	"github.com/gogo-framework/db/schema"
)

type FromClause struct {
	Source schema.Tabler
	Alias  string
	Joins  []any // We have no type yet for joins.
}

func (f *FromClause) As(alias string) {
	f.Alias = alias
}

func (f *FromClause) AppendJoins(joins ...any) {
	f.Joins = append(f.Joins, joins...)
}

// Implement the SqlWriter interface
func (f *FromClause) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	var args []any

	w.Write([]byte("FROM "))

	if f.Source != nil {
		if f.Alias != "" {
			w.Write([]byte(f.Source.Table().Name + " AS " + f.Alias))
		} else {
			w.Write([]byte(f.Source.Table().Name))
		}
	}

	return args, nil
}
