package main

import (
	"log"

	"webserver-go/bot"
	"webserver-go/config"
	"webserver-go/webserver"
)

func main() {
	log.Println("Init config...")
	config.Init()

	log.Println("Init bot...")
	bot.Init()

	log.Println("Init webserver...")
	webserver.Init()
}
