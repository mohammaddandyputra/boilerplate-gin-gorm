package services

import (
	"learn-gin-gorm/models"
	"learn-gin-gorm/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
