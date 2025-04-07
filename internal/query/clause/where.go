package clause

import (
	"github.com/gogo-framework/db/internal/query/expr"
)

type WhereClause struct {
	Conditions []expr.Condition
}
