package services

import (
	"fmt"
	"math/rand"
	"time"
	"turcompany/internal/models"
	"turcompany/internal/repositories"
)

type SMS_Service struct {
	Repo *repositories.SMSConfirmationRepository
}

// generateCode создает 6-значный код
func (s *SMS_Service) generateCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// SendSMS создает новую запись и отправляет код
func (s *SMS_Service) SendSMS(documentID int64, phone string) error {
	code := s.generateCode()
	sms := &models.SMSConfirmation{
		DocumentID:  documentID,
		Phone:       phone,
		SMSCode:     code,
		SentAt:      time.Now(),
		Confirmed:   false,
		ConfirmedAt: time.Time{},
	}
	_, err := s.Repo.Create(sms)
	if err != nil {
		return err
	}

	// test в хосте 
	fmt.Printf("📲 SMS sent to %s: code is %s\n", phone, code)
	return nil
}

// ResendSMS повторно отправляет код (если последний не подтвержден и не истёк)
func (s *SMS_Service) ResendSMS(documentID int64) error {
	existing, err := s.Repo.GetLatestByDocumentID(documentID)
	if err != nil {
		return err
	}
	if existing == nil || existing.Confirmed || s.IsCodeExpired(existing.SentAt) {
		return s.SendSMS(documentID, existing.Phone)
	}

	// Повторно отправляем тот же код
	fmt.Printf("🔁 Resending SMS to %s: code is %s\n", existing.Phone, existing.SMSCode)
	return nil
}

// ConfirmCode проверяет, совпадает ли код
func (s *SMS_Service) ConfirmCode(documentID int64, code string) (bool, error) {
	sms, err := s.Repo.GetByDocumentIDAndCode(documentID, code)
	if err != nil {
		return false, err
	}
	if sms == nil || sms.Confirmed {
		return false, nil
	}
	if s.IsCodeExpired(sms.SentAt) {
		return false, nil
	}

	sms.Confirmed = true
	sms.ConfirmedAt = time.Now()
	return true, s.Repo.Update(sms)
}

// IsCodeExpired проверяет, истёк ли срок действия кода (5 минут)
func (s *SMS_Service) IsCodeExpired(sentAt time.Time) bool {
	expiration := sentAt.Add(5 * time.Minute)
	return time.Now().After(expiration)
}

// DeleteConfirmation удаляет все подтверждения по документу
func (s *SMS_Service) DeleteConfirmation(documentID int64) error {
	return s.Repo.DeleteByDocumentID(documentID)
}

// GetLatestByDocumentID возвращает последнее подтверждение
func (s *SMS_Service) GetLatestByDocumentID(documentID int64) (*models.SMSConfirmation, error) {
	return s.Repo.GetLatestByDocumentID(documentID)
}
