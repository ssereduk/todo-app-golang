package main

import (
	"log"
)


func main() {
	srv:= new(todo.Server)
	if err := srv.Run("8000"); err != nill {
		log.Fatalf("error occured whilen running http server: %s, err.Error()")
	}
}