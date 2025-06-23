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

// @Summary      Создание документа
// @Description  Создает новый документ
// @Tags         Documents
// @Accept       json
// @Produce      json
// @Param        document  body      models.Document  true  "Данные документа"
// @Success      201  {object}  map[string]int64
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /documents [post]
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

// @Summary      Получить документ по ID
// @Description  Возвращает один документ по его ID
// @Tags         Documents
// @Produce      json
// @Param        id   path      int64  true  "ID документа"
// @Success      200  {object}  models.Document
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /documents/{id} [get]
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

// @Summary      Верифицировать документ
// @Description  Отмечает документ как верифицированный
// @Tags         Documents
// @Param        id   path  int64  true  "ID документа"
// @Success      202  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /documents/{id}/verify [patch]
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

// @Summary      Отправка SMS подтверждения
// @Description  Отправляет код подтверждения на номер, связанный с документом
// @Tags         Documents
// @Param        id    path  int64  true  "ID документа"
// @Param        code  path  string  true  "Код подтверждения"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /documents/{id}/sms/{code} [post]
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

// @Summary      Подтвердить документ
// @Description  Подтверждает документ по SMS-коду
// @Tags         Documents
// @Param        id    path  int64   true  "ID документа"
// @Param        code  path  string  true  "Код подтверждения"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /documents/{id}/confirm/{code} [post]
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

// @Summary      Получить документы сделки
// @Description  Возвращает все документы, связанные с определенной сделкой
// @Tags         Documents
// @Produce      json
// @Param        dealid  path  int64  true  "ID сделки"
// @Success      200  {array}   models.Document
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /documents/deal/{dealid} [get]
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

// @Summary      Удалить документ
// @Description  Удаляет документ по ID
// @Tags         Documents
// @Param        id   path  int64  true  "ID документа"
// @Success      204  "No Content"
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /documents/{id} [delete]
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

// @Summary      Создание документа из лида
// @Description  Создает документ на основе лида и типа документа
// @Tags         Documents
// @Accept       json
// @Produce      json
// @Param        input  body  object{lead_id=int,doc_type=string}  true  "ID лида и тип документа"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /documents/from-lead [post]
func (h *DocumentHandler) CreateDocumentFromLead(c *gin.Context) {
	var request struct {
		LeadID  int    `json:"lead_id" binding:"required"`
		DocType string `json:"doc_type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Неверные параметры запроса: " + err.Error()})
		return
	}

	doc, err := h.Service.CreateDocumentFromLead(request.LeadID, request.DocType)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Ошибка при создании документа: " + err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message":  "Документ успешно создан",
		"document": doc,
	})
}
