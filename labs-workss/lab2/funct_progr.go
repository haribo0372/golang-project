package main

func Map(slice []float64, predicate func(float64) float64) []float64 {
	array := make([]float64, len(slice))

	copy(array, slice)

	for i, v := range array {
		array[i] = predicate(v)
	}
	return array
}
