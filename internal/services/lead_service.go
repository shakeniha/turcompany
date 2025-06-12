package services

import (
	"turcompany/internal/models"
	"turcompany/internal/repositories"
)

type LeadService struct {
	Repo *repositories.LeadRepository
}

func NewLeadService(repo *repositories.LeadRepository) *LeadService {
	return &LeadService{Repo: repo}
}
func (s *LeadService) Create(lead *models.Leads) error {
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
