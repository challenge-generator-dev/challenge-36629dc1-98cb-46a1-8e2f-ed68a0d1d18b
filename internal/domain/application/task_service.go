package application

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/user/project/internal/domain"
	"github.com/user/project/internal/domain/infrastructure/repository"
)

var (
	ErrTaskNotFound  = errors.New("tarea no encontrada")
	ErrEmptyTitle    = errors.New("el título no puede estar vacío")
	ErrTitleTooLong  = errors.New("el título no puede exceder 100 caracteres")
)

const maxTitleLength = 100

type TaskService struct {
	repo *repository.InMemoryTaskRepository
}

func NewTaskService(repo *repository.InMemoryTaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(title string) (*domain.Task, error) {
	trimmedTitle := strings.TrimSpace(title)
	if trimmedTitle == "" {
		return nil, ErrEmptyTitle
	}
	if len(trimmedTitle) > maxTitleLength {
		return nil, ErrTitleTooLong
	}

	task, err := domain.NewTask(trimmedTitle)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) GetAllTasks() ([]*domain.Task, error) {
	tasks, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) UpdateTaskTitle(id uuid.UUID, newTitle string) (*domain.Task, error) {
	trimmedTitle := strings.TrimSpace(newTitle)
	if trimmedTitle == "" {
		return nil, ErrEmptyTitle
	}
	if len(trimmedTitle) > maxTitleLength {
		return nil, ErrTitleTooLong
	}

	task, err := s.repo.FindByID(id)
	if err != nil {
		return nil, ErrTaskNotFound
	}

	if err := task.UpdateTitle(trimmedTitle); err != nil {
		return nil, err
	}

	if err := s.repo.Save(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) CompleteTask(id uuid.UUID) (*domain.Task, error) {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return nil, ErrTaskNotFound
	}

	if err := task.Complete(); err != nil {
		return nil, err
	}

	if err := s.repo.Save(task); err != nil {
		return nil, err
	}

	return task, nil
}