package db

import (
	"database/sql"
	"fmt"
)


func CreateTable(conn *sql.DB, name string,columns map[string][]string) (sql.Result, error) {
	query := fmt.Sprintf("CREATE TABLE %s (", name)

	for k,v := range columns {
		query += k
		for _, param := range v {
			query += " " + param
		}
		query += " ,"
	}

	query = query[:len(query)-2] + ");"

	res, err := conn.Exec(query)

	return res, err
}
