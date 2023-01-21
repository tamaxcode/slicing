package sliceutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tamaxcode/sliceutil"
)

func TestForEach_IterateAll(t *testing.T) {
	numbers := make([]int64, 10)
	for i := 1; i <= 10; i++ {
		numbers[i-1] = int64(i)
	}

	got := make([]int64, 10)
	sliceutil.ForEach(numbers, func(i int, e int64) bool {
		got[i] = e
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
	sliceutil.ForEach(numbers, func(i int, e int64) bool {
		got = append(got, e)

		return !(i == 4)
	})

	want := []int64{1, 2, 3, 4, 5}

	assert.Equal(t, want, got)
}
