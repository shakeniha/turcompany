package models

import (
	"time"
)

// Определяем возможные статусы задачи для консистентности
type TaskStatus string

const (
	StatusNew        TaskStatus = "new"
	StatusInProgress TaskStatus = "in_progress"
	StatusDone       TaskStatus = "done"
)

// Task - это наша основная бизнес-сущность.
// Она не содержит тегов для json или базы данных. Это чистая структура.
type Task struct {
	ID          int64
	CreatorID   int64
	AssigneeID  int64
	EntityID    int64
	EntityType  string
	Title       string
	Description string
	Status      TaskStatus
	DueDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
