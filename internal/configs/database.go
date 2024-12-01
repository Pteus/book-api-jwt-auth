package configs

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDatabase() *sql.DB {
	connstr := "user=postgres password=gobank dbname=books sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("database not reachable: %v", err)
	}

	statement := `
	CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
	);`

	_, err = db.Exec(statement)
	if err != nil {
		log.Fatalf("Failed to create 'users' table")
	}

	return db
}
