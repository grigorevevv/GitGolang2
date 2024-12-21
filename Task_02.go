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
