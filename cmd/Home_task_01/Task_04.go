/*
4. Создайте два слайса из чисел типа int, объедините их в один слайс
и выведите результат. Используйте функцию append.
*/

package main

import "fmt"

func taskFour() {
	sliceOne := make([]int, 3, 10)
	sliceOne[0] = 1
	sliceOne[1] = 2
	sliceOne[2] = 7
	sliceTwo := []int{17, 36, 99, 643, -32, 78}

	//sliceNew
	sliceNew := append(sliceOne, sliceTwo...)

	for index, value := range sliceOne {
		fmt.Println(index, value)
	}
	fmt.Println(sliceNew)
}
