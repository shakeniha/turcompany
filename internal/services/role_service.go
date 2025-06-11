package services

import (
	"turcompany/internal/models"
	"turcompany/internal/repositories"
)

type RoleService interface {
	CreateRole(role *models.Role) error
	GetRoleByID(id int) (*models.Role, error)
	UpdateRole(role *models.Role) error
	DeleteRole(id int) error
	ListRoles() ([]*models.Role, error)
}

type roleService struct {
	repo repositories.RoleRepository
}

func NewRoleService(repo repositories.RoleRepository) RoleService {
	return &roleService{repo: repo}
}

func (s *roleService) CreateRole(role *models.Role) error {
	return s.repo.Create(role)
}

func (s *roleService) GetRoleByID(id int) (*models.Role, error) {
	return s.repo.GetByID(id)
}

func (s *roleService) UpdateRole(role *models.Role) error {
	return s.repo.Update(role)
}

func (s *roleService) DeleteRole(id int) error {
	return s.repo.Delete(id)
}

func (s *roleService) ListRoles() ([]*models.Role, error) {
	return s.repo.List()
}
