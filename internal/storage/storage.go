package storage

import "errors"

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLEXISTS = errors.New("url exists")
	
)
