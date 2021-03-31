package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MaxIdleConnection = 10
	MaxOpenConnection = 10
)

// WritePostgresDB function for creating database connection for write-access
func WriteMySQLDB() *sql.DB {
	return CreateDBConnection(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			os.Getenv("WRITE_DB_USER"),
			os.Getenv("WRITE_DB_PASSWORD"),
			os.Getenv("WRITE_DB_HOST"),
			os.Getenv("WRITE_DB_PORT"),
			os.Getenv("WRITE_DB_NAME"),
		))
}

// ReadPostgresDB function for creating database connection for write-access
func ReadMySQLDB() *sql.DB {
	return CreateDBConnection(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("READ_DB_USER"),
		os.Getenv("READ_DB_PASSWORD"),
		os.Getenv("READ_DB_HOST"),
		os.Getenv("READ_DB_PORT"),
		os.Getenv("READ_DB_NAME"),
	))

}

// CreateDBConnection function for creating database connection
func CreateDBConnection(descriptor string) *sql.DB {
	db, err := sql.Open("mysql", descriptor)
	if err != nil {
		defer db.Close()
		return db
	}

	db.SetMaxIdleConns(MaxIdleConnection)
	db.SetMaxOpenConns(MaxOpenConnection)

	return db
}

// CloseDb function for closing database connection
func CloseDb(db *sql.DB) {
	if db != nil {
		db.Close()
		db = nil
	}
}
