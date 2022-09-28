package delivery

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"testProject/config"
	"testProject/internal/service"
)

type Handler struct {
	service *service.Service
	logger  *zap.SugaredLogger
	cfg     *config.Configs
}

func NewHandler(services *service.Service, logger *zap.SugaredLogger, cfg *config.Configs) *Handler {
	return &Handler{
		service: services,
		logger:  logger,
		cfg:     cfg,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	return router
}
