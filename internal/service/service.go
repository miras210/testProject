package service

import (
	"go.uber.org/zap"
	"testProject/config"
	"testProject/internal/models"
	"testProject/internal/repository"
	"time"
)

type Service struct {
	Stats
}

type Stats interface {
	CreateStats(stats *models.Stats) error
	GetStats(from, to time.Time) ([]*models.Stats, error)
	DeleteStats() error
}

func NewService(repo *repository.Repository, cfg *config.Configs, logger *zap.SugaredLogger) *Service {
	return &Service{
		Stats: NewStatsService(repo, cfg, logger),
	}
}
