package model

import (
	"errors"
	"fmt"
)

var (
	ErrPermissionDenied = errors.New("permission denied")
	ErrNotFound         = errors.New("not found")
	ErrAlreadyExist     = errors.New("already exist")
)

var (
	ErrTypeNotFound = fmt.Errorf("content type not found: %w", ErrNotFound)
	ErrItemNotFound = fmt.Errorf("item not found: %w", ErrNotFound)
)
