package main

type Transfomer[T, S any] func(t T) S

func Map[T, S any](input []T, f Transfomer[T, S]) []S {
	result := []S{}

	for _, val := range input {
		result = append(result, f(val))
	}

	return result
}
