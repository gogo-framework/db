package clause

import (
	"context"
	"io"

	"github.com/gogo-framework/db/dialect"
	"github.com/gogo-framework/db/schema"
)

type SelectClause struct {
	columns []schema.Column
}

func (sc *SelectClause) AppendColumns(column ...schema.Column) {
	sc.columns = append(sc.columns, column...)
}

// Implement the SqlWriter interface
func (sc *SelectClause) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	var args []any
	for i, column := range sc.columns {
		if i > 0 {
			w.Write([]byte(", "))
		}
		w.Write([]byte(column.GetName()))
	}
	return args, nil
}
