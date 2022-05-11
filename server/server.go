package server

import (
	"avatar.com/avatar/server/conf"
	"github.com/gin-gonic/gin"
	"log"
)

func SetupServer() *gin.Engine {
	r := gin.New()

	err := r.SetTrustedProxies(conf.AllowedHosts)
	if err != nil {
		log.Fatalln("SetTrustedProxies error")
	}

	// Change default logger to custom
	r = setLogger(r)

	// Set paths to handler functions
	r = setRoutes(r)

	// Set a lower memory limit for multipart forms (default is 32 MB)
	r.MaxMultipartMemory = conf.MaxAvatarMiBSize << 20

	return r
}
