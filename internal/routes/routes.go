package routes

import (
	"github.com/gin-gonic/gin"
	//"turcompany/internal/middleware"

	//"database/sql"
	"turcompany/internal/handlers"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler, roleHandler *handlers.RoleHandler, leadHandler *handlers.LeadHandler, dealHandler *handlers.DealHandler, authHandler *handlers.AuthHandler, documentHandler *handlers.DocumentHandler, taskHandler *handlers.TaskHandler, messageHandler *handlers.MessageHandler, smsHandler *handlers.SMSHandler) *gin.Engine {

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

	//Document routes
	documents := r.Group("/documents")
	{
		documents.POST("/", documentHandler.CreateDocument)
		documents.GET("/:id", documentHandler.GetDocument)
		documents.DELETE("/:id", documentHandler.DeleteDocument)

		documents.GET("/deal/:dealid", documentHandler.ListDocumentsByDeal)
		documents.PUT("/verify/:id", documentHandler.VerifyDocument)
		documents.PUT("/send/:id/:code", documentHandler.SendSMSConfirmation)
		documents.PUT("/confirm/:id/:code", documentHandler.ConfirmDocument)
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

	//SMS routes
	sms := r.Group("/sms")
	{
		sms.POST("/send", smsHandler.SendSMSHandler)
		sms.POST("/resend", smsHandler.ResendSMSHandler)
		sms.POST("/confirm", smsHandler.ConfirmSMSHandler)
		sms.GET("/latest/:document_id", smsHandler.GetLatestSMSHandler)
		sms.DELETE("/:document_id", smsHandler.DeleteSMSHandler)
	}

	return r
}
