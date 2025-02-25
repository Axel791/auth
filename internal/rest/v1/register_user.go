package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Axel791/auth/internal/common"

	userAPI "github.com/Axel791/auth/internal/rest/v1/api"
	"github.com/Axel791/auth/internal/usecases/auth/dto"
	authScenarios "github.com/Axel791/auth/internal/usecases/auth/scenarios"
	log "github.com/sirupsen/logrus"
)

type RegistrationHandler struct {
	logger              *log.Logger
	registrationUseCase authScenarios.Registration
}

func NewRegistrationHandler(
	registrationUseCase authScenarios.Registration,
	logger *log.Logger,
) *RegistrationHandler {
	return &RegistrationHandler{
		registrationUseCase: registrationUseCase,
		logger:              logger,
	}
}

func (h *RegistrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var input userAPI.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Infof("err decode body: %v", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	userDTO := dto.UserDTO{
		Login:    input.Login,
		Password: input.Password,
	}

	err := h.registrationUseCase.Execute(r.Context(), userDTO)
	if err != nil {
		var appErr *common.AppError

		h.logger.Infof("err login: %v", err)

		if ok := errors.As(err, &appErr); ok {
			http.Error(w, appErr.Message, appErr.Code)
		} else {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
}
