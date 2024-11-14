package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

type User struct {
	Id        int
	Username  string
	Password  string
	DeletedAt *time.Time
}

type SignupHandlerInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) signupHandler(w http.ResponseWriter, r *http.Request) {
	input := SignupHandlerInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: Validate input.

	now := time.Now().UTC()
	_, err = s.db.Exec(
		r.Context(),
		"INSERT INTO users (username, password_hash, created_at, updated_at) VALUES (?, ?, ?, ?)",
		input.Username,
		// TODO: Hash the password.
		input.Password,
		now,
		now,
	)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
	// TODO: Return the user.
}

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

	var user User
	err = s.db.QueryRow(r.Context(), "SELECT id, username, password_hash FROM users WHERE username = ?", input.Username).
		Scan(&user.Id, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		slog.Info(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: Check if user is deleted.
	if input.Password != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid username or password"})
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
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

func generateToken(userId int) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour)
	claims := &TokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	// TODO: Move to env variable.
	secret := "super-secret-string"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
