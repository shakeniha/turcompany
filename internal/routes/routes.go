package routes

import (
	"github.com/gin-gonic/gin"
	//"database/sql"
	"turcompany/internal/handlers"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler, roleHandler *handlers.RoleHandler) *gin.Engine {

	users := r.Group("/users")
	{
		users.POST("/", userHandler.CreateUser)
		users.GET("/:id", userHandler.GetUserByID)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
		users.GET("/", userHandler.ListUsers)
	}

	roles := r.Group("/roles")
	{
		roles.POST("/", roleHandler.CreateRole)
		roles.GET("/:id", roleHandler.GetRoleByID)
		roles.PUT("/:id", roleHandler.UpdateRole)
		roles.DELETE("/:id", roleHandler.DeleteRole)
		roles.GET("/", roleHandler.ListRoles)
	}
	//
	//leadRepo := repositories.NewLeadRepository(db)
	//leadService := services.NewLeadService(leadRepo)
	//leadHandler := handlers.NewLeadHandler(leadService)
	//// Lead routes
	//leads := r.Group("/leads")
	//{
	//	leads.POST("/", leadHandler.Create)
	//	leads.GET("/:id", leadHandler.GetByID)
	//	leads.PUT("/:id", leadHandler.Update)
	//	leads.DELETE("/:id", leadHandler.Delete)
	//}
	//
	//dealRepo := repositories.NewDealRepository(db)
	//dealService := services.NewDealService(dealRepo)
	//dealHandler := handlers.NewDealHandler(dealService)
	////Deal routes
	//deals := r.Group("/deals")
	//{
	//	deals.POST("/", dealHandler.Create)
	//	deals.GET("/:id", dealHandler.GetByID)
	//	deals.PUT("/:id", dealHandler.Update)
	//	deals.DELETE("/:id", dealHandler.Delete)
	//}

	return r
}
