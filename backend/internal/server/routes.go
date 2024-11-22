package server

import (
	"backend/internal/modules/user"
	"database/sql"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HelloWorldHandler)

	mux.HandleFunc("/health", s.healthHandler)

	mux.HandleFunc("/api/signup", s.signupHandler)
	mux.HandleFunc("/api/login", s.loginHandler)

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

type ErrorOutput struct {
	Message string              `json:"message"`
	Details map[string][]string `json:"details"`
}

type LoginHandlerInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	input := LoginHandlerInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.validate.Struct(input)
	if err != nil {
		output := ErrorOutput{
			Message: "Validation error",
			Details: make(map[string][]string),
		}
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			field := strings.ToLower(err.Field())
			output.Details[field] = append(output.Details[field], err.Translate(s.trans))
		}

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(output)
		return
	}

	var user user.User
	err = s.db.QueryRow(r.Context(), "SELECT id, username, password_hash FROM users WHERE username = ? AND deleted_at IS NULL", input.Username).
		Scan(&user.Id, &user.Username, &user.PasswordHash)
	if err == sql.ErrNoRows {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		output := ErrorOutput{
			Message: "Invalid username or password",
		}
		json.NewEncoder(w).Encode(output)
		return
	} else if err != nil {
		slog.Info(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		output := ErrorOutput{
			Message: "Invalid username or password",
		}
		json.NewEncoder(w).Encode(output)
		return
	}

	token, err := generateToken(user.Id)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

type TokenClaims struct {
	UserId int64 `json:"userId"`
	jwt.RegisteredClaims
}

func generateToken(userId int64) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour)
	claims := &TokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("SECRET")

	return token.SignedString([]byte(secret))
}
