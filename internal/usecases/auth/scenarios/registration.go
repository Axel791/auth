package scenarios

import (
	"context"
	"fmt"

	"github.com/Axel791/auth/internal/common"
	"github.com/Axel791/auth/internal/domains"
	"github.com/Axel791/auth/internal/services"
	"github.com/Axel791/auth/internal/usecases/auth/dto"
	"github.com/Axel791/auth/internal/usecases/auth/repositories"
)

// RegistrationScenario - структура сценария регистрация
type RegistrationScenario struct {
	userRepository      repositories.UserRepository
	hashPasswordService services.HashPasswordService
}

// NewRegistrationScenario - создание сценария
func NewRegistrationScenario(userRepository repositories.UserRepository) *RegistrationScenario {
	return &RegistrationScenario{userRepository: userRepository}
}

// Execute - Функция выполнения сценария
func (s *RegistrationScenario) Execute(ctx context.Context, userDTO dto.UserDTO) error {
	userDomain := domains.User{
		Login:    userDTO.Login,
		Password: userDTO.Password,
	}

	if err := userDomain.ValidatePassword(); err != nil {
		return common.NewValidationError(fmt.Sprintf("invalid password: %v", err))
	}

	if err := userDomain.ValidateLogin(); err != nil {
		return common.NewValidationError(fmt.Sprintf("invalid login: %v", err))
	}

	user, err := s.userRepository.GetUserByLogin(ctx, userDomain.Login)
	if err != nil {
		return common.NewInternalError(fmt.Sprintf("invalid fetch user: %v", err))
	}

	if user.ID > 0 {
		return common.NewBadRequestError("user login already exists")
	}

	hashedPassword := s.hashPasswordService.Hash(userDomain.Password)
	userDomain.Password = hashedPassword

	err = s.userRepository.CreateUser(ctx, userDomain)
	if err != nil {
		return fmt.Errorf("error create user: %w", err)
	}
	return nil
}
