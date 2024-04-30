package clients_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"github.com/vnoitkumar/demyst-code-kata/clients"
	"github.com/vnoitkumar/demyst-code-kata/configurations"
)

type todoClientTestSuite struct {
	suite.Suite
	context        context.Context
	mockController *gomock.Controller
	mockServer     *httptest.Server
	config         *configurations.Config
	client         clients.TodoClient
}

// As mentioned in the FAQ section ["Do I need to write tests for connecting to API ? - That can be ommitted."]
// Ignoring the test (coverage) for this file, written one success case
func TestTodoClientTestSuite(t *testing.T) {
	suite.Run(t, new(todoClientTestSuite))
}

func (suite *todoClientTestSuite) SetupTest() {
	suite.context = context.TODO()
	suite.mockController = gomock.NewController(suite.T())
	suite.mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/todo/1":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"id": 1, "title": "Sample Todo", "completed": false}`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	suite.config = &configurations.Config{
		TodoURL: suite.mockServer.URL + "/todo/",
	}

	suite.client = clients.NewTodoClient(suite.config)
}

func (suite *todoClientTestSuite) TearDownTest() {
	suite.mockServer.Close()
	suite.mockController.Finish()
}

func (suite *todoClientTestSuite) TestGetTodoItem() {
	todoResponse, err := suite.client.GetTodoItem(1)

	suite.NoError(err)
	suite.NotNil(todoResponse)
	suite.Equal(1, todoResponse.Id)
	suite.Equal("Sample Todo", todoResponse.Title)
	suite.False(todoResponse.Completed)
}
