package peak_finding

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindAPeak(t *testing.T) {
	input1 := []int{1, 2, 3, 4, 5, 4, 3, 2, 6, 3}
	peakIndex1, number1 := FindAPeak(input1)

	input2 := []int{1}
	peakIndex2, number2 := FindAPeak(input2)

	assert.Equal(t, 4, peakIndex1)
	assert.Equal(t, 5, number1)
	assert.Equal(t, 0, peakIndex2)
	assert.Equal(t, 1, number2)
}

func TestFindAllPeaks(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 4, 3, 2, 6, 3,8}
	allPeaks := FindAllPeaks(input)
	expectedAllEquals := []peak{
		{4,5},
		{8,6},
		{10,8},
	}
	assert.Equal(t, expectedAllEquals, allPeaks)
}
