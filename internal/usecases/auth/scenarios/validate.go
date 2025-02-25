package scenarios

import (
	"context"
	"fmt"

	"github.com/Axel791/auth/internal/services"
	"github.com/Axel791/auth/internal/usecases/auth/repositories"
)

// ValidateScenario - структура сценария валидации
type ValidateScenario struct {
	userRepository repositories.UserRepository
	tokenService   services.TokenService
}

// NewValidateScenario - конструктор сценария валидации
func NewValidateScenario(
	userRepository repositories.UserRepository,
	tokenService services.TokenService,
) *ValidateScenario {
	return &ValidateScenario{
		userRepository: userRepository,
		tokenService:   tokenService,
	}
}

// Execute - сценарий валидации токена
func (s *ValidateScenario) Execute(ctx context.Context, token string) (bool, error) {
	userClaims, err := s.tokenService.ValidateToken(token)
	if err != nil {
		return false, fmt.Errorf("error validating token: %w", err)
	}

	user, err := s.userRepository.GetUserById(ctx, userClaims.UserID)
	if err != nil {
		return false, fmt.Errorf("error getting user by id: %w", err)
	}

	if user.ID == 0 {
		return false, fmt.Errorf("user not found")
	}
	return true, nil
}
