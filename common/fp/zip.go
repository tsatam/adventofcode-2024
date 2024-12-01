package fp

func Zip[T any](in1, in2 []T) [][2]T {
	return ZipFunc(in1, in2, zipCombineTuple)
}

func ZipFunc[T, U any](in1, in2 []T, combine func(elem1, elem2 T) U) []U {
	size := min(len(in1), len(in2))

	result := make([]U, size)

	for i := range size {
		result[i] = combine(in1[i], in2[i])
	}

	return result
}

func zipCombineTuple[T any](elem1, elem2 T) [2]T {
	return [2]T{elem1, elem2}
}
