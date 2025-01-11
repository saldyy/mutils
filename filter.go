package main

type Predicate[A any] func(A) bool

func Filter[T any](array []T, f Predicate[T]) []T {
	result := []T{}

	for _, val := range array {
		if f(val) {
			result = append(result, val)
		}
	}

	return result
}
