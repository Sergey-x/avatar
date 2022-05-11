package conf

import (
	"log"
	"os"
)

var baseDir, _ = os.Getwd()
var BaseAvatarPath string = baseDir + "/avatars/"

// MaxAvatarMiBSize 4 MB
const MaxAvatarMiBSize = 4

var ServiceAvatarPort = os.Getenv("SERVICE_AVATAR_PORT")
var AllowedHosts = []string{"0.0.0.0"}

func init() {
	if ServiceAvatarPort == "" {
		log.Fatal("environment variable `SERVICE_AVATAR_PORT` must be specified and must be non empty value")
	}
}
