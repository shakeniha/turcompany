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

// @Summary      Сводный отчет
// @Description  Выводит общее количество лидов и сделок.
// @Tags         Reports
// @Produce      json
// @Success      200  {object}  map[string]int  "Количество лидов и сделок"
// @Failure      500  {object}  map[string]string
// @Router       /reports/summary [get]
func (h *ReportHandler) GetSummary(c *gin.Context) {
	data, err := h.Service.GetSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary      Фильтр сделок
// @Description  Фильтрация сделок по указанным параметрам (например, статус, дата).
// @Tags         Reports
// @Produce      json
// @Param        status  query     string  false  "Статус сделки"
// @Param        from    query     string  false  "Начало диапазона дат (yyyy-mm-dd)"
// @Param        to      query     string  false  "Конец диапазона дат (yyyy-mm-dd)"
// @Success      200     {array}   models.Deals   "Список сделок"
// @Failure      400     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /reports/deals/filter [get]
func (h *ReportHandler) FilterDeals(c *gin.Context) {
	status := c.Query("status")
	fromDate := c.Query("from")
	toDate := c.Query("to")

	results, err := h.Service.FilterDeals(status, fromDate, toDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}
