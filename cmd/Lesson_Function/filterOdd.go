package main

func filterOdd(oldSlice []int) (newSlice []int) {

	for _, value := range oldSlice {
		if value%2 != 0 {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}
