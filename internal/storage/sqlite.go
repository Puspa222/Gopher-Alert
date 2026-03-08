package storage

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	conn *sql.DB
}

func NewSQLiteDB(filepath string) (*DB, error) {
	conn, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite database: %w ", err)
	}
	query := `
	CREATE TABLE IF NOT EXISTS notifications(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		provider TEXT,
		recipient TEXT,
		message TEXT,
		status TEXT,
		error TEXT,
		created_at DATETIME
	);
	`
	if _, err := conn.Exec(query); err != nil {
		return nil, fmt.Errorf("failed to create notifications table: %w", err)
	}
	return &DB{conn: conn}, nil
}

func (db *DB) LogNotification(provider, recipent, message, status, errMsg string) error {
	query := `
	INSERT INTO notifications(provider, recipient, message, status, error, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := db.conn.Exec(query, provider, recipent, message, status, errMsg, time.Now())
	return err
}
