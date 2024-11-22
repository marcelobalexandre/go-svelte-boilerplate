package server

import (
	"backend/internal/modules"
	"backend/internal/modules/user"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type SignupHandlerInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) signupHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Signup handler called")
	input := SignupHandlerInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userRepo := user.NewUserRepo(s.db)
	createUser := user.NewCreateUser(userRepo, s.validate, s.trans)
	user, err := createUser(r.Context(), user.CreateUserInput{
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
