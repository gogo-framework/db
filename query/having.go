package query

import (
	"context"
	"io"

	"github.com/gogo-framework/db/dialect"
)

// HavingClause represents a HAVING clause
type HavingClause struct {
	Conditions []Condition
}

func (h *HavingClause) ApplySelect(stmt *SelectStmt) {
	stmt.having = h
}

func (h *HavingClause) WriteSql(ctx context.Context, writer io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	if len(h.Conditions) == 0 {
		return nil, nil
	}

	var args []any

	for i, condition := range h.Conditions {
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

// Having creates a HAVING clause
func Having(conditions ...Condition) SelectPart {
	return &HavingClause{
		Conditions: conditions,
	}
}

// And adds additional conditions to an existing HAVING clause
func (h *HavingClause) And(conditions ...Condition) *HavingClause {
	h.Conditions = append(h.Conditions, conditions...)
	return h
}
