package sliceutil_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tamaxcode/sliceutil"
)

func TestMap_NumbersTimes2(t *testing.T) {
	numbers := make([]int64, 10)

	for i := 1; i <= 10; i++ {
		numbers[i-1] = int64(i)
	}

	got := sliceutil.Map(numbers, func(i int, e int64) int64 {
		return e * 2
	})

	want := []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	assert.Equal(t, want, got)
}

func TestMap_NumbersToString(t *testing.T) {
	numbers := make([]int64, 10)

	for i := 1; i <= 10; i++ {
		numbers[i-1] = int64(i)
	}

	got := sliceutil.Map(numbers, func(i int, e int64) string {
		return strconv.FormatInt(e, 10)
	})

	want := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	assert.Equal(t, want, got)
}
