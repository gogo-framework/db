/*
	This file contains a POC for mapping back rows to a table struct.
	It'll not be used in the final version, but tries to proof that the current implementation
	is good enough in order to map back results.
*/

package internal

import (
	"database/sql"
	"fmt"

	"github.com/gogo-framework/db/schema"
)

type RowMapper struct {
	table *schema.Table
}

func NewRowMapper(table *schema.Table) *RowMapper {
	return &RowMapper{table: table}
}

func (rm *RowMapper) MapRow(rows *sql.Rows) error {
	colNames, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("failed to get column names: %w", err)
	}

	values := make([]any, len(colNames))
	for i := range values {
		values[i] = new(any)
	}

	if err := rows.Scan(values...); err != nil {
		return fmt.Errorf("failed to scan row: %w", err)
	}

	for i, name := range colNames {
		col, ok := rm.table.GetColumn(name)
		if !ok {
			continue
		}

		value := *(values[i].(*any))
		if err := col.Scan(value); err != nil {
			return fmt.Errorf("failed to scan column %s: %w", name, err)
		}
	}

	return nil
}

func MapAll[T schema.Tabler](rows *sql.Rows) ([]T, error) {
	var result []T

	for rows.Next() {
		var model T
		table := model.Table()
		mapper := NewRowMapper(table)

		if err := mapper.MapRow(rows); err != nil {
			return nil, err
		}

		result = append(result, model)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func MapOne[T schema.Tabler](rows *sql.Rows) (*T, error) {
	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	var model T
	table := model.Table()
	mapper := NewRowMapper(table)

	if err := mapper.MapRow(rows); err != nil {
		return nil, err
	}

	if rows.Next() {
		return &model, fmt.Errorf("multiple rows returned for single result query")
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &model, nil
}

func FindByID[T schema.Tabler](db *sql.DB, id any) (*T, error) {
	var model T
	table := model.Table()

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", table.Name)

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return MapOne[T](rows)
}

func Find[T schema.Tabler](db *sql.DB, query string, args ...any) ([]T, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return MapAll[T](rows)
}

func FindOne[T schema.Tabler](db *sql.DB, query string, args ...any) (*T, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return MapOne[T](rows)
}
