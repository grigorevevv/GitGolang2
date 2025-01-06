package main

func averageMark(s []float64) (sum float64) {

	counter := float64(len(s))
	for _, value := range s {
		sum += value
	}
	return sum / counter
}
