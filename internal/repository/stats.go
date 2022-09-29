package repository

import (
	"context"
	"database/sql"
	"fmt"
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

func (s *StatsRepository) CreateStats(stats *models.FullStats) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	query := `INSERT INTO stats VALUES ($1, $2, $3, $4, $5, $6) RETURNING date`
	if err := s.db.QueryRowContext(ctx, query, stats.Date.Time, stats.Views,
		stats.Clicks, stats.Cost, stats.Cpc, stats.Cpm).Scan(&stats.Date.Time); err != nil {
		s.logger.Errorf("Error occured while querying to DB: %s", err.Error())
		return err
	}
	s.logger.Infof("Statistics successfully inserted with a date: %s", stats.Date)
	return nil
}

func (s *StatsRepository) GetStats(from, to time.Time, filter *models.Filter) ([]*models.FullStats, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	s.logger.Infof("Filter: %s", filter)
	query := fmt.Sprintf(`SELECT * FROM stats WHERE date BETWEEN $1 AND $2 ORDER BY $3 %s`, filter.SortAsc)
	rows, err := s.db.QueryContext(ctx, query, from, to, filter.SortColumn)
	if err != nil {
		s.logger.Errorf("Error during getting statistics: %s", err.Error())
		return nil, models.ErrDBError
	}
	defer rows.Close()

	stats := make([]*models.FullStats, 0)
	for rows.Next() {
		var stat models.FullStats
		if err := rows.Scan(&stat.Date.Time, &stat.Views, &stat.Clicks, &stat.Cost, &stat.Cpc, &stat.Cpm); err != nil {
			return nil, models.ErrDBError
		}
		s.logger.Infof("Stat: %v", stat)
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
