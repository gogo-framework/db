package sqlite

type InsertPart interface {
	ApplyInsert(*InsertStmt)
}

// https://www.sqlite.org/lang_insert.html
type InsertStmt struct {
}

func Insert(parts ...InsertPart) *InsertStmt {
	stmt := &InsertStmt{}
	for _, part := range parts {
		part.ApplyInsert(stmt)
	}
	return stmt
}
