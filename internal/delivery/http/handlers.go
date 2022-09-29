package delivery

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"testProject/config"
	"testProject/internal/models"
	"testProject/internal/service"
	"time"
)

const dateLayout = "2006-01-02"
const defaultFrom = "2000-01-01"
const defaultTo = "3000-01-01"

type Handler struct {
	service *service.Service
	logger  *zap.SugaredLogger
	cfg     *config.Configs
}

type StatsGetOutput struct {
	Date   models.CustomTime `json:"date"`
	Views  int               `json:"views"`
	Clicks int               `json:"clicks"`
	Cost   float32           `json:"cost"`
	Cpc    float32           `json:"cpc"`
	Cpm    float32           `json:"cpm"`
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
	router.POST("/statistics", h.CreateStats)
	router.GET("/statistics", h.GetStats)
	router.DELETE("/statistics", h.DeleteStats)

	return router
}

func (h *Handler) CreateStats(c *gin.Context) {
	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		h.logger.Errorf("Error occured while reading request body: %s", err.Error())
		c.JSON(401, gin.H{
			"error": models.ErrInvalidInput.Error(),
		})
		return
	}
	var statsRequest *models.Stats
	err = json.Unmarshal(requestBody, &statsRequest)
	if err != nil {
		h.logger.Errorf("Error occurred while unmarshalling request body: %s", err.Error())
		c.JSON(401, gin.H{
			"error": models.ErrInvalidInput.Error(),
		})
		return
	}
	h.logger.Infof("Unmarshalled object: %v", statsRequest)
	err = h.service.CreateStats(statsRequest)
	if err != nil {
		h.logger.Errorf("Error occured while creating statistics: %s", err.Error())
		c.JSON(500, gin.H{
			"error": models.ErrInternalServerError.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func (h *Handler) GetStats(c *gin.Context) {
	fromStr := c.DefaultQuery("from", defaultFrom)
	toStr := c.DefaultQuery("to", defaultTo)

	from, err := time.Parse(dateLayout, fromStr)
	if err != nil {
		h.logger.Errorf("Error occurred while parsing query params: %s", err.Error())
		c.JSON(401, gin.H{
			"error": models.ErrInvalidInput.Error(),
		})
		return
	}
	to, err := time.Parse(dateLayout, toStr)
	if err != nil {
		h.logger.Errorf("Error occurred while parsing query params: %s", err.Error())
		c.JSON(401, gin.H{
			"error": models.ErrInvalidInput.Error(),
		})
		return
	}

	stats, err := h.service.GetStats(from, to)
	if err != nil {
		h.logger.Errorf("Error occurred while getting statistics: %s", err.Error())
		c.JSON(500, gin.H{
			"error": models.ErrInternalServerError.Error(),
		})
		return
	}
	statsOutputs := make([]*StatsGetOutput, 0)
	for _, v := range stats {
		statsOutput := &StatsGetOutput{}
		statsOutput.Date = v.Date
		statsOutput.Views = v.Views
		statsOutput.Clicks = v.Clicks
		statsOutput.Cost = v.Cost
		statsOutput.Cpc = v.Cost / float32(v.Clicks)
		statsOutput.Cpm = v.Cost / float32(v.Views) * 1000

		statsOutputs = append(statsOutputs, statsOutput)
	}
	c.JSON(200, gin.H{
		"statistics": statsOutputs,
	})
}

func (h *Handler) DeleteStats(c *gin.Context) {
	err := h.service.DeleteStats()
	if err != nil {
		h.logger.Errorf("Error occurred while deleting statistics: %s", err.Error())
		c.JSON(500, gin.H{
			"error": models.ErrInternalServerError.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "OK",
	})
}
