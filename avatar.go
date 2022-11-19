package main

import (
	"avatar.com/avatar/server"
	"avatar.com/avatar/server/conf"
	"log"
)

func main() {
	r := server.SetupServer()

	err := r.Run(":" + conf.ServiceAvatarPort)
	if err != nil {
		log.Fatalln("Avatar server crashed")
		return
	}
}
