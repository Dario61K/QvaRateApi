package internal

import (
	h "qvarate_api/internal/handlers"
	"qvarate_api/internal/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	r := gin.Default()
	r.Use(middleware.SetCors())

	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) { // health check
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})


	api.GET("/get-currency/:startdate/:enddate", middleware.RateLimit(60, time.Minute*1, time.Minute*5), h.GetCurrency)
	api.GET("/get-excel/:startdate/:enddate", middleware.RateLimit(6, time.Minute*1, time.Minute*5), h.GetExcel)

	return r
}
