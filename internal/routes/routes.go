package routes

import (
	"database/sql"
	"turcompany/internal/handlers"
	"turcompany/internal/repositories"
	"turcompany/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) *gin.Engine {
	r := gin.Default()

	leadRepo := repositories.NewLeadRepository(db)
	leadService := services.NewLeadService(leadRepo)
	leadHandler := handlers.NewLeadHandler(leadService)
	// Lead routes
	leads := r.Group("/leads")
	{
		leads.POST("/", leadHandler.Create)
		leads.GET("/:id", leadHandler.GetByID)
		leads.PUT("/:id", leadHandler.Update)
		leads.DELETE("/:id", leadHandler.Delete)
	}

	dealRepo := repositories.NewDealRepository(db)
	dealService := services.NewDealService(dealRepo)
	dealHandler := handlers.NewDealHandler(dealService)
	//Deal routes
	deals := r.Group("/deals")
	{
		deals.POST("/", dealHandler.Create)
		deals.GET("/:id", dealHandler.GetByID)
		deals.PUT("/:id", dealHandler.Update)
		deals.DELETE("/:id", dealHandler.Delete)
	}


	
	return r
}
