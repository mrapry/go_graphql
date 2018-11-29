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

// WriteMysqlDB function for creating database connection for write-access
func WriteMysqlDB() *sql.DB {
	return CreateDBConnection(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")))

}

// create connection in database
func CreateDBConnection(descriptor string) *sql.DB {
	db, err := sql.Open("mysql", descriptor)
	if err != nil {
		print(err.Error())
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
