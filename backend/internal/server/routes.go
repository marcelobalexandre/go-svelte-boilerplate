package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HelloWorldHandler)

	mux.HandleFunc("/health", ensureAuthenticated(s.healthHandler))

	mux.HandleFunc("/ws", ensureAuthenticated(s.wsHandler))
	mux.HandleFunc("/sse", s.sseHandler)

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

	fmt.Println("User ID", r.Header.Get("x-user-id"))

	_, _ = w.Write(jsonResp)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var clients = make(map[int]*websocket.Conn)

func (s *Server) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	userId, _ := strconv.Atoi(r.Header.Get("x-user-id"))
	clients[userId] = conn

	defer delete(clients, userId)
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Message", string(p))

		for _, clientConn := range clients {
			clientConn.WriteMessage(messageType, p)
		}
	}
}

func (s *Server) sseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	w.Header().Set("Access-Control-Allow-Origin", "*")

	clientDisconnected := r.Context().Done()

	rc := http.NewResponseController(w)
	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
		case <-clientDisconnected:
			fmt.Println("Client disconnected")
			return
		case <-t.C:
			_, err := fmt.Fprintf(w, "data: Time is %s\n\n", time.Now().Format(time.UnixDate))
			if err != nil {
				return
			}
			if err = rc.Flush(); err != nil {
				return
			}
		}
	}
}

type ErrorOutput struct {
	Message string              `json:"message"`
	Details map[string][]string `json:"details"`
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
	secret := os.Getenv("PASSWORD_SECRET")

	return token.SignedString([]byte(secret))
}

func ensureAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// authHeader := r.Header.Get("Authorization")
		// bearerToken := strings.Split(authHeader, " ")
		// if len(bearerToken) != 2 {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }

		token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(os.Getenv("PASSWORD_SECRET")), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r.Header.Set("x-user-id", strconv.Itoa(int(claims["userId"].(float64))))

		next(w, r)
	}
}
