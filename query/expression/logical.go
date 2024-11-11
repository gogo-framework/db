package expression

import "fmt"

type And struct {
	Left  Expression
	Right Expression
}

func (a And) ToSQL() (string, []any) {
	leftSQL, leftArgs := a.Left.ToSQL()
	rightSQL, rightArgs := a.Right.ToSQL()
	return fmt.Sprintf("(%s AND %s)", leftSQL, rightSQL), append(leftArgs, rightArgs...)
}

type Or struct {
	Left  Expression
	Right Expression
}

func (o Or) ToSQL() (string, []any) {
	leftSQL, leftArgs := o.Left.ToSQL()
	rightSQL, rightArgs := o.Right.ToSQL()
	return fmt.Sprintf("(%s OR %s)", leftSQL, rightSQL), append(leftArgs, rightArgs...)
}

type Not struct {
	Expr Expression
}

func (n Not) ToSQL() (string, []any) {
	sql, args := n.Expr.ToSQL()
	return fmt.Sprintf("NOT (%s)", sql), args
}
