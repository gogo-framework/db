package query

import (
	"context"
	"io"

	"github.com/gogo-framework/db/dialect"
)

type WherePart interface {
	ApplyWhere(*SelectStmt)
}

type WhereClause struct {
	Conditions []Condition
}

// ApplyWhere implements the WherePart interface
func (w WhereClause) ApplyWhere(stmt *SelectStmt) {
	stmt.where = w
}

// ApplySelect implements the SelectPart interface
func (w WhereClause) ApplySelect(stmt *SelectStmt) {
	stmt.where = w
}

// WriteSql implements the SqlWriter interface
func (w WhereClause) WriteSql(ctx context.Context, writer io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	if len(w.Conditions) == 0 {
		return nil, nil
	}

	var args []any

	for i, condition := range w.Conditions {
		if i > 0 {
			writer.Write([]byte(" AND "))
		}

		conditionArgs, err := condition.WriteSql(ctx, writer, d, argPos+len(args))
		if err != nil {
			return nil, err
		}
		args = append(args, conditionArgs...)
	}

	return args, nil
}

// Where creates a new WHERE clause with the given conditions
func Where(conditions ...Condition) WhereClause {
	return WhereClause{Conditions: conditions}
}

// And adds additional conditions to an existing WHERE clause
func (w WhereClause) And(conditions ...Condition) WhereClause {
	w.Conditions = append(w.Conditions, conditions...)
	return w
}

// Or creates an OR condition between multiple conditions
func Or(conditions ...Condition) Condition {
	return &OrCondition{Conditions: conditions}
}
