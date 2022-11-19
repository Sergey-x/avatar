package conf

import (
	"log"
	"os"
)

var BaseAvatarPath string = os.Getenv("AVATAR_IMAGES_DIR")
var ScheduleService string = os.Getenv("SCHEDULE_SERVICE")

const UserIdHeader = "X-User-Identity"

// MaxAvatarMiBSize 4 MB
const MaxAvatarMiBSize = 4

var ServiceAvatarPort = os.Getenv("SERVICE_AVATAR_PORT")

func init() {
	if ServiceAvatarPort == "" {
		log.Fatal("environment variable `SERVICE_AVATAR_PORT` must be specified and must be non empty value")
	}
	if BaseAvatarPath == "" {
		log.Fatal("environment variable `AVATAR_IMAGES_DIR` must be specified and must be non empty value")
	}
	if ScheduleService == "" {
		log.Fatal("environment variable `SCHEDULE_SERVICE` must be specified and must be non empty value")
	}
}
