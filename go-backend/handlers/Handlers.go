package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"../data"

	uuid "github.com/satori/go.uuid"
)

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
})

// Handlers for posts

var GetPostHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(data.GetAllPosts())
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

var CreatePostHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var p data.Post
	err := decoder.Decode(&p)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	p.ID = uuid.Must(uuid.NewV4()).String()
	p.Status = "ACTIVE"
	p.Timestamp = time.Now().Unix()
	data.InsertPost(p)

	payload, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})
