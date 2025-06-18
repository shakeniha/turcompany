package handlers

import (
	"net/http"
	"turcompany/internal/services"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	Service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{Service: service}
}

// @Summary Сводный отчет
// @Description Выводит общее количество лидов и сделок.
// @Tags Reports
// @Produce json
// @Success 200 {object} map[string]int
// @Failure 500 {object} map[string]string
// @Router /reports/summary [get]
func (h *ReportHandler) GetSummary(c *gin.Context) {
	data, err := h.Service.GetSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary Фильтрация лидов
// @Description Фильтрует лиды по статусу и owner_id.
// @Tags Reports
// @Produce json
// @Param status query string false "Статус"
// @Param owner_id query int false "ID владельца"
// @Success 200 {array} models.Leads
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reports/leads/filter [get]
func (h *ReportHandler) FilterLeads(c *gin.Context) {
	status := c.Query("status")
	ownerID := c.Query("owner_id")

	leads, err := h.Service.FilterLeads(status, ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, leads)
}

// @Summary Фильтрация сделок
// @Description Фильтрует сделки по статусу и дате.
// @Tags Reports
// @Produce json
// @Param status query string false "Статус"
// @Param from query string false "Дата с (yyyy-mm-dd)"
// @Param to query string false "Дата по (yyyy-mm-dd)"
// @Success 200 {array} models.Deals
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reports/deals/filter [get]
func (h *ReportHandler) FilterDeals(c *gin.Context) {
	status := c.Query("status")
	from := c.Query("from")
	to := c.Query("to")

	deals, err := h.Service.FilterDeals(status, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deals)
}
