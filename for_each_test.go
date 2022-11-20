package sliceutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForEach_IterateAll(t *testing.T) {
	numbers := make([]int64, 10)
	for i := 1; i <= 10; i++ {
		numbers[i-1] = int64(i)
	}

	got := make([]int64, 0)
	ForEach(numbers, func(i int, e int64) bool {
		got = append(got, e)
		return true
	})

	assert.Equal(t, numbers, got)
}

func TestForEach_Break(t *testing.T) {
	numbers := make([]int64, 10)
	for i := 1; i <= 10; i++ {
		numbers[i-1] = int64(i)
	}

	got := make([]int64, 0)
	ForEach(numbers, func(i int, e int64) bool {
		got = append(got, e)

		if i == 4 {
			return false
		}
		return true
	})

	want := []int64{1, 2, 3, 4, 5}

	assert.Equal(t, want, got)
}
