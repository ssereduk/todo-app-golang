package service

import "github.com/ssereduk/todo-app-golang/pkg/reposi"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(*reposi.Repository) *Service {
	return &Service{}
}
