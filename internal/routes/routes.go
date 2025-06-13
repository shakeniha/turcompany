package routes

import (
	"github.com/gin-gonic/gin"
	//"turcompany/internal/middleware"

	//"database/sql"
	"turcompany/internal/handlers"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler, roleHandler *handlers.RoleHandler, leadHandler *handlers.LeadHandler, dealHandler *handlers.DealHandler, authHandler *handlers.AuthHandler) *gin.Engine {

	r.POST("/login", authHandler.Login)

	r.POST("/users/", userHandler.CreateUser)

	users := r.Group("/users")
	// users.Use(middleware.AuthMiddleware())
	{
		users.GET("/:id", userHandler.GetUserByID)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
		users.GET("/", userHandler.ListUsers)
	}

	roles := r.Group("/roles")
	// roles.Use(middleware.AuthMiddleware())
	{
		roles.POST("/", roleHandler.CreateRole)
		roles.GET("/:id", roleHandler.GetRoleByID)
		roles.PUT("/:id", roleHandler.UpdateRole)
		roles.DELETE("/:id", roleHandler.DeleteRole)
		roles.GET("/", roleHandler.ListRoles)
	}

	// Lead routes
	leads := r.Group("/leads")
	{
		leads.POST("/", leadHandler.Create)
		leads.GET("/:id", leadHandler.GetByID)
		leads.PUT("/:id", leadHandler.Update)
		leads.DELETE("/:id", leadHandler.Delete)
	}

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
