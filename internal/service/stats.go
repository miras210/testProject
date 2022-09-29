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
	fullStats := &models.FullStats{
		Date:   stats.Date,
		Views:  stats.Views,
		Clicks: stats.Clicks,
		Cost:   stats.Cost,
		Cpc:    stats.Cost / float32(stats.Clicks),
		Cpm:    stats.Cost / float32(stats.Views) * 1000,
	}
	return s.statsRepo.CreateStats(fullStats)
}

func (s *StatsService) GetStats(from, to time.Time, filter *models.Filter) ([]*models.FullStats, error) {
	return s.statsRepo.GetStats(from, to, filter)
}

func (s *StatsService) DeleteStats() error {
	return s.statsRepo.DeleteStats()
}
