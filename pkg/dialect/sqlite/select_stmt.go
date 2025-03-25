package sqlite

type SelectPart interface {
	Apply(*SelectStmt)
}

// https://www.sqlite.org/lang_select.html
type SelectStmt struct {
	with     any
	selects  any
	distinct bool
	from     any
	where    any
	groupBy  any
	having   any
	windows  any
	orderBy  any
	limit    any
	offset   any
}

/*

The API for the query builder will be something like;

user := model.New[UserWithAggs]()
sqlite.Select(
	sqlite.Columns(user.Name, user.UppercaseName),
	sqlite.From(user),
	sqlite.Where(user.Name.Eq("HENKIE"), sqlite.And(user.UppercaseName.Eq("henkie"))),
	sqlite.GroupBy(user.Name),
	sqlite.Having(user.Name.Eq("HENKIE")),
	sqlite.OrderBy(user.Name),
	sqlite.Limit(10),
	sqlite.Offset(10),
)

So instead of the traditional builder pattern, it's a more function approach.
This allows more flexibility in the API and allowing others to implement custom functions.

Under the hood, the functions will apply the appropriate logic to the SelectStmt struct.
*/

func Select(parts ...SelectPart) *SelectStmt {
	stmt := &SelectStmt{}
	for _, part := range parts {
		part.Apply(stmt)
	}
	return stmt
}

type FromClause struct {
	table any
}

func (f *FromClause) Apply(stmt *SelectStmt) {
	stmt.from = f.table
}

func From(table any) *FromClause {
	return &FromClause{table}
}
