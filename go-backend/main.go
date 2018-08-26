package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
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
var secret = []byte("SECRET")

var getTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["admin"] = true
	claims["name"] = "Ado Kukic"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString(secret)

	w.Write([]byte(tokenString))
})

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func main() {
	r := mux.NewRouter()

	// Routes

	r.Handle("/status", statusHandler).Methods("GET")
	r.Handle("/posts", getPostHandler).Methods("GET")
	r.Handle("/posts", jwtMiddleware.Handler(createPostHandler)).Methods("POST")
	r.Handle("/get-token", getTokenHandler).Methods("GET")

	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
}
