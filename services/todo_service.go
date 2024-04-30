package services

import (
	"github.com/rs/zerolog/log"
	"github.com/vnoitkumar/demyst-code-kata/clients"
	"github.com/vnoitkumar/demyst-code-kata/models/responses"
)

//go:generate mockgen -source=./todo_service.go -destination=./mocks/todo_service_mock.go -package=mocks

type TodoService interface {
	GetTodoItem(id int) (todoResponse *responses.TodoResponse, err error)
}

type todoService struct {
	todoClient clients.TodoClient
}

func NewTodoService(todoClient clients.TodoClient) TodoService {
	return &todoService{
		todoClient: todoClient,
	}
}

func (service *todoService) GetTodoItem(id int) (todoResponse *responses.TodoResponse, err error) {
	log.Info().Msgf("Making API call to get todo item for %d", id)
	todoResponse, err = service.todoClient.GetTodoItem(id)
	if err != nil {
		log.Err(err).Msgf("Error from client for %d", id)
		return nil, err
	}

	return
}
