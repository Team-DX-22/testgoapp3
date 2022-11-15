package main

import (
	_ "github.com/diegoluisi/testgoapp3/config"
	"github.com/diegoluisi/testgoapp3/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "Hello world for your app testgoapp3",
		})
	})
	server.GET("/health-check/liveness", controllers.HealthCheckLiveness)
	server.GET("/health-check/readiness", controllers.HealthCheckReadiness)
	server.Run()
}
