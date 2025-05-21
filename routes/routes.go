package routes

import (
	"go-backend/controllers"
	"go-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	auth := api.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	authorized := api.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		// Customer Management
		authorized.GET("/customers", controllers.GetCustomers)
		authorized.GET("/customers/:id", controllers.GetCustomer)
		authorized.POST("/customers", controllers.CreateCustomer)
		authorized.PUT("/customers/:id", controllers.UpdateCustomer)
		authorized.DELETE("/customers/:id", controllers.DeleteCustomer)

		// Order Management
		authorized.GET("/orders", controllers.GetOrders)
		authorized.GET("/orders/:id", controllers.GetOrder)
		authorized.POST("/orders", controllers.CreateOrder)
		authorized.PUT("/orders/:id", controllers.UpdateOrder)
		authorized.DELETE("/orders/:id", controllers.DeleteOrder)
	}
}
