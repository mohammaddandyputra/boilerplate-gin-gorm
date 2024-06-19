package services

import (
	"learn-gin-gorm/dto"
	"learn-gin-gorm/models"
	"learn-gin-gorm/repositories"
	"learn-gin-gorm/utils"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepo}
}

func (s *AuthService) ProfileUser(email string) (*models.User, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) RegisterUser(requestBody dto.RegisterDTO) error {
	hashedPassword, err := utils.HashPassword(requestBody.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     requestBody.Name,
		Email:    requestBody.Email,
		Password: (hashedPassword),
	}

	return s.userRepository.CreateUser(user)
}

func (s *AuthService) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = utils.CheckPasswordHash(user.Password, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
