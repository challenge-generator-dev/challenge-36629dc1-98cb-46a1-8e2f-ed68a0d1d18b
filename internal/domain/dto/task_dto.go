package dto

import (
	"time"

	"github.com/google/uuid"
)

type TaskCreateDTO struct {
	Title string `json:"title" binding:"required,max=100"`
}

type TaskUpdateDTO struct {
	Title string `json:"title" binding:"required,max=100"`
}

type TaskResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTaskResponseDTO(id uuid.UUID, title string, status string, createdAt, updatedAt time.Time) TaskResponseDTO {
	return TaskResponseDTO{
		ID:        id,
		Title:     title,
		Status:    status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}