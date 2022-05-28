package main

import (
	"log"

	"github.com/ssereduk/todo-app-golang"

	"github.com/ssereduk/todo-app-golang/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured whilen running http server: %s", err.Error())
	}
}
