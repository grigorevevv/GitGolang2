/*
1. Напишите программу, которая принимает массив из 10 целых чисел,
введенных пользователем, и выводит их сумму. Используйте цикл для перебора элементов массива.
*/

package main

import "fmt"

func taskOne() {
	var sum int
	nums := [10]int{}
	fmt.Print("Введите значения для массива: ")

	for i := 0; i < 10; i++ {
		_, err := fmt.Scan(&nums[i])
		if err != nil {
			fmt.Printf("ошибка при вводе элемента %d: %v\n", i, err)
		}
	}

	for _, number := range nums {
		sum += number
		//fmt.Println(i, number)
	}
	fmt.Println(sum)
}
