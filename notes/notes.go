package main

import (
	"bujem/notes/controller"
	"log"
)

func main() {
	log.Println("BuJEm Notes service")
	log.Println("Configuring routing")
	controller.Listen()
}
