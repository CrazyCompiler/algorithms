package sorts

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAscSort(t *testing.T) {
	listOfNumber := []int{503, -319, 245, -537, -167, -804, 469, -457, 143, 184, 376, 825, 11, -309, -144, 152, -493, 684, 165, 29}
	sortedList := AscSort(listOfNumber)
	expectedList := []int{-804, -537, -493, -457, -319, -309, -167, -144, 11, 29, 143, 152, 165, 184, 245, 376, 469, 503, 684, 825}
	assert.Equal(t, expectedList, sortedList)
}

func TestDescSort(t *testing.T) {
	listOfNumber := []int{503, -319, 245, -537, -167, -804, 469, -457, 143, 184, 376, 825, 11, -309, -144, 152, -493, 684, 165, 29}
	sortedList := DescSort(listOfNumber)
	expectedList := []int{825, 684, 503, 469, 376, 245, 184, 165, 152, 143, 29, 11, -144, -167, -309, -319, -457, -493, -537, -804}
	assert.Equal(t, expectedList, sortedList)
}
