package main

import (
	"fmt"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gogo-framework/db/dialect/sqlite"
	"github.com/gogo-framework/db/internal"
	"github.com/gogo-framework/db/schema"
)

type User struct {
	ID   sqlite.Text
	Name sqlite.Text
	Bio  sqlite.Text
}

func (u *User) Table() *schema.Table {
	return sqlite.NewTable("users", func(t *schema.Table) {
		t.RegisterColumn("id", &u.ID)
		t.RegisterColumn("name", &u.Name)
		t.RegisterColumn("bio", &u.Bio)
	})
}

func main() {
	user := schema.NewTable[User]()
	query := sqlite.Select(
		&user.ID, &user.Name, &user.Bio,
		sqlite.From(user).As("u"),
	).ToSql()
	println(query)
	// testMapping()
}

func testMapping() {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM users").WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "bio"}).
			AddRow("1", "Alice", "Software developer").
			AddRow("2", "Bob", "UX designer").
			AddRow("3", "Charlie", "Project manager"),
	)

	mock.ExpectQuery("SELECT (.+) FROM users WHERE id = ?").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "bio"}).
			AddRow("1", "Alice", "Software developer"))

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}

	fmt.Println("All users:")
	for rows.Next() {
		user := &User{}
		table := user.Table()
		mapper := internal.NewRowMapper(table)

		err := mapper.MapRow(rows)
		if err != nil {
			log.Printf("Error mapping row: %v\n", err)
			continue
		}

		fmt.Printf("- %s (%s): %s\n",
			user.ID.Get(), user.Name.Get(), user.Bio.Get())
	}
	rows.Close()

	idRows, err := db.Query("SELECT * FROM users WHERE id = ?", "1")
	if err != nil {
		log.Fatalf("Error executing ID query: %v", err)
	}

	fmt.Println("\nUser with ID=1:")
	if idRows.Next() {
		user := &User{}
		table := user.Table()
		mapper := internal.NewRowMapper(table)

		err := mapper.MapRow(idRows)
		if err != nil {
			log.Fatalf("Error mapping ID row: %v", err)
		}

		fmt.Printf("ID=%s, Name=%s, Bio=%s\n",
			user.ID.Get(), user.Name.Get(), user.Bio.Get())
	}
	idRows.Close()
}
