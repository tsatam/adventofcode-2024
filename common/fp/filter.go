package fp

type Predicate[T any] func(int T) bool

func Filter[T any](in []T, predicate Predicate[T]) []T {
	result := make([]T, 0, len(in))
	for _, item := range in {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func AllMatch[T any](in []T, predicate Predicate[T]) bool {
	for _, item := range in {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func Not[T any](predicate Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return !predicate(t)
	}
}
