package fp

func Map[T, U any](in []T, mapper func(t T) U) []U {
	result := make([]U, len(in))
	for i, v := range in {
		result[i] = mapper(v)
	}
	return result
}

func MapParallel[T, U any](in []T, mapper func(t T) U) []U {
	result := make([]U, len(in))

	type mapped struct {
		i int
		v U
	}
	resultChannel := make(chan mapped)
	for i, v := range in {
		go func(i int, v T) {
			resultChannel <- mapped{i, mapper(v)}
		}(i, v)
	}
	for i := 0; i < len(in); i++ {
		r := <-resultChannel
		result[r.i] = r.v
	}
	return result
}
