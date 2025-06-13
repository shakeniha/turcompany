package handlers

import (
	"strconv"
	"turcompany/internal/models"
	"turcompany/internal/services"

	"github.com/gin-gonic/gin"
)

type DocumentHandler struct {
	Service *services.DocumentService
}

func NewDocumentHandler(service *services.DocumentService) *DocumentHandler {
	return &DocumentHandler{Service: service}
}

func (h *DocumentHandler) CreateDocument(c *gin.Context) {
	var document models.Document

	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.Service.CreateDocument(&document)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"id": id})
}

func (h *DocumentHandler) GetDocument(c *gin.Context) {
	idparam := c.Param("id")

	id, err := strconv.ParseInt(idparam, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	doc, err := h.Service.GetDocument(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "document not found"})
		return
	}

	c.JSON(200, doc)
}

func (h *DocumentHandler) VerifyDocument(c *gin.Context) {
	idparam := c.Param("id")
	id, err := strconv.ParseInt(idparam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	if err := h.Service.VerifyDocument(id); err != nil {
		c.JSON(404, gin.H{"error": "document not found"})
		return
	}
	c.JSON(202, gin.H{"message": "verified"})
}

func (h *DocumentHandler) SendSMSConfirmation(c *gin.Context) {
	code := c.Param("code")
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	if err := h.Service.SendSMSConfirmation(id, code); err != nil {
		c.JSON(500, gin.H{"sms": "sms not confirmed"})
		return
	}

	c.JSON(200, gin.H{"sms": "sms confirmed"})
}

func (h *DocumentHandler) ConfirmDocument(c *gin.Context) {
	idParam := c.Param("id")
	code := c.Param("code")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	if err := h.Service.ConfirmDocument(id, code); err != nil {
		c.JSON(404, gin.H{"error": "confirmation failed", "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "document signed"})
}

func (h *DocumentHandler) ListDocumentsByDeal(c *gin.Context) {
	dealIDParam := c.Param("dealid")
	dealID, err := strconv.ParseInt(dealIDParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid deal id"})
		return
	}

	docs, err := h.Service.ListDocumentsByDeal(dealID)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not fetch documents"})
		return
	}

	c.JSON(200, docs)
}

func (h *DocumentHandler) DeleteDocument(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	if err := h.Service.DeleteDocument(id); err != nil {
		c.JSON(500, gin.H{"error": "failed to delete document"})
		return
	}

	c.Status(204)
}
