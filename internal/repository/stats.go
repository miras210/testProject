package repository

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
	"testProject/internal/models"
	"time"
)

type StatsRepository struct {
	db      *sql.DB
	timeout time.Duration
	logger  *zap.SugaredLogger
}

func newStatsRepo(db *sql.DB, timeout time.Duration, logger *zap.SugaredLogger) *StatsRepository {
	return &StatsRepository{
		db:      db,
		timeout: timeout,
		logger:  logger,
	}
}

func (s *StatsRepository) CreateStats(stats *models.Stats) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	query := `INSERT INTO stats VALUES ($1, $2, $3, $4) RETURNING stats_date`
	if err := s.db.QueryRowContext(ctx, query, stats.Date.Time, stats.Views, stats.Clicks, stats.Cost).Scan(&stats.Date.Time); err != nil {
		s.logger.Errorf("Error occured while querying to DB: %s", err.Error())
		return err
	}
	s.logger.Infof("Statistics successfully inserted with a date: %s", stats.Date)
	return nil
}

func (s *StatsRepository) GetStats(from, to time.Time) ([]*models.Stats, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	query := `SELECT * FROM stats WHERE stats_date BETWEEN $1 AND $2 ORDER BY stats_date DESC`
	rows, err := s.db.QueryContext(ctx, query, from, to)
	if err != nil {
		return nil, models.ErrDBError
	}
	defer rows.Close()

	stats := make([]*models.Stats, 0)
	for rows.Next() {
		var stat models.Stats
		if err := rows.Scan(&stat.Date.Time, &stat.Views, &stat.Clicks, &stat.Cost); err != nil {
			return nil, models.ErrDBError
		}
		stats = append(stats, &stat)
	}
	if err := rows.Err(); err != nil {
		return nil, models.ErrDBError
	}

	return stats, nil
}

func (s *StatsRepository) DeleteStats() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	query := `DELETE FROM stats`
	_, err := s.db.ExecContext(ctx, query)
	if err != nil {
		return models.ErrDBError
	}
	return nil
}
