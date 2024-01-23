package main

import (
	"github.com/rifansyah/go-crud/app/initializers"
	"github.com/rifansyah/go-crud/app/models"
)


func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.User{})
}
