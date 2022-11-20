package sliceutil

// ForEachFunc represents ForEach callback function.
// It will receives index and slice's element of each iteration as parameter
// and returns bool.
// Will stop iteration if return false
type ForEachFunc[T any] func(int, T) bool

// ForEach will iterate all elements in slice unless callback return false
func ForEach[T any](src []T, cb ForEachFunc[T]) {
	for i, e := range src {
		if !cb(i, e) {
			break
		}
	}
}
