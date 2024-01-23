package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rifansyah/go-crud/app/http/controllers"
)

type PostsRoute struct {
	ge *gin.RouterGroup
}

func NewPostsRoute(ge *gin.RouterGroup) *PostsRoute {
	return &PostsRoute{ge}
}

func (r *PostsRoute) SetRoutes() {
	postController := controllers.NewPostsController()

	r.ge.POST("/posts", postController.Create)
	r.ge.GET("/posts", postController.Index)
	r.ge.GET("/posts/:id", postController.GetPost)
	r.ge.PUT("/posts/:id", postController.Update)
	r.ge.DELETE("/posts/:id", postController.Detele)
}
