package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain/application"
)

type TaskHandler struct {
	service *application.TaskService
}

func NewTaskHandler(service *application.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// CreateTask handles POST /tasks
// TODO(reto): Implementar la lógica de creación de tareas
func (h *TaskHandler) CreateTask(c *gin.Context) {
	panic("TODO(reto): CreateTask")
}

// ListTasks handles GET /tasks
// TODO(reto): Implementar la lógica de listado de tareas
func (h *TaskHandler) ListTasks(c *gin.Context) {
	panic("TODO(reto): ListTasks")
}

// UpdateTask handles PUT /tasks/:id
// TODO(reto): Implementar la lógica de actualización de tareas
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	panic("TODO(reto): UpdateTask")
}

// CompleteTask handles PATCH /tasks/:id/complete
// TODO(reto): Implementar la lógica de completado de tareas
func (h *TaskHandler) CompleteTask(c *gin.Context) {
	panic("TODO(reto): CompleteTask")
}