package sliceutil

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestFilter_Numbers(t *testing.T) {
    numbers := make([]int64, 10)
    for i := 1; i <= 10; i++ {
        numbers[i-1] = int64(i)
    }

    got := Filter(numbers, func(idx int, e int64) bool {
        return e % 2 == 0
    })

    want := []int64{2, 4, 6, 8, 10}

    assert.Equal(t, got, want)
}