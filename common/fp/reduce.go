package fp

import "golang.org/x/exp/constraints"

func Reduce[T, U any](in []T, identity U, combine func(curr U, next T) U) U {
	result := identity
	for _, next := range in {
		result = combine(result, next)
	}
	return result
}

type number interface {
	constraints.Float | constraints.Integer
}

func Sum[N number](in []N) N {
	return Reduce(in, identity[N](), func(curr N, next N) N {
		return curr + next
	})
}

func SumFrom[T any, N number](in []T, toNumber func(t T) N) N {
	return Reduce(in, identity[N](), func(curr N, next T) N {
		return curr + toNumber(next)
	})
}

func identity[N number]() N {
	return 0
}
