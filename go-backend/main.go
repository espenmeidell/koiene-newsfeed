package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/auth0-community/auth0"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	jose "gopkg.in/square/go-jose.v2"
)

type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

var posts = []Post{Post{ID: "5e7995cb-7422-4b2d-a5ed-8907a3bbe435", Title: "Første Koiemøte", Content: "Semesterets første koiemøte er klokken 19:30 på Idrettsbygget Gløshaugen. Møtet er åpent for alle som er interessert. Etter møtet blir det kake og koieprat =)", Status: "ACTIVE", Timestamp: 1535310038}}

var notImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

var statusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
})

// Handlers for posts

var getPostHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(posts)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

var createPostHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var p Post
	err := decoder.Decode(&p)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	p.ID = uuid.Must(uuid.NewV4()).String()
	p.Status = "ACTIVE"
	p.Timestamp = time.Now().Unix()
	posts = append(posts, p)

	payload, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

// AUTH

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := []byte(os.Getenv("AUTH0_API_CLIENT_SECRET"))
		secretProvider := auth0.NewKeyProvider(secret)
		audience := []string{os.Getenv("AUTH0_API_AUDIENCE")}

		configuration := auth0.NewConfiguration(secretProvider, audience, os.Getenv("AUTH0_DOMAIN"), jose.HS256)
		validator := auth0.NewValidator(configuration, nil)

		token, err := validator.ValidateRequest(r)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Token is not valid:", token)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func main() {
	r := mux.NewRouter()

	// Routes

	r.Handle("/status", statusHandler).Methods("GET")
	r.Handle("/posts", getPostHandler).Methods("GET")
	r.Handle("/posts", authMiddleware(createPostHandler)).Methods("POST")

	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
}
