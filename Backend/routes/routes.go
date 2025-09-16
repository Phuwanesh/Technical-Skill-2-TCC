package routes

import (
	"authdemo/handlers"
	"authdemo/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(store handlers.UserStore) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		allowedOrigins := []string{
			"http://localhost:5173",
			"http://localhost:5174",
		}

		for _, o := range allowedOrigins {
			if origin == o {
				c.Header("Access-Control-Allow-Origin", origin)
				break
			}
		}
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	authHandler := handlers.NewAuthHandler(store)

	api := r.Group("/api/v1")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
		api.GET("/me", middleware.AuthRequired(), handlers.Me)
	}
	return r
}
