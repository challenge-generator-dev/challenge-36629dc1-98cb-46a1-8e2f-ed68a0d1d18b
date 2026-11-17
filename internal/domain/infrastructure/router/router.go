// TODO(reto): archivo stubbeado; implementa los cuerpos.
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure/handler"
	"github.com/user/project/internal/domain/infrastructure/middleware"
	"github.com/user/project/internal/domain/infrastructure/repository"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Aplicar middleware de logging global
	router.Use(middleware.LoggingMiddleware())

	// Inicializar repositorio
	taskRepo := repository.NewInMemoryTaskRepository()

	// Inicializar servicio
	taskService := application.NewTaskService(taskRepo)

	// Inicializar handler
	taskHandler := handler.NewTaskHandler(taskService)

	// Definir rutas
	tasks := router.Group("/tasks")
	{
		tasks.POST("", taskHandler.CreateTask)
		tasks.GET("", taskHandler.ListTasks)
		tasks.PUT("/:id", taskHandler.UpdateTask)
		tasks.PATCH("/:id/complete", taskHandler.CompleteTask)
	}

	return router
}