package main

import (
	"strings"

	"github.com/gogo-framework/db/pkg/model"
	"github.com/gogo-framework/db/pkg/schema"
	"github.com/gogo-framework/db/pkg/types"
)

type User struct {
	Name types.String
}

func (u *User) Table() *schema.Table {
	table := &schema.Table{Name: "users", Schema: "public"}
	table.RegisterColumn(&u.Name, schema.NewColumnDefinition("name"))
	return table
}

type UserWithAggs struct {
	User
	UppercaseName types.String
}

func (u *UserWithAggs) Table() *schema.Table {
	table := u.User.Table()
	table.RegisterColumn(&u.UppercaseName, schema.NewColumnDefinition("uppercase_name"))
	return table
}

func main() {
	user := model.New[UserWithAggs]()
	user.Name = "HENKIE"
	user.UppercaseName = types.String(strings.ToUpper(string(user.Name)))
	println(user.Name)
	println(user.UppercaseName)
	table := user.Table()
	colDef, _ := table.GetColumnFromValuePtr(&user.Name)
	println(colDef.GetTable())
	println(colDef.Name)
	colDef, _ = table.GetColumnFromValuePtr(&user.UppercaseName)
	println(colDef.Name)
	println(colDef.GetTable())

}
