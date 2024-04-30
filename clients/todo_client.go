package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/vnoitkumar/demyst-code-kata/configurations"
	"github.com/vnoitkumar/demyst-code-kata/models/responses"
)

//go:generate mockgen -source=./todo_client.go -destination=./mocks/todo_client_mock.go -package=mocks

type TodoClient interface {
	GetTodoItem(id int) (todoResponse *responses.TodoResponse, err error)
}

type todoClient struct {
	config *configurations.Config
}

func NewTodoClient(config *configurations.Config) TodoClient {
	return &todoClient{
		config: config,
	}
}

func (client *todoClient) GetTodoItem(id int) (todoResponse *responses.TodoResponse, err error) {
	log.Info().Msgf("Construct url to get todo item %d", id)

	url := fmt.Sprintf("%s%d", client.config.TodoURL, id)

	log.Info().Msgf("Getting todo item from %s", url)
	response, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Msgf("Error getting todo item on %s", url)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Error().Msgf("Error http status code is not 200 but got %d", response.StatusCode)
		return
	}

	log.Info().Msgf("Reading response body for item %d", id)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msgf("Error reading response body for item %d", id)
		return
	}

	log.Info().Msgf("Unmarshalling todoResponse for item %d", id)
	err = json.Unmarshal(body, &todoResponse)
	if err != nil {
		log.Error().Err(err).Msgf("Error unmarshalling todoResponse for item %d", id)
		return
	}

	log.Info().Msgf("Successfully got todo item from '%s'", url)
	return
}
