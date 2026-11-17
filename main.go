package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure/middleware"
	"github.com/user/project/internal/domain/infrastructure/repository"
	"github.com/user/project/internal/domain/infrastructure/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	taskRepo := repository.NewInMemoryTaskRepository()
	taskService := application.NewTaskService(taskRepo)

	r := router.SetupRouter()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not start server: %v\n", err)
	}
}