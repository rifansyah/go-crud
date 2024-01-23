package routes

import (
	"github.com/gin-gonic/gin"
	routes "github.com/rifansyah/go-crud/app/routes/api/v1"
)

func InitRouter(ge *gin.Engine) {
	v1 := ge.Group("/api/v1")
	{
		postsRouter := routes.NewPostsRoute(v1)
		postsRouter.SetRoutes()

		usersRouter := routes.NewUsersRoute(v1)
		usersRouter.SetRoutes()
	}
}
