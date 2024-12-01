package fp

func Filter[T any](in []T, predicate func(it T) bool) []T {
	result := make([]T, 0, len(in))
	for _, item := range in {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func AllMatch[T any](in []T, predicate func(it T) bool) bool {
	for _, item := range in {
		if !predicate(item) {
			return false
		}
	}
	return true
}
