package storage

import (
	"errors"
)

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExist    = errors.New("url exists")
)

type Store interface {
	GetURL(alias string) (string, error)
	SaveURL(urlToSave string, alias string) error
	DeleteUrl(alias string) error
	Close() error
}
