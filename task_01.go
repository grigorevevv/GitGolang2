/*
1. Напишите программу, которая принимает массив из 10 целых чисел,
введенных пользователем, и выводит их сумму. Используйте цикл для перебора элементов массива.
*/

package main

import "fmt"

func taskOne() {
	var x int
	nums := [10]int{}
	fmt.Print("Введите значения для массива: ")

	for i := 0; i <= 9; i++ {
		fmt.Scan(&nums[i])
	}

	for _, number := range nums {
		x += number
		//fmt.Println(i, number)
	}
	fmt.Println(x)
}
