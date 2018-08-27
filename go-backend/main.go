package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	h "./handlers"
	"github.com/auth0-community/auth0"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	jose "gopkg.in/square/go-jose.v2"
)

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

	r.Handle("/status", h.StatusHandler).Methods("GET")
	r.Handle("/posts", h.GetPostHandler).Methods("GET")
	r.Handle("/posts", authMiddleware(h.CreatePostHandler)).Methods("POST")

	cr := cors.AllowAll().Handler(r)

	log.Println("Server Started")
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, cr))
}
