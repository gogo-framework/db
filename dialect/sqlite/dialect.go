package sqlite

import "fmt"

// SqliteDialect implements the dialect.SqliteDialect interface for SQLite
type SqliteDialect struct{}

// QuoteIdentifier quotes a table or column name for use in queries
func (d *SqliteDialect) QuoteIdentifier(name string) string {
	return fmt.Sprintf("\"%s\"", name)
}

// Placeholder returns the placeholder for a parameter at the given position
func (d *SqliteDialect) Placeholder(position int) string {
	return "?"
}

// NamedPlaceholder returns the placeholder for a named parameter
func (d *SqliteDialect) NamedPlaceholder(name string) string {
	return ":" + name
}

// SupportsNamedPlaceholders returns whether the dialect supports named placeholders
func (d *SqliteDialect) SupportsNamedPlaceholders() bool {
	return true
}

// SupportsReturning returns whether the dialect supports RETURNING clause
func (d *SqliteDialect) SupportsReturning() bool {
	// SQLite supports RETURNING since version 3.35.0 (2021-03-12)
	return true
}

// LimitOffset returns the SQL for LIMIT and OFFSET clauses
func (d *SqliteDialect) LimitOffset(limit, offset *int) string {
	if limit == nil && offset == nil {
		return ""
	}

	sql := ""
	if limit != nil {
		sql = fmt.Sprintf(" LIMIT %d", *limit)
	}
	if offset != nil {
		sql += fmt.Sprintf(" OFFSET %d", *offset)
	}
	return sql
}
