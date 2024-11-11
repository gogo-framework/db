package expression

import "fmt"

// Expression represents anything that can be converted to SQL
type Expression interface {
	ToSQL() (string, []any)
}

// Comparison represents binary operations like =, >, <, etc.
// ToDo: Check if we can make Right type safe with generics.
type Comparison struct {
	Left  string
	Op    string
	Right any
}

// ToDo: Types for switch case ops
func (c Comparison) ToSQL() (string, []any) {
	switch c.Op {
	case "IN", "NOT IN":
		return fmt.Sprintf("%s %s (?)", c.Left, c.Op), []any{c.Right}
	case "BETWEEN":
		// Handle different types of slices
		switch v := c.Right.(type) {
		case []int64:
			return fmt.Sprintf("%s BETWEEN ? AND ?", c.Left), []any{v[0], v[1]}
		case []float64:
			return fmt.Sprintf("%s BETWEEN ? AND ?", c.Left), []any{v[0], v[1]}
		case []string:
			return fmt.Sprintf("%s BETWEEN ? AND ?", c.Left), []any{v[0], v[1]}
		case []any:
			return fmt.Sprintf("%s BETWEEN ? AND ?", c.Left), []any{v[0], v[1]}
		default:
			panic(fmt.Sprintf("unsupported BETWEEN type: %T", c.Right))
		}
	case "IS NULL", "IS NOT NULL":
		return fmt.Sprintf("%s %s", c.Left, c.Op), nil
	default:
		return fmt.Sprintf("%s %s ?", c.Left, c.Op), []any{c.Right}
	}
}
