package handlers

import (
	"net/http"
	"strconv"
	"turcompany/internal/models"
	"turcompany/internal/services"

	"github.com/gin-gonic/gin"
)

type DealHandler struct {
	Service *services.DealService
}

func NewDealHandler(service *services.DealService) *DealHandler {
	return &DealHandler{Service: service}
}

// @Summary      Создание сделки
// @Description  Создает новую сделку, связанную с лидом
// @Tags         Deals
// @Accept       json
// @Produce      json
// @Param        deals  body      models.Deals  true  "Данные сделки"
// @Success      201   {object}  models.Deals
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /deals [post]
func (h *DealHandler) Create(c *gin.Context) {
	var deal models.Deals
	if err := c.ShouldBindJSON(&deal); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if _, err := h.Service.Create(&deal); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(201)
}

// @Summary      Обновление сделки
// @Description  Обновляет данные сделки по ее ID.
// @Tags         Deals
// @Accept       json
// @Produce      json
// @Param        id    path      int           true  "ID сделки"
// @Param        deal  body      models.Deals  true  "Новые данные сделки"
// @Success      200   {object}  models.Deals
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /deals/{id} [put]
func (h *DealHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var deal models.Deals
	if err := c.ShouldBindJSON(&deal); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	deal.ID = id

	if err := h.Service.Update(&deal); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(200)
}

// @Summary      Получить сделку по ID
// @Description  Возвращает данные одной сделки
// @Tags         Deals
// @Produce      json
// @Param        id   path      int  true  "ID сделки"
// @Success      200  {object}  models.Deals
// @Failure      404  {object}  map[string]string
// @Router       /deals/{id} [get]
func (h *DealHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	deal, err := h.Service.GetByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Deal not found"})
		return
	}
	c.JSON(200, deal)
}

// @Summary      Удалить сделку
// @Description  Удаляет сделку по ID
// @Tags         Deals
// @Param        id   path  int  true  "ID сделки"
// @Success      204  "No Content"
// @Failure      500  {object}  map[string]string
// @Router       /deals/{id} [delete]
func (h *DealHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.Service.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}

// List Deals
// @Summary      Get all deals
// @Description  Returns a list of all deals
// @Tags         Deals
// @Produce      json
// @Success      200  {array}  models.Deals
// @Failure 500  {object}  map[string]string
// @Router       /deals/ [get]
func (h *DealHandler) List(c *gin.Context) {
	deals, err := h.Service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve deals"})
		return
	}
	c.JSON(http.StatusOK, deals)
}
