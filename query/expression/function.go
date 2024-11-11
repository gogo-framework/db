package expression

import (
	"fmt"
	"strings"
)

type Function struct {
    Name string
    Args []Expression
}

func (f Function) ToSQL() (string, []any) {
    var args []string
    var sqlArgs []any

    for _, arg := range f.Args {
        argSQL, argValues := arg.ToSQL()
        args = append(args, argSQL)
        sqlArgs = append(sqlArgs, argValues...)
    }

    return fmt.Sprintf("%s(%s)", f.Name, strings.Join(args, ", ")), sqlArgs
}

// Common function constructors
func Count(expr Expression) Function {
    return Function{
        Name: "COUNT",
        Args: []Expression{expr},
    }
}

func Sum(expr Expression) Function {
    return Function{
        Name: "SUM",
        Args: []Expression{expr},
    }
}

func Max(expr Expression) Function {
    return Function{
        Name: "MAX",
        Args: []Expression{expr},
    }
}

func Min(expr Expression) Function {
    return Function{
        Name: "MIN",
        Args: []Expression{expr},
    }
}

func Avg(expr Expression) Function {
    return Function{
        Name: "AVG",
        Args: []Expression{expr},
    }
}
