package main

import (
	"fmt"

	"github.com/gogo-framework/db/dialect/sqlite"
	"github.com/gogo-framework/db/schema"
)

// User represents a user in the database
type User struct {
	ID    sqlite.Text
	Name  sqlite.Text
	Bio   sqlite.Text
	table *schema.Table
}

// Table implements the Tabler interface
func (u *User) Table() *schema.Table {
	if u.table == nil {
		u.table = sqlite.NewTable("users", func(t *schema.Table) {
			t.RegisterColumn("id", &u.ID)
			t.RegisterColumn("name", &u.Name)
			t.RegisterColumn("bio", &u.Bio)
		})
	}
	return u.table
}

func main() {
	user := schema.NewTable[User]()

	// Comprehensive example demonstrating all supported clauses
	query, args := sqlite.Select(
		&user.ID, &user.Name, &user.Bio,
		sqlite.From(user).As("u"),
		sqlite.Where(
			user.Name.Like("J%"),
			sqlite.Or(
				user.ID.Gt(5),
				user.ID.Lt(10),
			),
		).And(
			user.Bio.Like("%developer%"),
		),
		sqlite.GroupBy(&user.Bio),
		sqlite.Having(
			user.ID.Gt(3),
		),
		sqlite.OrderBy(&user.Name, &user.ID),
		sqlite.Limit(10),
		sqlite.Offset(20),
		sqlite.Distinct(),
	).ToSql()

	fmt.Println("Generated SQL Query:")
	fmt.Println(query)
	fmt.Println("\nQuery Arguments:")
	fmt.Println(args)
}
