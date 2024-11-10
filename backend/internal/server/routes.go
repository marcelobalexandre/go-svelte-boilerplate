package server

import (
	"encoding/json"
	"log"
	"net/http"
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

func (s *Server) signupHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]map[string]string)
	errors := make(map[string]string)
	errors["username"] = "Username is invalid"
	errors["email"] = "Email is invalid"
	resp["errors"] = errors

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.WriteHeader(http.StatusUnprocessableEntity)
	_, _ = w.Write(jsonResp)
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["errorMessage"] = "Invalid username or password"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write(jsonResp)
}
