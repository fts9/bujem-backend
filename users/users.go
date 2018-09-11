package main

import (
	"bujem/users/controller"
	"log"
)

func main() {
	log.Println("BuJEm Users service")
	log.Println("Configuring routing")
	controller.Listen()
}
