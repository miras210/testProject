package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"time"
)

type Repository struct {
}

func NewRepository(db *sql.DB, timeout time.Duration, logger *zap.SugaredLogger) *Repository {
	return &Repository{}
}
