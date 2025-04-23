package sqlite

import (
	"github.com/gogo-framework/db/internal/query"
	"github.com/gogo-framework/db/internal/schema"
)

func Or(conditions ...query.Condition) query.Condition {
	return &query.OrCondition{Conditions: conditions}
}

func Eq[T any](column schema.Column, value T) query.Condition {
	return query.Eq(column, value)
}

func Neq[T any](column schema.Column, value T) query.Condition {
	return query.Neq(column, value)
}

func Gt[T any](column schema.Column, value T) query.Condition {
	return query.Gt(column, value)
}

func Gte[T any](column schema.Column, value T) query.Condition {
	return query.Gte(column, value)
}

func Lt[T any](column schema.Column, value T) query.Condition {
	return query.Lt(column, value)
}

func Lte[T any](column schema.Column, value T) query.Condition {
	return query.Lte(column, value)
}

func Like(column schema.Column, pattern string) query.Condition {
	return query.Like(column, pattern)
}

func In[T any](column schema.Column, values ...T) query.Condition {
	return query.In(column, values...)
}
