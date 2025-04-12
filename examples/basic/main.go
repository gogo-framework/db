package main

import (
	"fmt"

	"github.com/gogo-framework/db/dialect/sqlite"
	"github.com/gogo-framework/db/schema"
)

// User represents a user in the database
type User struct {
	table     *schema.Table
	ID        sqlite.Integer
	Username  sqlite.Text
	Email     sqlite.Text
	Age       sqlite.Integer
	Score     sqlite.Float
	CreatedAt sqlite.Text
}

// Table implements the Tabler interface
func (u *User) Table() *schema.Table {
	if u.table == nil {
		u.table = sqlite.NewTable("users", func(t *schema.Table) {
			t.RegisterColumn("id", &u.ID)
			t.RegisterColumn("username", &u.Username)
			t.RegisterColumn("email", &u.Email)
			t.RegisterColumn("age", &u.Age)
			t.RegisterColumn("score", &u.Score)
			t.RegisterColumn("created_at", &u.CreatedAt)
		})
	}
	return u.table
}

func main() {
	user := schema.NewTable[User]()

	// Comprehensive example demonstrating all supported clauses
	query, args := sqlite.Select(
		&user.ID, &user.Username, &user.Email, &user.Age, &user.Score, &user.CreatedAt,
		sqlite.From(user).As("u"),
		sqlite.Where(
			user.Username.Like("J%"),
			sqlite.Or(
				user.ID.Gt(5),
				user.ID.Lt(10),
			),
		).And(
			user.Email.Like("%developer%"),
		),
		sqlite.GroupBy(&user.Email),
		sqlite.Having(user.ID.Gt(3)),
		sqlite.OrderBy(&user.Username, &user.ID),
		sqlite.Limit(10),
		sqlite.Offset(20),
		sqlite.Distinct(),
	).ToSql()

	fmt.Println("Generated SQL Query:")
	fmt.Println(query)
	fmt.Println("\nQuery Arguments:")
	fmt.Println(args)
}
