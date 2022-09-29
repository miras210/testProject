package service

import (
	"go.uber.org/zap"
	"testProject/config"
	"testProject/internal/models"
	"testProject/internal/repository"
	"time"
)

type StatsService struct {
	statsRepo repository.Stats
	cfg       *config.Configs
	log       *zap.SugaredLogger
}

func NewStatsService(statsRepo repository.Stats, cfg *config.Configs, log *zap.SugaredLogger) *StatsService {
	return &StatsService{
		statsRepo: statsRepo,
		cfg:       cfg,
		log:       log,
	}
}

func (s *StatsService) CreateStats(stats *models.Stats) error {
	return s.statsRepo.CreateStats(stats)
}

func (s *StatsService) GetStats(from, to time.Time) ([]*models.Stats, error) {
	return s.statsRepo.GetStats(from, to)
}

func (s *StatsService) DeleteStats() error {
	return s.statsRepo.DeleteStats()
}
