package requests

type PostsRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
