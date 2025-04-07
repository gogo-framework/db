package sqlite

import (
	"github.com/gogo-framework/db/internal/query/clause"
	"github.com/gogo-framework/db/schema"
)

type SelectClause struct{ clause.SelectClause }

type FromClause struct{ clause.FromClause }

func (f FromClause) As(alias string) FromClause {
	f.FromClause.As(alias)
	return f
}

// Implement the SelectPart interface.
func (f FromClause) ApplySelect(stmt *SelectStmt) {
	stmt.from = f
}

func From(source schema.Tabler) *FromClause {
	return &FromClause{
		FromClause: clause.FromClause{Source: source},
	}
}
