package services

import (
	"turcompany/internal/models"
	"turcompany/internal/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
	ListUsers() ([]*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type userService struct {
	repo         repositories.UserRepository
	emailService EmailService
}

func NewUserService(repo repositories.UserRepository, emailService EmailService) UserService {
	return &userService{
		repo:         repo,
		emailService: emailService,
	}
}

func (s *userService) CreateUser(user *models.User) error {
	if err := s.repo.Create(user); err != nil {
		return err
	}

	if err := s.emailService.SendWelcomeEmail(user.Email, user.CompanyName); err != nil {
		// might add proper logging
		return nil
	}

	return nil
}

func (s *userService) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}

func (s *userService) ListUsers() ([]*models.User, error) {
	return s.repo.List()
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetByEmail(email)
}
