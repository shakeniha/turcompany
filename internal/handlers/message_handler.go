package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"turcompany/internal/handlers"
)

// sendMessageRequest определяет DTO для отправки нового сообщения.
type sendMessageRequest struct {
	ReceiverID int64  `json:"receiver_id"`
	Content    string `json:"content"`
}

// messageResponse - это DTO для ответа API, представляющий одно сообщение.
type messageResponse struct {
	ID         int64      `json:"id"`
	SenderID   int64      `json:"sender_id"`
	ReceiverID int64      `json:"receiver_id"`
	Content    string     `json:"content"`
	SentAt     time.Time  `json:"sent_at"`
	ReadAt     *time.Time `json:"read_at,omitempty"` // Поле будет опущено, если nil
}

// conversationResponse - это DTO для ответа API, представляющий один диалог в списке.
type conversationResponse struct {
	Partner      userShortResponse `json:"partner"`       // Информация о собеседнике
	LastMessage  messageResponse   `json:"last_message"`  // Последнее сообщение в диалоге
	UnreadCount  int               `json:"unread_count"`  // Количество непрочитанных сообщений
}

// userShortResponse - это урезанная информация о пользователе для списков.
type userShortResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"` // Предполагаем, что у User есть поле Name
}


// handleSendMessage обрабатывает запрос на отправку нового сообщения.
// POST /api/v1/messages
func (h *Handler) handleSendMessage(w http.ResponseWriter, r *http.Request) {
	var req sendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.ReceiverID == 0 || req.Content == "" {
		h.respondWithError(w, http.StatusBadRequest, "ReceiverID and content are required")
		return
	}

	senderID, ok := r.Context().Value(userIDKey).(int64)
	if !ok {
		h.respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if senderID == req.ReceiverID {
		h.respondWithError(w, http.StatusBadRequest, "Cannot send message to yourself")
		return
	}

	msg := &domain.Message{
		SenderID:   senderID,
		ReceiverID: req.ReceiverID,
		Content:    req.Content,
	}

	createdMsg, err := h.messageUseCase.Send(r.Context(), msg)
	if err != nil {
		// Здесь может быть проверка на существование receiver_id
		h.respondWithError(w, http.StatusInternalServerError, "Failed to send message")
		return
	}

	h.respondWithJSON(w, http.StatusCreated, toMessageResponse(createdMsg))
}

// handleGetConversationHistory обрабатывает запрос на получение истории сообщений с одним пользователем.
// GET /api/v1/messages/history/{partner_id}
func (h *Handler) handleGetConversationHistory(w http.ResponseWriter, r *http.Request) {
	// Получаем ID собеседника из URL
	partnerIDStr := r.PathValue("partner_id")
	partnerID, err := strconv.ParseInt(partnerIDStr, 10, 64)
	if err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid partner ID")
		return
	}

	// Получаем ID текущего пользователя из контекста
	userID, ok := r.Context().Value(userIDKey).(int64)
	if !ok {
		h.respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	
	// Вызываем бизнес-логику для получения истории
	messages, err := h.messageUseCase.GetConversationHistory(r.Context(), userID, partnerID)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to get conversation history")
		return
	}
	
	// Конвертируем доменные модели в DTO для ответа
	response := make([]messageResponse, len(messages))
	for i, msg := range messages {
		response[i] = toMessageResponse(&msg)
	}

	h.respondWithJSON(w, http.StatusOK, response)
}


// handleGetConversations обрабатывает запрос на получение списка всех диалогов пользователя.
// GET /api/v1/messages/conversations
func (h *Handler) handleGetConversations(w http.ResponseWriter, r *http.Request) {
    userID, ok := r.Context().Value(userIDKey).(int64)
    if !ok {
        h.respondWithError(w, http.StatusUnauthorized, "Unauthorized")
        return
    }

    // UseCase должен вернуть специальную структуру, представляющую диалоги
    conversations, err := h.messageUseCase.GetConversations(r.Context(), userID)
    if err != nil {
        h.respondWithError(w, http.StatusInternalServerError, "Failed to get conversations")
        return
    }

    // Конвертируем результат в DTO для ответа.
    // Эта логика зависит от того, какую структуру вернет UseCase.
    // Примерная реализация:
    // response := make([]conversationResponse, len(conversations))
    // for i, conv := range conversations {
    //     response[i] = toConversationResponse(&conv)
    // }

    h.respondWithJSON(w, http.StatusOK, conversations) // Пока просто возвращаем то, что дал UseCase
}


// toMessageResponse - вспомогательная функция-конвертер из доменной модели в DTO.
func toMessageResponse(m *domain.Message) messageResponse {
	return messageResponse{
		ID:         m.ID,
		SenderID:   m.SenderID,
		ReceiverID: m.ReceiverID,
		Content:    m.Content,
		SentAt:     m.SentAt,
		ReadAt:     m.ReadAt,
	}
}
