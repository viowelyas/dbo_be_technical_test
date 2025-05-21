package main

import (
	"go-backend/config"
	"go-backend/models"
	"go-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Customer{}, &models.User{}, &models.Order{})
	routes.SetupRoutes(r)
	r.Run(":8080")
}
