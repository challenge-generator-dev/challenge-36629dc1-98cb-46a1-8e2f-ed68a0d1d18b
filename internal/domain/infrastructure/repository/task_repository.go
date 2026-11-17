// TODO(reto): Implementar el repositorio en memoria para tareas.
// Requisitos:
// - Almacenar tareas en un mapa con UUID como clave.
// - Implementar métodos para guardar, listar y buscar tareas.
// - Manejar concurrencia con sync.RWMutex.
package repository

import (
	"sync"

	"github.com/google/uuid"
	"github.com/user/project/internal/domain"
)

type InMemoryTaskRepository struct {
	tasks map[uuid.UUID]*domain.Task
	mu    sync.RWMutex
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[uuid.UUID]*domain.Task),
	}
}

// TODO(reto): Implementar Save para guardar una tarea.
func (r *InMemoryTaskRepository) Save(task *domain.Task) error {
	return nil
}

// TODO(reto): Implementar FindAll para listar todas las tareas.
func (r *InMemoryTaskRepository) FindAll() ([]*domain.Task, error) {
	return nil, nil
}

// TODO(reto): Implementar FindByID para buscar una tarea por ID.
func (r *InMemoryTaskRepository) FindByID(id uuid.UUID) (*domain.Task, error) {
	return nil, nil
}