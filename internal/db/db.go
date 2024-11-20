package db

import (
	"database/sql"
)

func CreateTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS urls (
    id Integer NOT NULL PRIMARY KEY,
    short_url TEXT NOT NULL,
    original_url TEXT NOT NULL);`

	_, err := db.Exec(query)
	return err
}

func StoreURL(db *sql.DB, shortURL string, originalURL string) error {
	query := `INSERT INTO urls (short_url, original_url) VALUES (?, ?)`
	_, err := db.Exec(query, shortURL, originalURL)
	return err
}

func GetOriginalURL(db *sql.DB, shortURL string) (string, error) {
	var original string
	query := `SELECT original_url FROM urls WHERE short_url = ?`
	err := db.QueryRow(query, shortURL).Scan(&original)
	if err != nil {
		return "", err
	}
	return original, nil
}
