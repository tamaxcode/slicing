package sliceutil

// MapFunc represents map callback function.
// It will receives index and slice's element of each iteration as parameter
// and return expected type value
type MapFunc[T, R any] func(int, T) R

// Map will map each element in slice to desired output
func Map[T, R any](src []T, cb MapFunc[T, R]) []R {
	l := make([]R, len(src))

	for i, e := range src {
		l[i] = cb(i, e)
	}

	return l
}
