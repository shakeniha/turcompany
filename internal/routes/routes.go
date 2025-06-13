package routes

import (
	"github.com/gin-gonic/gin"
	//"database/sql"
	"turcompany/internal/handlers"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler, roleHandler *handlers.RoleHandler, leadHandler *handlers.LeadHandler, dealHandler *handlers.DealHandler, taskHandler *handlers.TaskHandler, messageHandler *handlers.MessageHandler) *gin.Engine {

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

	// --- Роуты для Задач (Tasks) ---
	tasks := r.Group("/tasks")
	tasks.POST("/", taskHandler.Create)
	tasks.GET("/", taskHandler.GetAll)
	tasks.GET("/:id", taskHandler.GetByID)
	tasks.PUT("/:id", taskHandler.Update)
	tasks.DELETE("/:id", taskHandler.Delete)

	// --- Роуты для Сообщений (Messages) ---
	messages := r.Group("/messages")
	messages.POST("/", messageHandler.Send)
	messages.GET("/conversations", messageHandler.GetConversations)
	messages.GET("/history/:partner_id", messageHandler.GetConversationHistory)

	return r
}
