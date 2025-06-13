package services

import (
	"errors"
	"strconv"
	"time"
	"turcompany/internal/models"
	"turcompany/internal/repositories"
)

type LeadService struct {
	Repo     *repositories.LeadRepository
	LeadRepo *repositories.LeadRepository
	DealRepo *repositories.DealRepository
}

func NewLeadService(leadRepo *repositories.LeadRepository, dealRepo *repositories.DealRepository) *LeadService {
	return &LeadService{
		Repo:     leadRepo,
		LeadRepo: leadRepo,
		DealRepo: dealRepo,
	}
}
func (s *LeadService) Create(lead *models.Leads) error {
	if s.Repo == nil {
		return errors.New("LeadRepository is not initialized")
	}
	if lead.Status == "" {
		lead.Status = "new"
	}
	return s.Repo.Create(lead)
}
func (s *LeadService) Update(lead *models.Leads) error {
	return s.Repo.Update(lead)
}
func (s *LeadService) GetByID(id int) (*models.Leads, error) {
	return s.Repo.GetByID(id)
}
func (s *LeadService) Delete(id int) error {
	return s.Repo.Delete(id)
}

func (s *LeadService) ConvertLeadToDeal(leadID int, amount, currency string) (*models.Deals, error) {
	lead, err := s.LeadRepo.GetByID(leadID)
	if err != nil {
		return nil, errors.New("lead not found")
	}

	if lead.Status != "confirmed" {
		return nil, errors.New("lead is not in a convertible status")
	}

	// Шаг 3: Создать новый объект сделки
	deal := &models.Deals{
		LeadID:    strconv.Itoa(lead.ID),
		Amount:    amount,
		Currency:  currency,
		Status:    "new",
		CreatedAt: time.Now(),
	}

	// Шаг 4: Сохранить сделку в базе данных через DealRepository
	if err := s.DealRepo.Create(deal); err != nil {
		return nil, err
	}

	// Шаг 5: Обновить статус лидов в базе данных
	lead.Status = "converted"
	if err := s.LeadRepo.Update(lead); err != nil {
		return nil, err
	}

	// Вернуть созданную сделку
	return deal, nil
}
