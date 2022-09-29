package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"testProject/internal/models"
	"time"
)

type Repository struct {
	Stats
}

type Stats interface {
	CreateStats(stats *models.Stats) error
	GetStats(from, to time.Time) ([]*models.Stats, error)
	DeleteStats() error
}

func NewRepository(db *sql.DB, timeout time.Duration, logger *zap.SugaredLogger) *Repository {
	return &Repository{
		Stats: newStatsRepo(db, timeout, logger),
	}
}
