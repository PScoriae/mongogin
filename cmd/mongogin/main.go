package main

import (
	"mongogin/internal/app/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	users := r.Group("/users")

	{
		users.GET("", handlers.GetAllUsers)

		users.POST("", handlers.CreateUser)

		users.GET("/:id", handlers.GetUsersById)
	}

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
