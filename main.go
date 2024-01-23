package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rifansyah/go-crud/app/http/middlewares"
	"github.com/rifansyah/go-crud/app/initializers"
	"github.com/rifansyah/go-crud/app/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.Use(middlewares.RequireAuth())
	
	routes.InitRouter(r)

	r.Run()
}
