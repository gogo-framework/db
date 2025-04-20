package main

import (
	"fmt"

	"github.com/gogo-framework/db/dialect/sqlite"
	"github.com/gogo-framework/db/schema"
)

// User represents a user in the database
type User struct {
	ID        sqlite.Integer
	Username  sqlite.Text
	Email     sqlite.Text
	Age       sqlite.Integer
	Score     sqlite.Float
	CreatedAt sqlite.Text
}

// Table implements the Tabler interface
func (u *User) Table() *schema.Table {
	return sqlite.NewTable("users", func(t *schema.Table) {
		t.RegisterColumn("id", &u.ID)
		t.RegisterColumn("username", &u.Username)
		t.RegisterColumn("email", &u.Email)
		t.RegisterColumn("age", &u.Age)
		t.RegisterColumn("score", &u.Score)
		t.RegisterColumn("created_at", &u.CreatedAt)
	})
}

// UserStats represents aggregated user statistics
type UserStats struct {
	User
	AverageAge    sqlite.Float
	UppercaseName sqlite.Text
}

func main() {
	user := schema.NewTable[User]()
	query, args := sqlite.Select(
		sqlite.Distinct(),
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
	).ToSql()

	fmt.Println("Generated SQL Query:")
	fmt.Println(query)
	fmt.Println("\nQuery Arguments:")
	fmt.Println(args)

	// New example with aggregation and function
	userStats := schema.NewTable[UserStats]()
	query, args = sqlite.Select(
		&userStats.ID,
		&userStats.Username,
		sqlite.Avg(&userStats.Age, &userStats.AverageAge),
		sqlite.Upper(&userStats.Username, &userStats.UppercaseName),
		sqlite.From(userStats),
		sqlite.GroupBy(&userStats.ID, &userStats.Username),
	).ToSql()

	fmt.Println("\nGenerated SQL Query:")
	fmt.Println(query)
	fmt.Println("\nQuery Arguments:")
	fmt.Println(args)

	// New example with SQLite function
	query, args = sqlite.Select(
		&user.ID,
		sqlite.Upper(&user.Username, &user.Username),
		sqlite.From(user),
	).ToSql()

	fmt.Println("\nGenerated SQL Query with Function:")
	fmt.Println(query)
	fmt.Println("\nQuery Arguments:")
	fmt.Println(args)
}
