package query

import (
	"fmt"
	"strings"

	"github.com/gogo-framework/db/query/expression"
	"github.com/gogo-framework/db/schema"
)

type OrderDirection string

const (
	Asc  OrderDirection = "ASC"
	Desc OrderDirection = "DESC"
)

type SelectQuery[T schema.Model] struct {
	model   T
	columns []string
	from    string
	where   []expression.Expression
	orderBy []string
	limit   *int
	offset  *int
	groupBy []string
	having  []expression.Expression
}

// Select starts building a SELECT query
func Select[T schema.Model](columns ...string) *SelectQuery[T] {
	return &SelectQuery[T]{
		columns: columns,
	}
}

// From specifies the table to select from
func (q *SelectQuery[T]) From(table string) *SelectQuery[T] {
	q.from = table
	return q
}

// Where adds WHERE conditions
func (q *SelectQuery[T]) Where(conditions ...expression.Expression) *SelectQuery[T] {
	q.where = append(q.where, conditions...)
	return q
}

// OrderBy adds ORDER BY clauses
func (q *SelectQuery[T]) OrderBy(column string, direction OrderDirection) *SelectQuery[T] {
	q.orderBy = append(q.orderBy, fmt.Sprintf("%s %s", column, direction))
	return q
}

// Limit sets the LIMIT clause
func (q *SelectQuery[T]) Limit(limit int) *SelectQuery[T] {
	q.limit = &limit
	return q
}

// Offset sets the OFFSET clause
func (q *SelectQuery[T]) Offset(offset int) *SelectQuery[T] {
	q.offset = &offset
	return q
}

// GroupBy adds GROUP BY clauses
func (q *SelectQuery[T]) GroupBy(columns ...string) *SelectQuery[T] {
	q.groupBy = append(q.groupBy, columns...)
	return q
}

// Having adds HAVING conditions
func (q *SelectQuery[T]) Having(conditions ...expression.Expression) *SelectQuery[T] {
	q.having = append(q.having, conditions...)
	return q
}

func (q *SelectQuery[T]) ToSQL() (string, []any) {
	var query strings.Builder
	var args []any

	// SELECT
	query.WriteString("SELECT ")
	if len(q.columns) == 0 {
		query.WriteString("*")
	} else {
		query.WriteString(strings.Join(q.columns, ", "))
	}

	// FROM
	query.WriteString(" FROM ")
	query.WriteString(q.model.Table().Name)

	// WHERE
	if len(q.where) > 0 {
		query.WriteString(" WHERE ")
		whereStrs := make([]string, 0, len(q.where))
		for _, w := range q.where {
			sql, wargs := w.ToSQL()
			whereStrs = append(whereStrs, sql)
			args = append(args, wargs...)
		}
		query.WriteString(strings.Join(whereStrs, " AND "))
	}

	// GROUP BY
	if len(q.groupBy) > 0 {
		query.WriteString(" GROUP BY ")
		query.WriteString(strings.Join(q.groupBy, ", "))
	}

	// HAVING
	if len(q.having) > 0 {
		query.WriteString(" HAVING ")
		havingStrs := make([]string, 0, len(q.having))
		for _, having := range q.having {
			sql, havingArgs := having.ToSQL()
			havingStrs = append(havingStrs, sql)
			args = append(args, havingArgs...)
		}
		query.WriteString(strings.Join(havingStrs, " AND "))
	}

	// ORDER BY
	if len(q.orderBy) > 0 {
		query.WriteString(" ORDER BY ")
		query.WriteString(strings.Join(q.orderBy, ", "))
	}

	// LIMIT
	if q.limit != nil {
		query.WriteString(fmt.Sprintf(" LIMIT %d", *q.limit))
	}

	// OFFSET
	if q.offset != nil {
		query.WriteString(fmt.Sprintf(" OFFSET %d", *q.offset))
	}

	return query.String(), args
}
