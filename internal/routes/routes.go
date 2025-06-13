package routes

import (
	"github.com/gin-gonic/gin"
	//"database/sql"
	"turcompany/internal/handlers"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler, roleHandler *handlers.RoleHandler, leadHandler *handlers.LeadHandler, dealHandler *handlers.DealHandler, documentHandler *handlers.DocumentHandler, smsHandler *handlers.SMSHandler) *gin.Engine {

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

	//Document routes
	documents := r.Group("/documents")
	{
		documents.POST("/", documentHandler.CreateDocument)
		documents.GET("/:id", documentHandler.GetDocument)
		documents.DELETE("/:id", documentHandler.DeleteDocument)

		documents.GET("/deal/:dealid", documentHandler.ListDocumentsByDeal)
		documents.PUT("/:id/verify", documentHandler.VerifyDocument)
		documents.POST("/:id/sms/:code", documentHandler.SendSMSConfirmation)
		documents.POST("/:id/confirm/:code", documentHandler.ConfirmDocument)
	}

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
