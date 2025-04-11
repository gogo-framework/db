package query

import (
	"context"
	"io"

	"github.com/gogo-framework/db/dialect"
)

type SqlWriter interface {
	WriteSql(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int) ([]any, error)
}

// WriteSqlWithSeparator writes multiple SqlWriter items with a separator between them
func WriteSqlWithSeparator(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int, separator string, items ...SqlWriter) ([]any, error) {
	var args []any

	for i, item := range items {
		if i > 0 {
			w.Write([]byte(separator))
		}

		itemArgs, err := item.WriteSql(ctx, w, d, argPos+len(args))
		if err != nil {
			return nil, err
		}
		args = append(args, itemArgs...)
	}

	return args, nil
}

// WriteSqlWithPrefix writes a SqlWriter with a prefix
func WriteSqlWithPrefix(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int, prefix string, item SqlWriter) ([]any, error) {
	w.Write([]byte(prefix))
	return item.WriteSql(ctx, w, d, argPos)
}

// WriteSqlWithSuffix writes a SqlWriter with a suffix
func WriteSqlWithSuffix(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int, suffix string, item SqlWriter) ([]any, error) {
	args, err := item.WriteSql(ctx, w, d, argPos)
	if err != nil {
		return nil, err
	}
	w.Write([]byte(suffix))
	return args, nil
}

// WriteSqlWithParentheses writes a SqlWriter wrapped in parentheses
func WriteSqlWithParentheses(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int, item SqlWriter) ([]any, error) {
	w.Write([]byte("("))
	args, err := item.WriteSql(ctx, w, d, argPos)
	if err != nil {
		return nil, err
	}
	w.Write([]byte(")"))
	return args, nil
}

// WriteSqlIf writes a SqlWriter only if the condition is true
func WriteSqlIf(ctx context.Context, w io.Writer, d dialect.Dialect, argPos int, condition bool, item SqlWriter) ([]any, error) {
	if !condition {
		return nil, nil
	}
	return item.WriteSql(ctx, w, d, argPos)
}
