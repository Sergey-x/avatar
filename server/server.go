package server

import (
	"avatar.com/avatar/server/conf"
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	r := gin.New()

	// Change default logger to custom
	r = setLogger(r)

	// Set paths to handler functions
	r = setRoutes(r)

	// Set a lower memory limit for multipart forms (override default 32 MB)
	r.MaxMultipartMemory = conf.MaxAvatarMiBSize << 20

	return r
}
