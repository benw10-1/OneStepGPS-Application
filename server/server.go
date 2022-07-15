package main

import (
	"log"
	"server/auth"
	. "server/config"
	"server/handlers"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	port := GetEnv("PORT", "3000")
	router := gin.Default()
	router.Use(auth.AuthMiddleWare())

	router.Use(static.Serve("/", static.LocalFile("../client/dist", true)))

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.File("../client/dist/favicon.ico")
	})
	router.POST("/api/login", handlers.Login)
	router.POST("/api/signup", handlers.SignUp)
	router.GET("/api/devices", handlers.GetDevices)
	router.POST("/api/setAPIKey", handlers.SetAPIKey)
	router.GET("/api/preferences", handlers.Preferences)
	router.POST("/api/preferences", handlers.Preferences)
	router.POST("/api/reverseGeocode", handlers.ReverseGeocode)

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(302, "/")
	})

	log.Fatal(router.Run(":" + port))
}
