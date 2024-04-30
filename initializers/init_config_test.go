package initializers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vnoitkumar/demyst-code-kata/configurations"
	"github.com/vnoitkumar/demyst-code-kata/initializers"
)

func TestLoadConfig_ShouldNotReturnError_WhenGivenCorrectFilePath_WithCorrectData(t *testing.T) {
	var config configurations.Config
	actualErr := initializers.LoadConfig("./test-data/test_config.json", &config)

	assert.Nil(t, actualErr)
}

func TestLoadConfig_ShouldReturnNoSuchFileOrDirectoryError_WhenFilePathIsInvalid(t *testing.T) {
	var config configurations.Config
	actualErr := initializers.LoadConfig("./invalid_path_test_config.json", &config)

	assert.Equal(t, "open ./invalid_path_test_config.json: no such file or directory", actualErr.Error())
}

func TestLoadConfig_ShouldReturnCannotUnmarshalError_WhenFileHasInvalidDataType(t *testing.T) {
	var config configurations.Config
	actualErr := initializers.LoadConfig("./test-data/invalid_data_type_test_config.json", &config)

	assert.Equal(t, "json: cannot unmarshal number into Go struct field Config.todo_url of type string", actualErr.Error())
}

func TestLoadConfig_ShouldReturnEOFError_WhenFileHasInvalidFileContent(t *testing.T) {
	var config configurations.Config
	actualErr := initializers.LoadConfig("./test-data/invalid_json_test_config.json", &config)

	assert.Equal(t, "EOF", actualErr.Error())
}

func TestValidateConfig_ShouldNotReturnError_WhenStuctIsValid(t *testing.T) {
	config := configurations.Config{
		TodoURL:       "https://jsonplaceholder.typicode.com/todos/",
		TodoListSize:  20,
		TodoChunkSize: 6,
	}
	actualErr := initializers.ValidateConfig(config)

	assert.Nil(t, actualErr)
}

func TestValidateConfig_ShouldReturnError_WhenTodoURLIsEmpty(t *testing.T) {
	config := configurations.Config{
		TodoURL:       "",
		TodoListSize:  20,
		TodoChunkSize: 6,
	}
	actualErr := initializers.ValidateConfig(config)

	assert.Equal(t, "Key: 'Config.TodoURL' Error:Field validation for 'TodoURL' failed on the 'required' tag", actualErr.Error())
}

func TestValidateConfig_ShouldReturnError_WhenTodoListSizeIsEmpty(t *testing.T) {
	config := configurations.Config{
		TodoURL:       "https://jsonplaceholder.typicode.com/todos/",
		TodoListSize:  0,
		TodoChunkSize: 6,
	}
	actualErr := initializers.ValidateConfig(config)

	assert.Equal(t, "Key: 'Config.TodoListSize' Error:Field validation for 'TodoListSize' failed on the 'required' tag", actualErr.Error())
}

func TestValidateConfig_ShouldReturnError_WhenTodoChunkSizeIsEmpty(t *testing.T) {
	config := configurations.Config{
		TodoURL:       "https://jsonplaceholder.typicode.com/todos/",
		TodoListSize:  20,
		TodoChunkSize: 0,
	}
	actualErr := initializers.ValidateConfig(config)

	assert.Equal(t, "Key: 'Config.TodoChunkSize' Error:Field validation for 'TodoChunkSize' failed on the 'required' tag", actualErr.Error())
}
