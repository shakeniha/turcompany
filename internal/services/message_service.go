package services

import (
	"context"
	"time"
	
        "turcompany/internal/models"       
	"turcompany/internal/repositories"
)

// MessageService определяет интерфейс для сервиса сообщений.
type MessageService interface {
	Send(ctx context.Context, msg *models.Message) (*models.Message, error)
	GetConversationHistory(ctx context.Context, userID, partnerID int64) ([]models.Message, error)
	GetConversations(ctx context.Context, userID int64) ([]models.Conversation, error)
}

// messageService - конкретная реализация.
type messageService struct {
	repo repositories.MessageRepository
}

// NewMessageService - конструктор.
func NewMessageService(repo repositories.MessageRepository) MessageService {
	return &messageService{
		repo: repo,
	}
}

// Send - бизнес-логика отправки сообщения.
func (s *messageService) Send(ctx context.Context, msg *models.Message) (*models.Message, error) {
	// Устанавливаем время отправки
	msg.SentAt = time.Now()

	// Здесь может быть дополнительная логика:
	// - Проверить, не заблокировал ли один пользователь другого
	// - Отправить real-time уведомление через WebSocket (если используется)

	// Сохраняем сообщение через репозиторий
	err := s.repo.Store(ctx, msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

// GetConversationHistory - логика получения истории переписки двух пользователей.
func (s *messageService) GetConversationHistory(ctx context.Context, userID, partnerID int64) ([]models.Message, error) {
	// В будущем здесь можно добавить логику "отметить сообщения как прочитанные"
	return s.repo.FindConversationHistory(ctx, userID, partnerID)
}

// GetConversations - логика получения списка всех диалогов пользователя.
func (s *messageService) GetConversations(ctx context.Context, userID int64) ([]models.Conversation, error) {
	// Этот метод самый сложный, он требует от репозитория специального запроса,
	// который сгруппирует сообщения по собеседникам и вернет последнее сообщение для каждого.
	return s.repo.FindConversations(ctx, userID)
}
