package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const schema = `
CREATE TABLE IF NOT EXISTS coordinates(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	x INTEGER NOT NULL,
	y INTEGER NOT NULL,
	map_id INTEGER NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	FOREIGN KEY(map_id) REFERENCES map(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS map(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid TEXT UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);
`

func NewDatabase(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath+"?cache=shared&mode=rwc&_journal_mode=WAL&_busy_timeout=5000&_synchronous=NORMAL")
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func InitializeSchema(db *sql.DB) error {
	_, err := db.Exec(schema)
	return err
}

