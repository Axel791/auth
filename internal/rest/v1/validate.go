package v1

import (
	"encoding/json"

	"github.com/Axel791/auth/internal/common"
	userAPI "github.com/Axel791/auth/internal/rest/v1/api"
	authScenarios "github.com/Axel791/auth/internal/usecases/auth/scenarios"
	log "github.com/sirupsen/logrus"

	"errors"
	"net/http"
)

type ValidationHandler struct {
	logger            *log.Logger
	validationUseCase authScenarios.Validate
}

func NewValidationHandler(
	logger *log.Logger,
	validationUseCase authScenarios.Validate,
) *ValidationHandler {
	return &ValidationHandler{
		logger:            logger,
		validationUseCase: validationUseCase,
	}
}

func (h *ValidationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var input userAPI.Token
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Infof("err decode body: %v", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.validationUseCase.Execute(r.Context(), input.Token)
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
}
