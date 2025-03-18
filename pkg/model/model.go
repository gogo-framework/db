package model

import (
	"fmt"

	"github.com/gogo-framework/db/pkg/schema"
)

type Model interface {
	// Table returns the table for the model.
	// This method should be called before using the model to ensure everything is set up correctly.
	// You can use the `New` helper function to help with this.
	Table() *schema.Table
}

// New is a helper function to create a new model and call the Table method on it.
func New[T any]() *T {
	model := new(T)
	if m, ok := any(model).(Model); ok {
		m.Table()
	} else {
		panic(fmt.Sprintf("Type %T does not implement Model interface", model))
	}
	return model
}
