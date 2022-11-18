package main

import (
	"avatar.com/avatar/avatar/config"
	"avatar.com/avatar/avatar/server"
	"fmt"
	"log"
)

func main() {
	r := server.SetupServer()

	err := r.Run(fmt.Sprintf(":%s", config.ServiceAvatarPort))
	if err != nil {
		log.Fatalln("Avatar server crashed")
		return
	}
}
