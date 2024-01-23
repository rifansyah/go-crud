package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rifansyah/go-crud/app/http/controllers"
)

type UsersRoute struct {
	ge *gin.RouterGroup
}

func NewUsersRoute(ge *gin.RouterGroup) *UsersRoute {
	return &UsersRoute{ge}
}

func (r *UsersRoute) SetRoutes() {
	usersController := controllers.NewUsersController()

	r.ge.POST("/signup", usersController.SignUp)
	r.ge.POST("/login", usersController.Login)
}
