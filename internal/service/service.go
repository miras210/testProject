package service

import (
	"go.uber.org/zap"
	"testProject/config"
	"testProject/internal/repository"
)

type Service struct {
}

func NewService(repo *repository.Repository, cfg *config.Configs, logger *zap.SugaredLogger) *Service {
	return &Service{}
}
