/*
2. Создайте функцию, которая принимает слайс из чисел типа int и возвращает
минимальное и максимальное значения. Программа должна вывести результаты в main.
*/

package main

import "fmt"

func taskTwo() {
	nums := []int{20, 13, 113, 15, 34, 7, 52, 41, -5}

	valuemax, valuemin := maxminvalueslice(nums)
	fmt.Printf("минимальное значение: %2d\n", valuemin)
	fmt.Printf("максимальное значение: %2d\n", valuemax)
}

func maxminvalueslice(nums []int) (int, int) {
	var maxvalue int
	var minvalue int
	minvalue = nums[0]
	maxvalue = nums[0]

	for _, number := range nums {
		if maxvalue < number {
			maxvalue = number
		}
		if minvalue > number {
			minvalue = number
		}
	}
	return minvalue, maxvalue
}
