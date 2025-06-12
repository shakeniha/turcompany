package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"turcompany/internal/handlers"

)

// createTaskRequest определяет структуру для создания задачи.
// Мы используем DTO (Data Transfer Object), чтобы не привязывать API к доменной модели.
type createTaskRequest struct {
	AssigneeID  int64  `json:"assignee_id"`
	EntityID    int64  `json:"entity_id"`
	EntityType  string `json:"entity_type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"` // Принимаем строку, чтобы гибко парсить
}

// updateTaskRequest определяет структуру для полного обновления задачи.
type updateTaskRequest struct {
	AssigneeID  int64  `json:"assignee_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date"`
}

// taskResponse определяет структуру для ответа API.
type taskResponse struct {
	ID          int64     `json:"id"`
	CreatorID   int64     `json:"creator_id"`
	AssigneeID  int64     `json:"assignee_id"`
	EntityID    int64     `json:"entity_id"`
	EntityType  string    `json:"entity_type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// handleCreateTask обрабатывает запрос на создание новой задачи.
// POST /api/v1/tasks
func (h *Handler) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	var req createTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Валидация входных данных (здесь может быть более сложная логика)
	if req.Title == "" || req.AssigneeID == 0 {
		h.respondWithError(w, http.StatusBadRequest, "Title and assignee_id are required")
		return
	}

	// Получаем ID создателя из контекста (помещается туда middleware'ом аутентификации)
	creatorID, ok := r.Context().Value(userIDKey).(int64)
	if !ok {
		h.respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	dueDate, err := time.Parse(time.RFC3339, req.DueDate)
	if err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid due_date format, use RFC3339")
		return
	}
	
	// Создаем доменную модель
	task := &domain.Task{
		CreatorID:   creatorID,
		AssigneeID:  req.AssigneeID,
		EntityID:    req.EntityID,
		EntityType:  req.EntityType,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     dueDate,
	}

	// Вызываем бизнес-логику
	createdTask, err := h.taskUseCase.Create(r.Context(), task)
	if err != nil {
		// Здесь должна быть обработка конкретных ошибок из UseCase
		h.respondWithError(w, http.StatusInternalServerError, "Failed to create task")
		return
	}

	// Отправляем успешный ответ
	h.respondWithJSON(w, http.StatusCreated, toTaskResponse(createdTask))
}

// handleGetTaskByID обрабатывает запрос на получение задачи по ее ID.
// GET /api/v1/tasks/{id}
func (h *Handler) handleGetTaskByID(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	task, err := h.taskUseCase.GetByID(r.Context(), id)
	if err != nil {
		// Проверяем, является ли ошибка "не найдено"
		if errors.Is(err, domain.ErrNotFound) { // domain.ErrNotFound - кастомная ошибка
			h.respondWithError(w, http.StatusNotFound, "Task not found")
			return
		}
		h.respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	h.respondWithJSON(w, http.StatusOK, toTaskResponse(task))
}

// handleGetTasksList обрабатывает запрос на получение списка задач с фильтрами.
// GET /api/v1/tasks?assignee_id=12&status=new
func (h *Handler) handleGetTasksList(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filter := domain.TaskFilter{} // Предполагаем, что такая структура есть в домене

	if assigneeIDStr := query.Get("assignee_id"); assigneeIDStr != "" {
		if id, err := strconv.ParseInt(assigneeIDStr, 10, 64); err == nil {
			filter.AssigneeID = &id
		}
	}
	if status := query.Get("status"); status != "" {
		filter.Status = &status
	}
	// ... другие фильтры ...

	tasks, err := h.taskUseCase.GetAll(r.Context(), filter)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to get tasks list")
		return
	}

	// Конвертируем список доменных моделей в список DTO для ответа
	response := make([]taskResponse, len(tasks))
	for i, task := range tasks {
		response[i] = toTaskResponse(&task)
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// handleDeleteTask обрабатывает запрос на удаление задачи.
// DELETE /api/v1/tasks/{id}
func (h *Handler) handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	err = h.taskUseCase.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			h.respondWithError(w, http.StatusNotFound, "Task not found")
			return
		}
		// Здесь может быть ошибка domain.ErrForbidden, если пользователь не имеет прав
		h.respondWithError(w, http.StatusInternalServerError, "Failed to delete task")
		return
	}

	h.respondWithJSON(w, http.StatusNoContent, nil)
}

// toTaskResponse - это вспомогательная функция-конвертер из доменной модели в DTO.
func toTaskResponse(t *domain.Task) taskResponse {
	return taskResponse{
		ID:          t.ID,
		CreatorID:   t.CreatorID,
		AssigneeID:  t.AssigneeID,
		EntityID:    t.EntityID,
		EntityType:  t.EntityType,
		Title:       t.Title,
		Description: t.Description,
		Status:      string(t.Status),
		DueDate:     t.DueDate,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
