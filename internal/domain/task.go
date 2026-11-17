// TODO(reto): Implementar la entidad Task con sus reglas de dominio.
// Requisitos:
// - El título no puede estar vacío y debe tener máximo 100 caracteres.
// - El ID debe ser generado automáticamente como UUID.
// - La tarea debe tener un estado (pendiente, completada).
package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEmptyTitle    = errors.New("el título no puede estar vacío")
	ErrTitleTooLong  = errors.New("el título no puede exceder 100 caracteres")
	ErrInvalidStatus = errors.New("estado inválido")
)

type TaskStatus string

const (
	StatusPending   TaskStatus = "pendiente"
	StatusCompleted TaskStatus = "completada"
)

type Task struct {
	ID        uuid.UUID   `json:"id"`
	Title     string      `json:"title"`
	Status    TaskStatus  `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// TODO(reto): Implementar NewTask para crear una nueva tarea con validaciones.
func NewTask(title string) (*Task, error) {
	return nil, nil
}

// TODO(reto): Implementar UpdateTitle para actualizar el título con validaciones.
func (t *Task) UpdateTitle(newTitle string) error {
	return nil
}

// TODO(reto): Implementar Complete para marcar la tarea como completada.
func (t *Task) Complete() error {
	return nil
}