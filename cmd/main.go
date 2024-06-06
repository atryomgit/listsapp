package main

import (
	todo "listsapp"
	"listsapp/pack/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
