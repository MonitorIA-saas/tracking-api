package gingonic

import "github.com/gin-gonic/gin"

func StartAPI() {
	router := gin.Default()
	group := router.Group("monitoring/api")

	// adding middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// adding our routers groups
	handleTrack(group)

	router.Run()
}
