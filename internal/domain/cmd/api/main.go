package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure/handler"
	"github.com/user/project/internal/domain/infrastructure/middleware"
	"github.com/user/project/internal/domain/infrastructure/repository"
	"github.com/user/project/internal/domain/infrastructure/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	taskRepo := repository.NewInMemoryTaskRepository()
	taskService := application.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	engine := router.SetupRouter()

	engine.Use(gin.Recovery())
	engine.Use(middleware.LoggingMiddleware())

	engine.POST("/tasks", func(c *gin.Context) {
		taskHandler.CreateTask(c)
	})

	engine.GET("/tasks", func(c *gin.Context) {
		taskHandler.ListTasks(c)
	})

	engine.PUT("/tasks/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
			return
		}
		c.Set("taskID", id)
		taskHandler.UpdateTask(c)
	})

	engine.PATCH("/tasks/:id/complete", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
			return
		}
		c.Set("taskID", id)
		taskHandler.CompleteTask(c)
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not start server: %v\n", err)
	}
}