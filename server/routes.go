package server

import (
	"avatar.com/avatar/server/handlers"
	"github.com/gin-gonic/gin"
)

func setRoutes(router *gin.Engine) *gin.Engine {
	v1 := router.Group("")
	{
		v1.GET("/avatar", handlers.GetAvatarPath)
		v1.POST("/avatar", handlers.SetAvatar)
	}

	return router
}
