package db

import (
	"database/sql"
)

// execMutation runs a mutation SQL query (INSERT, UPDATE, DELETE) with any number of parameters.
func Mutate(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

