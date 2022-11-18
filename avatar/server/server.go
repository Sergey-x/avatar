package server

import (
	"github.com/gin-gonic/gin"
)

// MaxAvatarMiBSize 4 MB
const MaxAvatarMiBSize = 4

func SetupServer() *gin.Engine {
	r := gin.New()

	// Change default logger to custom
	r = setLogger(r)

	// Set paths to handler functions
	r = setRoutes(r)

	// Set a lower memory limit for multipart forms (override default 32 MB)
	r.MaxMultipartMemory = MaxAvatarMiBSize << 20

	return r
}
