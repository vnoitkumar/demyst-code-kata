package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"github.com/vnoitkumar/demyst-code-kata/configurations"
	"github.com/vnoitkumar/demyst-code-kata/models/responses"
	"github.com/vnoitkumar/demyst-code-kata/services"
	"github.com/vnoitkumar/demyst-code-kata/services/mocks"
	"go.uber.org/goleak"
)

type todoCheckerServiceTestSuite struct {
	suite.Suite
	context         context.Context
	mockController  *gomock.Controller
	mockTodoService *mocks.MockTodoService
	config          *configurations.Config
	service         services.TodoCheckerService
}

func TestTodoCheckerServiceTestSuite(t *testing.T) {
	suite.Run(t, new(todoCheckerServiceTestSuite))
}

func (suite *todoCheckerServiceTestSuite) SetupTest() {
	suite.context = context.TODO()
	suite.mockController = gomock.NewController(suite.T())
	suite.mockTodoService = mocks.NewMockTodoService(suite.mockController)
	suite.config = &configurations.Config{
		TodoListSize:  1,
		TodoChunkSize: 1,
	}

	suite.service = services.NewTodoCheckerService(suite.config, suite.mockTodoService)
}

func (suite *todoCheckerServiceTestSuite) TearDownTest() {
	suite.mockController.Finish()
}

func (suite *todoCheckerServiceTestSuite) TestCheckTodoListStatus_Success() {
	expectedTodoResponses := []responses.TodoResponse{{Id: 2, UserId: 2, Title: "delectus aut autem", Completed: false}}
	expectedResponse := responses.TodoResponse{Id: 2, UserId: 2, Title: "delectus aut autem", Completed: false}
	suite.mockTodoService.EXPECT().GetTodoItem(2).Return(&expectedResponse, nil).Times(1)

	actualTodoResponses, err := suite.service.CheckTodoListStatus()
	suite.Nil(err)
	suite.Len(actualTodoResponses, 1)
	suite.Equal(expectedTodoResponses, actualTodoResponses)
}

func (suite *todoCheckerServiceTestSuite) TestCheckTodoListStatus_Failuer() {
	suite.mockTodoService.EXPECT().GetTodoItem(2).Return(nil, errors.New("some error")).Times(1)
	actualTodoResponses, err := suite.service.CheckTodoListStatus()
	suite.NotNil(err)
	suite.Equal("errors occurred during processing", err.Error())
	suite.Len(actualTodoResponses, 0)
}

func (suite *todoCheckerServiceTestSuite) TestCheckTodoListStatus_CheckMemoryLeaks() {
	defer goleak.VerifyNone(suite.T())

	config := &configurations.Config{
		TodoListSize:  5,
		TodoChunkSize: 2,
	}

	service := services.NewTodoCheckerService(config, suite.mockTodoService)

	for _, v := range []int{2, 4, 6, 8, 10} {
		expectedResponse := responses.TodoResponse{Id: v, UserId: v, Title: "delectus aut autem", Completed: false}
		suite.mockTodoService.EXPECT().GetTodoItem(v).Return(&expectedResponse, nil).Times(1)
	}

	_, err := service.CheckTodoListStatus()
	suite.Nil(err)
}
