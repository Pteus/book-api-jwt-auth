package configs

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDatabase() *sql.DB {
	connstr := "user=postgres password=gobank dbname=books sslmodel=disabled"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("database not reachable: %v", err)
	}

	return db
}
