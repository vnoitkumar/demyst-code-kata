package services

import (
	"errors"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/vnoitkumar/demyst-code-kata/configurations"
	"github.com/vnoitkumar/demyst-code-kata/models/responses"
	"github.com/vnoitkumar/demyst-code-kata/utils"
)

type TodoCheckerService interface {
	CheckTodoListStatus() (todoResponses []responses.TodoResponse, err error)
}

type todoCheckerService struct {
	config      *configurations.Config
	todoService TodoService
}

func NewTodoCheckerService(config *configurations.Config, todoService TodoService) TodoCheckerService {
	return &todoCheckerService{
		config:      config,
		todoService: todoService,
	}
}

func (service *todoCheckerService) CheckTodoListStatus() (todoResponses []responses.TodoResponse, err error) {
	evenNumberedSlice := utils.GetEvenNumberedSlice(service.config.TodoListSize)

	wg := sync.WaitGroup{}
	doneChan := make(chan interface{}, 1)
	errChan := make(chan error, len(evenNumberedSlice))
	semaphore := make(chan struct{}, service.config.TodoChunkSize)
	todoResponseChan := make(chan *responses.TodoResponse, len(evenNumberedSlice))
	errs := []error{}

	for _, evenNumber := range evenNumberedSlice {
		log.Info().Msgf("Acquire a semaphore slot to limit concurrency on iteration for %d", evenNumber)
		semaphore <- struct{}{}

		log.Info().Msgf("Increment wait group counter on iteration for %d", evenNumber)
		wg.Add(1)

		log.Info().Msgf("Start a goroutine for each even number on iteration for %d", evenNumber)
		go func(evenNumber int) {
			defer func() {
				log.Info().Msgf("Release the semaphore slot and decrement wait group counter after finishing, on iteration for %d", evenNumber)
				<-semaphore
				wg.Done()
			}()

			todoResponse, err := service.todoService.GetTodoItem(evenNumber)
			if err != nil {
				log.Error().Msgf("Send error to the error channel on iteration for %d", evenNumber)
				errChan <- err
				return
			}

			log.Info().Msgf("Send todo response to the response channel on iteration for %d", evenNumber)
			todoResponseChan <- todoResponse
		}(evenNumber)
	}

	go func() {
		log.Info().Msg("Wait for all goroutines to finish")
		wg.Wait()
		doneChan <- nil
	}()

	log.Info().Msg("Process results from all channels")
out:
	for {
		select {
		case err := <-errChan:
			log.Error().Err(err).Msgf("Error getting todo item")
			errs = append(errs, err)
		case todoResponse := <-todoResponseChan:
			todoResponses = append(todoResponses, *todoResponse)
			log.Debug().Msgf("todoResponse %#v", todoResponse)
			log.Trace().Msgf("Todo item of id %d with title '%s' is '%s'", todoResponse.Id, todoResponse.Title, utils.GetStatus(todoResponse.Completed))
		case <-doneChan:
			log.Info().Msg("Break the result process loop")
			break out
		}
	}

	if len(errs) > 0 {
		return todoResponses, errors.New("errors occurred during processing")
	}

	return
}
