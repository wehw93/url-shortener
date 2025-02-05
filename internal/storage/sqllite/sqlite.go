package sqllite

import (
	"database/sql"
	"errors"
	"fmt"
	"url-shortener/internal/storage"

	"github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS url(
			id INTEGER PRIMARY KEY,
			alias TEXT NOT NULL UNIQUE,
			url TEXT NOT NULL);
		CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}
	return &Storage{
		db: db,
	}, nil
}
func (s *Storage) SaveURL(urlToSave string, alias string) error {
	const op = "storage.sqlite.SaveURL"

	stmt, err := s.db.Prepare("INSERT INTO url(url, alias) VALUES(?, ?)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(urlToSave, alias)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return  fmt.Errorf("%s: %w", op, storage.ErrURLExist)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
func (s *Storage) GetURL(alias string) (string, error) {
	const op = "storage.sqlite.GetURL"
	stmt, err := s.db.Prepare("SELECT url FROM url WHERE alias = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	gettedUrl := ""
	err = stmt.QueryRow(alias).Scan(&gettedUrl)
	if errors.Is(err, sql.ErrNoRows) {
		return "", fmt.Errorf("%s %w", op, storage.ErrURLNotFound)
	}
	return gettedUrl, nil

}

func (s *Storage) DeleteUrl(alias string) error {
	const op = "storage.sqlite.GetURL"
	stmt, err := s.db.Prepare("DELETE FROM url WHERE alias = ?")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	res, err := stmt.Exec(alias)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	val, err := res.RowsAffected()
	if val == 0 {
		return fmt.Errorf("alias to delete not dound")
	}
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil

}

func (s *Storage) Close() error {
	return s.db.Close()
}