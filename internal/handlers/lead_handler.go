package handlers
import(
	"turcompany/internal/models"
	"turcompany/internal/services"
	"github.com/gin-gonic/gin"
)

type LeadHandler struct {
	Service *services.LeadService
}
func NewLeadHandler(service *services.LeadService) *LeadHandler {
	return &LeadHandler{Service: service}
}
func (h *LeadHandler) Create(c *gin.Context) {
	var lead models.Leads
	if err := c.ShouldBindJSON(&lead); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Create(&lead); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(201)
}
func (h *LeadHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var lead models.Leads
	if err := c.ShouldBindJSON(&lead); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	lead.ID = id
	if err := h.Service.Update(&lead); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(200)
}
func (h *LeadHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	lead, err := h.Service.GetByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Lead not found"})
		return
	}
	c.JSON(200, lead)
}
func (h *LeadHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}