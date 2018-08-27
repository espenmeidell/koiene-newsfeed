package data

// Post struct
type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

// Posts stored in app
var posts = []Post{Post{ID: "5e7995cb-7422-4b2d-a5ed-8907a3bbe435", Title: "Første Koiemøte", Content: "Semesterets første koiemøte er klokken 19:30 på Idrettsbygget Gløshaugen. Møtet er åpent for alle som er interessert. Etter møtet blir det kake og koieprat =)", Status: "ACTIVE", Timestamp: 1535310038}}

// GetAllPosts returns all posts
func GetAllPosts() []Post {
	return posts
}

// InsertPost stores a post
func InsertPost(post Post) {
	posts = append(posts, post)
}
