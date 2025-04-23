/*
	This file contains a POC for mapping back rows to a table struct.
	It'll not be used in the final version, but tries to proof that the current implementation
	is good enough in order to map back results.
*/

package internal

import (
	"database/sql"
	"fmt"

	"github.com/gogo-framework/db/internal/schema"
)

type RowMapper struct {
	columns []schema.Column
}

func NewRowMapper(columns []schema.Column) *RowMapper {
	return &RowMapper{columns: columns}
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

	// Create a map of column names to their indices for faster lookup
	colNameToIndex := make(map[string]int)
	for i, name := range colNames {
		colNameToIndex[name] = i
	}

	// Process each column in our schema
	for _, col := range rm.columns {
		colName := col.GetName()
		index, exists := colNameToIndex[colName]
		if !exists {
			continue
		}

		value := *(values[index].(*any))
		if err := col.Scan(value); err != nil {
			return fmt.Errorf("failed to scan column %s: %w", colName, err)
		}
	}

	return nil
}

func MapAll[T schema.Tabler](rows *sql.Rows, columns []schema.Column, model T) ([]T, error) {
	var result []T

	for rows.Next() {
		// Create a copy of the model for each row
		modelCopy := model
		mapper := NewRowMapper(columns)

		if err := mapper.MapRow(rows); err != nil {
			return nil, err
		}

		result = append(result, modelCopy)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func MapOne[T schema.Tabler](rows *sql.Rows, columns []schema.Column, model T) (*T, error) {
	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	mapper := NewRowMapper(columns)

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
