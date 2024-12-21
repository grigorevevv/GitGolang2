/*
4. Создайте два слайса из чисел типа int, объедините их в один слайс
и выведите результат. Используйте функцию append.
*/

package main

import "fmt"

func taskFour() {
	sliceOne := []int{20, 13, 113, -5, 77}
	sliceTwo := []int{17, 36, 99, 643, -32, 78}

	sliceNew := append(sliceOne[:], sliceTwo[:]...)

	for index, value := range sliceNew {
		fmt.Println(index, value)
	}
}
