package db 

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func QueryTable[T any](conn *sql.DB, tableName string) ([]T, error) {
	// Create a zero value of T to inspect its fields
	var sample T
	sampleType := reflect.TypeOf(sample)
	if sampleType.Kind() != reflect.Struct {
		return nil, errors.New("generic type must be a struct")
	}

	// Build SELECT statement from field names
	var fields []string
	for i := 0; i < sampleType.NumField(); i++ {
		fields = append(fields, sampleType.Field(i).Name)
	}
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(fields, ", "), tableName)

	// Execute the query
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []T

	for rows.Next() {
		elemPtr := reflect.New(sampleType) // *T
		elem := elemPtr.Elem()             // T

		dest := make([]interface{}, elem.NumField())
		for i := 0; i < elem.NumField(); i++ {
			dest[i] = elem.Field(i).Addr().Interface()
		}

		if err := rows.Scan(dest...); err != nil {
			return nil, err
		}

		results = append(results, elem.Interface().(T))
	}

	return results, nil
}


func Query(conn *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	// Execute the query
	rows, err := conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
