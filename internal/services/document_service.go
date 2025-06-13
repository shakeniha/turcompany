package services

import (
	"fmt"
	"time"
	"turcompany/internal/models"
	"turcompany/internal/repositories"
)

type DocumentService struct {
	Repo    *repositories.DocumentRepository
	smsRepo *repositories.SMSConfirmationRepository
}

func NewDocumentService(repo *repositories.DocumentRepository) *DocumentService {
	return &DocumentService{Repo: repo}
}

func (s *DocumentService) CreateDocument(doc *models.Document) (int64, error) {
	return s.Repo.Create(doc)
}

func (s *DocumentService) GetDocument(id int64) (*models.Document, error) {
	return s.Repo.GetByID(id)
}

func (s *DocumentService) VerifyDocument(id int64) error {
	return s.Repo.UpdateStatus(id, "verified")
}

func (s *DocumentService) SendSMSConfirmation(docID int64, code string) error {
	confirmation := &models.SMSConfirmation{
		DocumentID: docID,
		SMSCode:    code,
		SentAt:     time.Now(),
		Confirmed:  false,
	}
	_, err := s.smsRepo.Create(confirmation)
	return err
}

func (s *DocumentService) ConfirmDocument(docID int64, code string) error {
	confirmation, err := s.smsRepo.GetByID(docID)
	if err != nil || confirmation == nil {
		return fmt.Errorf("confirmation not found")
	}

	if confirmation.SMSCode != code {
		return fmt.Errorf("invalid code")
	}

	confirmation.Confirmed = true
	confirmation.ConfirmedAt = time.Now()

	if err := s.smsRepo.Update(confirmation); err != nil {
		return err
	}

	return s.Repo.UpdateStatus(docID, "signed")
}

func (s *DocumentService) ListDocumentsByDeal(dealID int64) ([]*models.Document, error) {
	return s.Repo.ListDocumentsByDeal(dealID)
}

func (s *DocumentService) DeleteDocument(id int64) error {
	return s.Repo.Delete(id)
}
