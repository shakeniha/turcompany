package app

import (
	"database/sql"
	"fmt"
	"log"
	"turcompany/internal/config"
	"turcompany/internal/handlers"
	"turcompany/internal/repositories"
	"turcompany/internal/routes"
	"turcompany/internal/services"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Run() {
	cfg := config.LoadConfig()
	dsn := cfg.Database.DSN
	port := cfg.Server.Port

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService()
	userHandler := handlers.NewUserHandler(userService, authService)

	roleRepo := repositories.NewRoleRepository(db)
	roleService := services.NewRoleService(roleRepo)
	roleHandler := handlers.NewRoleHandler(roleService)

	leadRepo := repositories.NewLeadRepository(db)
	leadService := services.NewLeadService(leadRepo)
	leadHandler := handlers.NewLeadHandler(leadService)

	dealRepo := repositories.NewDealRepository(db)
	dealService := services.NewDealService(dealRepo)
	dealHandler := handlers.NewDealHandler(dealService)

	authHandler := handlers.NewAuthHandler(userService, authService)

	documentRepo := repositories.NewDocumentRepository(db)
	documentService := services.NewDocumentService(documentRepo)
	documentHandler := handlers.NewDocumentHandler(documentService)

	taskRepo := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	messageRepo := repositories.NewMessageRepository(db)
	messageService := services.NewMessageService(messageRepo)
	messageHandler := handlers.NewMessageHandler(messageService)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	routes.SetupRoutes(
		router,
		userHandler,
		roleHandler,
		leadHandler,
		dealHandler,
		authHandler,
		documentHandler,
		taskHandler,
		messageHandler,
	)

	listenAddr := fmt.Sprintf(":%d", port)
	log.Printf("Server started on %s", listenAddr)
	if err := router.Run(listenAddr); err != nil {
		log.Fatal(err)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
