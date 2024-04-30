package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"github.com/vnoitkumar/demyst-code-kata/clients/mocks"
	"github.com/vnoitkumar/demyst-code-kata/models/responses"
	"github.com/vnoitkumar/demyst-code-kata/services"
)

type todoServiceTestSuite struct {
	suite.Suite
	context        context.Context
	mockController *gomock.Controller
	mockToDoClient *mocks.MockTodoClient
	service        services.TodoService
}

func TestTodoServiceTestSuite(t *testing.T) {
	suite.Run(t, new(todoServiceTestSuite))
}

func (suite *todoServiceTestSuite) SetupTest() {
	suite.context = context.TODO()
	suite.mockController = gomock.NewController(suite.T())
	suite.mockToDoClient = mocks.NewMockTodoClient(suite.mockController)

	suite.service = services.NewTodoService(suite.mockToDoClient)
}

func (suite *todoServiceTestSuite) TearDownTest() {
	suite.mockController.Finish()
}

func (suite *todoServiceTestSuite) TestGetTodoItem_ShouldReturnTodoResponse_WhenClientReturnsTodoResponseWithoutError() {
	expectedTodoResponse := &responses.TodoResponse{
		Id:        10,
		UserId:    1,
		Title:     "quis ut nam facilis et officia qui",
		Completed: true,
	}

	suite.mockToDoClient.EXPECT().GetTodoItem(10).Return(expectedTodoResponse, nil).Times(1)
	actualTodoResponse, err := suite.service.GetTodoItem(10)

	suite.Nil(err)
	suite.Equal(expectedTodoResponse, actualTodoResponse)
}

func (suite *todoServiceTestSuite) TestGetTodoItem_ShouldReturnError_WhenClientReturnsError() {
	expectedError := errors.New("some client error")
	suite.mockToDoClient.EXPECT().GetTodoItem(10).Return(nil, expectedError).Times(1)
	todoResponse, actualError := suite.service.GetTodoItem(10)

	suite.Nil(todoResponse)
	suite.Equal(expectedError, actualError)
}
