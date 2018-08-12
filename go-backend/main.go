package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Post entity
type Post struct {
	ID        string `json:"id" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Status    string `json:"status" binding:"required"`
	Timestamp int    `json:"timestamp" binding:"required"`
}

var posts = []Post{
	Post{"6d503941-8d24-437c-9f19-88c0c6dce48b", "Complaints at Fosenkoia", "We have received some complaints from neighbours at Fosen. Don't walk through private gardens! When walking from the speed boat you need to go up to the main road behind the store. The road along the docks is a dead end and will lead you into private gardens.", "ACTIVE", 1534108248199},
}

func getPosts(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, posts)
}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Setup route group for the API
	api := router.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	api.GET("/posts", getPosts)

	// Start and run the server
	router.Run(":3000")
}
