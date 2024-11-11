package main

import (
	"github.com/gogo-framework/db/query"
	"github.com/gogo-framework/db/schema"
	"github.com/gogo-framework/db/schema/column"
)

// Define our User model
type User struct {
	ID      column.Integer
	Name    column.Text
	Email   column.Text
	Age     column.Integer
	IsAdmin column.Boolean
}

func (User) Table() schema.Table {
	return schema.Table{
		Name: "users",
	}
}

func NewUser() User {
	return User{
		ID:      column.Integer{Name: "id"},
		Name:    column.Text{Name: "name"},
		Email:   column.Text{Name: "email"},
		Age:     column.Integer{Name: "age"},
		IsAdmin: column.Boolean{Name: "is_admin"},
	}
}

func main() {
	user := NewUser()

	q1 := query.Select[User]("id", "name", "email").Where(user.Age.Gt(18))
	sql1, args1 := q1.ToSQL()

	if sql1 != "SELECT id, name, email FROM users WHERE age > ?" {
		panic("Query 1 is incorrect")
	}

	if len(args1) != 1 || args1[0].(int64) != 18 {
		panic("Query 1 args are incorrect")
	}

	q2 := query.Select[User]("id", "name").
		Where(
			user.Age.Between(18, 65),
			user.Email.Like("%@example.com"),
			user.IsAdmin.Eq(true),
		).
		GroupBy("age").
		Having(user.Age.Gt(21)).
		OrderBy("name", query.Asc).
		Limit(10).
		Offset(20)
	sql2, args2 := q2.ToSQL()

	if sql2 != "SELECT id, name FROM users WHERE age BETWEEN ? AND ? AND email LIKE ? AND is_admin = ? GROUP BY age HAVING age > ? ORDER BY name ASC LIMIT 10 OFFSET 20" {
		panic("Query 2 is incorrect")
	}

	if len(args2) != 5 || args2[0].(int64) != 18 || args2[1].(int64) != 65 || args2[2].(string) != "%@example.com" || args2[3].(bool) != true || args2[4].(int64) != 21 {
		panic("Query 2 args are incorrect")
	}
}
