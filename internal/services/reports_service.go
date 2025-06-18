package services

import (
	"errors"
	"turcompany/internal/models"
	"turcompany/internal/repositories"
)

type ReportService struct {
	LeadRepo *repositories.LeadRepository
	DealRepo *repositories.DealRepository
}

func NewReportService(leadRepo *repositories.LeadRepository, dealRepo *repositories.DealRepository) *ReportService {
	return &ReportService{
		LeadRepo: leadRepo,
		DealRepo: dealRepo,
	}
}

func (s *ReportService) GetSummary() (map[string]int, error) {
	totalLeads, err := s.LeadRepo.CountLeads()
	if err != nil {
		return nil, err
	}

	totalDeals, err := s.DealRepo.CountDeals()
	if err != nil {
		return nil, err
	}

	return map[string]int{
		"totalLeads": totalLeads,
		"totalDeals": totalDeals,
	}, nil
}
func (s *ReportService) FilterDeals(status, fromDate, toDate string) ([]models.Deals, error) {
	if fromDate == "" && toDate == "" && status == "" {
		return nil, errors.New("не указаны фильтры")
	}
	return s.DealRepo.FilterDeals(status, fromDate, toDate)
}
