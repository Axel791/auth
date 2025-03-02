package scenarios

import (
	"context"
	"github.com/Axel791/appkit"
	"net/http"

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
func NewRegistrationScenario(
	userRepository repositories.UserRepository,
	hashPasswordService services.HashPasswordService,
) *RegistrationScenario {
	return &RegistrationScenario{
		userRepository:      userRepository,
		hashPasswordService: hashPasswordService,
	}
}

// Execute - Функция выполнения сценария
func (s *RegistrationScenario) Execute(ctx context.Context, userDTO dto.UserDTO) error {
	userDomain := domains.User{
		Login:    userDTO.Login,
		Password: userDTO.Password,
	}

	if err := userDomain.ValidatePassword(); err != nil {
		return apikit.ValidationError(err.Error())
	}

	if err := userDomain.ValidateLogin(); err != nil {
		return apikit.ValidationError(err.Error())
	}

	user, err := s.userRepository.GetUserByLogin(ctx, userDomain.Login)
	if err != nil {
		return apikit.WrapError(
			http.StatusInternalServerError,
			"error getting user by login",
			err,
		)
	}

	if user.ID > 0 {
		return apikit.BadRequestError("user login already exists")
	}

	hashedPassword := s.hashPasswordService.Hash(userDomain.Password)
	userDomain.Password = hashedPassword

	err = s.userRepository.CreateUser(ctx, userDomain)
	if err != nil {
		return apikit.WrapError(http.StatusInternalServerError, "error creating user", err)
	}
	return nil
}
