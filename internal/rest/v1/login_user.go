package v1

import (
	"encoding/json"
	"errors"

	"github.com/Axel791/auth/internal/common"
	userAPI "github.com/Axel791/auth/internal/rest/v1/api"
	"github.com/Axel791/auth/internal/usecases/auth/dto"

	authScenarios "github.com/Axel791/auth/internal/usecases/auth/scenarios"
	log "github.com/sirupsen/logrus"

	"net/http"
)

type LoginHandler struct {
	logger       *log.Logger
	loginUseCase authScenarios.Login
}

func NewLoginHandler(
	logger *log.Logger,
	loginUseCase authScenarios.Login,
) *LoginHandler {
	return &LoginHandler{
		logger:       logger,
		loginUseCase: loginUseCase,
	}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var input userAPI.UserLogin

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.logger.Infof("err decode body: %v", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	loginDTO := dto.UserDTO{
		Login:    input.Login,
		Password: input.Password,
	}

	token, err := h.loginUseCase.Execute(r.Context(), loginDTO)
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
	common.ResponseJSON(w, userAPI.Token{Token: token.Token})
}
