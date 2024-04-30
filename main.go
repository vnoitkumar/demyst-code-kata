package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vnoitkumar/demyst-code-kata/clients"
	"github.com/vnoitkumar/demyst-code-kata/configurations"
	"github.com/vnoitkumar/demyst-code-kata/constants"
	"github.com/vnoitkumar/demyst-code-kata/initializers"
	"github.com/vnoitkumar/demyst-code-kata/services"
)

var initStart time.Time
var mainStart time.Time
var config configurations.Config

func init() {
	initStart = time.Now()
	defer func() {
		log.Debug().Msgf("Init execution took, %v", time.Since(initStart))
	}()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampNano})
	log.Logger = log.With().Caller().Logger()
	zerolog.TimeFieldFormat = time.StampNano

	err := initializers.LoadConfig(constants.CONFIG_FILE_PATH, &config)
	if err != nil {
		log.Err(err).Msg("Error loading config file, using default values")
		config = configurations.Config{
			TodoURL:       "https://jsonplaceholder.typicode.com/todos/",
			TodoListSize:  20,
			TodoChunkSize: 6,
		}
		return
	}

	err = initializers.ValidateConfig(config)
	if err != nil {
		log.Fatal().Err(err).Msg("Error validating config file")
		return
	}
}

func main() {
	mainStart = time.Now()
	defer func() {
		log.Debug().Msgf("Main execution took %v", time.Since(mainStart))
	}()

	todoClient := clients.NewTodoClient(&config)
	todoService := services.NewTodoService(todoClient)
	todoCheckerService := services.NewTodoCheckerService(&config, todoService)

	_, err := todoCheckerService.CheckTodoListStatus()
	if err != nil {
		log.Err(err).Msg("Error checking todo list status")
	}
}
