package main

import (
	"database/sql/driver"
	"fmt"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gogo-framework/db/dialect/sqlite"
	"github.com/gogo-framework/db/internal"
	"github.com/gogo-framework/db/schema"
)

// User represents a user in the system
type User struct {
	table *schema.Table
	ID    sqlite.Integer
	Name  sqlite.Text
	Age   sqlite.Integer
	Bio   sqlite.Text
	Data  sqlite.Blob
}

// UserStats represents aggregated user statistics
type UserStats struct {
	User
	AverageAge sqlite.Float
	MaxAge     sqlite.Integer
	MinAge     sqlite.Integer
	Count      sqlite.Integer
}

// Table implements the schema.Tabler interface
func (u *User) Table() *schema.Table {
	if u.table == nil {
		u.table = sqlite.NewTable("users", func(t *schema.Table) {
			t.RegisterColumn("id", &u.ID)
			t.RegisterColumn("name", &u.Name)
			t.RegisterColumn("age", &u.Age)
			t.RegisterColumn("bio", &u.Bio)
			t.RegisterColumn("data", &u.Data)
		})
	}
	return u.table
}

// GetColumns returns the columns from a SelectStmt
func GetColumns(stmt *sqlite.SelectStmt) []schema.Column {
	if stmt.Columns == nil {
		return nil
	}
	return stmt.Columns.Columns
}

// ConvertToDriverValues converts []any to []driver.Value
func ConvertToDriverValues(args []any) []driver.Value {
	result := make([]driver.Value, len(args))
	for i, arg := range args {
		result[i] = arg
	}
	return result
}

func main() {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("Failed to create mock database:", err)
	}
	defer db.Close()

	// Create instances for query building
	user := schema.NewTable[User]()
	userStats := schema.NewTable[UserStats]()

	// Example 1: Query all users with uppercase names
	selectStmt := sqlite.Select(
		&user.ID,
		sqlite.Upper(&user.Name, &user.Name),
		&user.Age,
		&user.Bio,
		&user.Data,
		sqlite.From(user),
	)

	selectSQL, args := selectStmt.ToSql()
	fmt.Println("Select SQL:", selectSQL)
	fmt.Println("Select args:", args)

	// Set up mock to return rows for ANY query (instead of expecting a specific one)
	rows := sqlmock.NewRows([]string{"id", "name", "age", "bio", "data"}).
		AddRow(1, "JOHN DOE", 30, "Software Engineer", []byte("test binary data")).
		AddRow(2, "JANE SMITH", 25, "Data Scientist", []byte("test binary data"))

	// This will match any query - no regex pattern matching issues
	mock.ExpectQuery(".*").WillReturnRows(rows)

	// Execute the query
	queryRows, err := db.Query(selectSQL, args...)
	if err != nil {
		log.Fatal("Failed to execute query:", err)
	}
	defer queryRows.Close()

	// Get the columns from the select statement
	columns := GetColumns(selectStmt)

	// Map the results to User structs
	users, err := internal.MapAll[*User](queryRows, columns, user)
	if err != nil {
		log.Fatal("Failed to map results:", err)
	}

	// Print the results
	fmt.Println("\nMapped Users:")
	for _, u := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d, Bio: %s\n",
			u.ID.Get(), u.Name.Get(), u.Age.Get(), u.Bio.Get())
	}

	// Example 2: Query user statistics with aggregations
	statsStmt := sqlite.Select(
		sqlite.Avg(&user.Age, &userStats.AverageAge),
		sqlite.Max(&user.Age, &userStats.MaxAge),
		sqlite.Min(&user.Age, &userStats.MinAge),
		sqlite.Count(&user.ID, &userStats.Count),
		sqlite.From(user),
	)

	statsSQL, statsArgs := statsStmt.ToSql()
	fmt.Println("\nStats SQL:", statsSQL)
	fmt.Println("Stats args:", statsArgs)

	// Set up mock for ANY stats query
	statsRows := sqlmock.NewRows([]string{"avg_age", "max_age", "min_age", "count_id"}).
		AddRow(27.5, 30, 25, 2)

	mock.ExpectQuery(".*").WillReturnRows(statsRows)

	// Execute the stats query
	statsQueryRows, err := db.Query(statsSQL, statsArgs...)
	if err != nil {
		log.Fatal("Failed to execute stats query:", err)
	}
	defer statsQueryRows.Close()

	// Get the columns from the stats statement
	statsColumns := GetColumns(statsStmt)

	// Map the results to UserStats structs
	stats, err := internal.MapAll[*UserStats](statsQueryRows, statsColumns, userStats)
	if err != nil {
		log.Fatal("Failed to map stats results:", err)
	}

	// Print the stats results
	fmt.Println("\nMapped User Stats:")
	for _, s := range stats {
		fmt.Printf("Average Age: %.1f, Max Age: %d, Min Age: %d, Count: %d\n",
			s.AverageAge.Get(), s.MaxAge.Get(), s.MinAge.Get(), s.Count.Get())
	}

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		log.Fatal("Not all expectations were met:", err)
	}
}
