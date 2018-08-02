package sorts

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	listOfNumber := []int{503, -319, 245, -537, -167, -804, 469, -457, 143, 184, 376, 825, 11, -309, -144, 152, -493, 684, 165, 29}
	sortedList := Sort(listOfNumber)
	expectedList := []int{-804, -537, -493, -457, -319, -309, -167, -144, 11, 29, 143, 152, 165, 184, 245, 376, 469, 503, 684, 825}
	assert.Equal(t, expectedList, sortedList)
}
