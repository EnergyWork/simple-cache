package internal

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrNotInitialized = errors.New("cache is not initialized")
	ErrExpired        = errors.New("expired")
)
