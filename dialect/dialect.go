package dialect

// Dialect defines the interface for database-specific SQL query generation
type Dialect interface {
	// QuoteIdentifier quotes a table or column name for use in queries
	QuoteIdentifier(name string) string

	// Placeholder returns the placeholder for a parameter at the given position
	Placeholder(position int) string

	// NamedPlaceholder returns the placeholder for a named parameter
	// e.g. ":name" for SQLite/PostgreSQL, "@name" for MySQL
	NamedPlaceholder(name string) string

	// SupportsNamedPlaceholders returns whether the dialect supports named placeholders
	// Some dialects might only support positional parameters
	SupportsNamedPlaceholders() bool

	// SupportsReturning returns whether the dialect supports RETURNING clause
	// e.g. SQLite 3.35.0+ and PostgreSQL support it, MySQL doesn't
	SupportsReturning() bool

	// LimitOffset returns the SQL for LIMIT and OFFSET clauses
	// Some dialects use different syntax (e.g. FETCH FIRST n ROWS ONLY)
	LimitOffset(limit, offset *int) string
}
