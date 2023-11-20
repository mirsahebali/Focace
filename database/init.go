package database

import (
	"database/sql"
	"fmt"

	"mirsahebali/focace/helpers/parsers"

	_ "github.com/lib/pq"
)

var (
	user = parsers.GetEnvVars("DB_USER")
	pass = parsers.GetEnvVars("DB_PASS")
	host = parsers.GetEnvVars("DB_HOST")
	port = parsers.GetEnvVars("DB_PORT")
	name = parsers.GetEnvVars("DB_NAME")
)

func InitDB() *sql.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, pass, name, host, port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error with driver")
		fmt.Println(err)
	}
	// WARN: Remove this later while refactoring
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database")
		fmt.Println(err)
	}
	fmt.Println("Database Connected")
	return db
}
