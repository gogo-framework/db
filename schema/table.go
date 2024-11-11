package schema

type Table struct {
	Name string
	// Could add table-level configurations like:
	// Indexes   []Index
	// Engine    string
	// Charset   string
	// etc.
}

type Model interface {
	Table() Table
}
