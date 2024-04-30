package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vnoitkumar/demyst-code-kata/utils"
)

func TestGetEvenNumberedSlice_Where_sliceSize_IsTen_ShouldReturnEvenNumberSliceWithLengthTen(t *testing.T) {
	evenNumberSlice := utils.GetEvenNumberedSlice(10)
	actualSliceSize := len(evenNumberSlice)

	assert.Equal(t, 10, actualSliceSize)
}

func TestGetEvenNumberedSlice_Where_sliceSize_IsTen_ShouldReturnEvenNumberSliceWithAllEvenNumbers(t *testing.T) {
	actualSlice := utils.GetEvenNumberedSlice(10)
	expectedSlice := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	assert.Equal(t, expectedSlice, actualSlice)
}

func TestGetEvenNumberedSlice_Where_sliceSize_IsZero_ShouldReturnEmptySliceWithLengthZero(t *testing.T) {
	evenNumberSlice := utils.GetEvenNumberedSlice(0)
	actualSliceSize := len(evenNumberSlice)

	assert.Equal(t, 0, actualSliceSize)
}
