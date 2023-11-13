package sqlitedb

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type service struct {
	db *sql.DB
}

func NewSqlite(dbName string) *service {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS wiki_pages (
		title TEXT PRIMARY KEY,
        pageid INTEGER,
        extract TEXT,
        summary TEXT,
        eli5 TEXT
    );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating wiki_pages table: %v", err)
	}

	return &service{db: db}
}
