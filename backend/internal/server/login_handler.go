package server

import (
	"backend/internal/modules"
	"backend/internal/modules/user"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type LoginHandlerInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	input := LoginHandlerInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userRepo := user.NewUserRepo(s.db)
	authenticateUser := user.NewAuthenticateUser(userRepo, s.validate, s.trans)
	output, err := authenticateUser(r.Context(), user.AuthenticateUserInput{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		if errors.As(err, &modules.ValidationError{}) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err)
			return
		}

		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(output)
}
