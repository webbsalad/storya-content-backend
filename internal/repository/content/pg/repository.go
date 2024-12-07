package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/webbsalad/storya-content-backend/internal/repository/content"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (content.Repository, error) {
	return &Repository{db: db}, nil
}
