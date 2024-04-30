package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vnoitkumar/demyst-code-kata/utils"
)

func TestGetStatus_Where_isCompleted_IsTrue_ShouldReturnCompleted(t *testing.T) {
	actualStatus := utils.GetStatus(true)

	assert.Equal(t, "Completed", actualStatus)
}

func TestGetStatus_Where_isCompleted_IsFalse_ShouldReturnNotCompleted(t *testing.T) {
	actualStatus := utils.GetStatus(false)

	assert.Equal(t, "Not Completed", actualStatus)
}
