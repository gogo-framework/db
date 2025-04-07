package sqlite

import "github.com/gogo-framework/db/schema"

func NewTable(name string, setup func(t *schema.Table)) *schema.Table {
	t := &schema.Table{Name: name}
	setup(t)
	return t
}
