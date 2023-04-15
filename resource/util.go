package resource

// remove removes the element at index i from arr.
func remove[T any](arr []T, i int) []T {
	return arr[:i+copy(arr[i:], arr[i+1:])]
}
