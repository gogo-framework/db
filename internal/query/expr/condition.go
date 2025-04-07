package expr

import (
	"context"
	"io"

	"github.com/gogo-framework/db/dialect"
	"github.com/gogo-framework/db/internal/query"
)

type Operator string

const (
	Equal              Operator = "="
	NotEqual           Operator = "!="
	GreaterThan        Operator = ">"
	GreaterThanOrEqual Operator = ">="
	LessThan           Operator = "<"
	LessThanOrEqual    Operator = "<="
	Like               Operator = "LIKE"
	NotLike            Operator = "NOT LIKE"
	In                 Operator = "IN"
	NotIn              Operator = "NOT IN"
	Glob               Operator = "GLOB"
	Match              Operator = "MATCH"
	Regexp             Operator = "REGEXP"
)

type Condition struct {
	Left  query.SqlWriter
	Op    Operator
	Right query.SqlWriter
}

func (c *Condition) WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error) {
	var args []any

	leftArgs, err := c.Left.WriteSql(ctx, w, d, argPos)
	if err != nil {
		return nil, err
	}
	args = append(args, leftArgs...)

	w.Write([]byte(" " + string(c.Op) + " "))

	rightArgs, err := c.Right.WriteSql(ctx, w, d, argPos+len(args))
	if err != nil {
		return nil, err
	}
	args = append(args, rightArgs...)

	return args, nil
}
