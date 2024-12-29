/*
5. Напишите программу, которая принимает слайс из чисел типа float64,
вычисляет их среднее значение и выводит результат.
*/

package main

import "fmt"

func taskFive() {
	sliceFloat := []float64{64, 335, 47.98, -5, 19.99}

	var sum float64
	var countElem int = len(sliceFloat)

	for _, value := range sliceFloat {
		sum += value
		//countElem++
	}
	fmt.Printf("Среднее значение слайса: %.2f", sum/float64(countElem))
}
