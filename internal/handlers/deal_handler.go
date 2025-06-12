package handlers

import (
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

func (h *DealHandler) Create(c *gin.Context) {
	var deal models.Deals
	if err := c.ShouldBindJSON(&deal); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Create(&deal); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(201)
}

func (h *DealHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	var deal models.Deals
	if err := c.ShouldBindJSON(&deal); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(idStr)
	deal.ID = id
	if err := h.Service.Update(&deal); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(200)
}

func (h *DealHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	deal, err := h.Service.GetByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Deal not found"})
		return
	}
	c.JSON(200, deal)
}

func (h *DealHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}
