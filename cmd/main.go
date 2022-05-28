package main

import (
	"log"

	"github.com/ssereduk/todo-app-golang"

	"github.com/ssereduk/todo-app-golang/pkg/handler"
	"github.com/ssereduk/todo-app-golang/pkg/reposi"
	"github.com/ssereduk/todo-app-golang/pkg/service"
)

func main() {
	repos := reposi.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured whilen running http server: %s", err.Error())
	}
}
