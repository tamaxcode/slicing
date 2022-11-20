package sliceutil

// FilterFunc represents filter callback function. It will receives index and slice's element of each iteration as parameter
type FilterFunc[T any] func (idx int, e T) bool

// Filter will filter slice based on FilterFunc provided
func Filter[T any](list []T, cb FilterFunc[T]) []T {
    filtered := make([]T, 0)

    for i, e := range list {
        if cb(i, e) {
            filtered = append(filtered, e)
        }
    }

    return filtered
}