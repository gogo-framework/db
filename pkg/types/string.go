package types

import (
	"database/sql"
	"database/sql/driver"
)

type String string

func (s *String) Scan(value interface{}) error {
	var ns sql.Null[string]

	if err := ns.Scan(value); err != nil {
		return err
	}

	if ns.Valid {
		*s = String(ns.V)
	} else {
		*s = ""
	}
	return nil
}

func (s String) Value() (driver.Value, error) {
	return string(s), nil
}

func (s String) Eq(other string) bool {
	return string(s) == other
}
