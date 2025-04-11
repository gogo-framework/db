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

	// Example 1: Basic select with method chaining for conditions
	query1, args1 := sqlite.Select(
		&user.ID, &user.Name, &user.Bio,
		sqlite.From(user).As("u"),
		sqlite.Where(
			user.Name.Eq("John"),
			user.ID.Gt(5),
		),
	).ToSql()
	fmt.Println("Query 1:", query1)
	fmt.Println("Args 1:", args1)

	// Example 2: Using OR conditions
	query2, args2 := sqlite.Select(
		&user.ID, &user.Name,
		sqlite.From(user),
		sqlite.Where(
			sqlite.Or(
				user.Name.Eq("John"),
				user.Name.Eq("Jane"),
			),
		),
	).ToSql()
	fmt.Println("\nQuery 2:", query2)
	fmt.Println("Args 2:", args2)

	// Example 3: Using LIKE and IN
	query3, args3 := sqlite.Select(
		&user.ID, &user.Name,
		sqlite.From(user),
		sqlite.Where(
			user.Name.Like("J%"),
			user.ID.In(1, 2, 3, 4, 5),
		),
	).ToSql()
	fmt.Println("\nQuery 3:", query3)
	fmt.Println("Args 3:", args3)

	// Example 4: Combining AND and OR
	query4, args4 := sqlite.Select(
		&user.ID, &user.Name, &user.Bio,
		sqlite.From(user),
		sqlite.Where(
			user.Name.Eq("John"),
			sqlite.Or(
				user.ID.Gt(5),
				user.ID.Lt(10),
			),
		),
	).ToSql()
	fmt.Println("\nQuery 4:", query4)
	fmt.Println("Args 4:", args4)

	// Example 5: Using ORDER BY
	query5, args5 := sqlite.Select(
		&user.ID, &user.Name,
		sqlite.From(user),
		sqlite.OrderBy(&user.Name),
	).ToSql()
	fmt.Println("\nQuery 5:", query5)
	fmt.Println("Args 5:", args5)

	// Example 6: Using ORDER BY with multiple columns
	query6, args6 := sqlite.Select(
		&user.ID, &user.Name, &user.Bio,
		sqlite.From(user),
		sqlite.OrderBy(&user.Name, &user.ID),
	).ToSql()
	fmt.Println("\nQuery 6:", query6)
	fmt.Println("Args 6:", args6)

	// Example 7: Using LIMIT
	query7, args7 := sqlite.Select(
		&user.ID, &user.Name,
		sqlite.From(user),
		sqlite.Limit(10),
	).ToSql()
	fmt.Println("\nQuery 7:", query7)
	fmt.Println("Args 7:", args7)

	// Example 8: Using OFFSET
	query8, args8 := sqlite.Select(
		&user.ID, &user.Name,
		sqlite.From(user),
		sqlite.Offset(20),
	).ToSql()
	fmt.Println("\nQuery 8:", query8)
	fmt.Println("Args 8:", args8)

	// Example 9: Combining ORDER BY, LIMIT, and OFFSET
	query9, args9 := sqlite.Select(
		&user.ID, &user.Name, &user.Bio,
		sqlite.From(user),
		sqlite.OrderBy(&user.Name),
		sqlite.Limit(10),
		sqlite.Offset(20),
	).ToSql()
	fmt.Println("\nQuery 9:", query9)
	fmt.Println("Args 9:", args9)
}
