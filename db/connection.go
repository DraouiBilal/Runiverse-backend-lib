package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)


func ConnectDB(Host, User ,Password ,DBName, SSL string, Port int) (*sql.DB, error) {
	// PostgreSQL connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",Host,Port,User,Password,DBName,SSL)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Check if the connection works
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
