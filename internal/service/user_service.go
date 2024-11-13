package service

import (
	"go-api-gin/internal/domain"
	"go-api-gin/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	repo *repository.UserRepository
}

// Constructor
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Lógica de servicio para crear un nuevo usuario
func (s *UserService) CreateUser(user *domain.User) (*domain.User, error) {
	user.ID = primitive.NewObjectID() // Genera un nuevo ID
	_, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Lógica de servicio para obtener un usuario por su email
func (s *UserService) GetUserByEmail(email string) (*domain.User, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Lógica de servicio para obtener un usuario por su ID
func (s *UserService) GetUserByID(id string) (*domain.User, error) {
	user, err := s.repo.FindUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
