package main

import (
	todo "listsapp"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatal("error occured")
	}
}
