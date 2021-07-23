package main

import (
	"demo/hello-svr/boot"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := boot.Init(); err != nil {
		log.Fatalf("cmd init error: %v", err)
	}
}
