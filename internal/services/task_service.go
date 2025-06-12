package services

import (
	"context"
	"time"

	"turcompany/internal/models"       // Путь к вашим моделям
	"turcompany/internal/repositories" // Путь к вашим репозиториям
)

// TaskService определяет интерфейс для сервиса задач.
// Использование интерфейсов - хорошая практика для тестирования.
type TaskService interface {
	Create(ctx context.Context, task *models.Task) (*models.Task, error)
	GetByID(ctx context.Context, id int64) (*models.Task, error)
	GetAll(ctx context.Context, filter models.TaskFilter) ([]models.Task, error)
	Update(ctx context.Context, task *models.Task) (*models.Task, error)
	Delete(ctx context.Context, id int64) error
}

// taskService - это конкретная реализация TaskService.
type taskService struct {
	repo repositories.TaskRepository
}

// NewTaskService - это конструктор для taskService.
// Он принимает репозиторий как зависимость (Dependency Injection).
func NewTaskService(repo repositories.TaskRepository) TaskService {
	return &taskService{
		repo: repo,
	}
}

// Create - бизнес-логика создания задачи.
func (s *taskService) Create(ctx context.Context, task *models.Task) (*models.Task, error) {
	// Устанавливаем значения по умолчанию
	task.Status = models.StatusNew // Предполагаем, что константы в пакете models
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	// Вызываем метод репозитория для сохранения в БД
	err := s.repo.Store(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// GetByID - бизнес-логика получения задачи по ID.
func (s *taskService) GetByID(ctx context.Context, id int64) (*models.Task, error) {
	return s.repo.FindByID(ctx, id)
}

// GetAll - бизнес-логика получения списка задач с фильтрами.
func (s *taskService) GetAll(ctx context.Context, filter models.TaskFilter) ([]models.Task, error) {
	return s.repo.FindAll(ctx, filter)
}

// Update - бизнес-логика обновления задачи.
func (s *taskService) Update(ctx context.Context, task *models.Task) (*models.Task, error) {
	// Перед обновлением можно добавить проверки прав доступа
	
	// Получаем существующую задачу, чтобы убедиться, что она есть
	existingTask, err := s.repo.FindByID(ctx, task.ID)
	if err != nil {
		return nil, err // Возвращаем ошибку, если задача не найдена
	}
	
	// Обновляем поля (здесь можно использовать более сложную логику, чтобы не затирать нужные данные)
	existingTask.Title = task.Title
	existingTask.Description = task.Description
	existingTask.AssigneeID = task.AssigneeID
	existingTask.Status = task.Status
	existingTask.DueDate = task.DueDate
	existingTask.UpdatedAt = time.Now()


	err = s.repo.Update(ctx, existingTask)
	if err != nil {
		return nil, err
	}

	return existingTask, nil
}

// Delete - бизнес-логика удаления задачи.
func (s *taskService) Delete(ctx context.Context, id int64) error {
	// Здесь также можно добавить проверку прав: может ли текущий пользователь удалять эту задачу
	return s.repo.Delete(ctx, id)
}
