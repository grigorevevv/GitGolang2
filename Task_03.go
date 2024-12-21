/*
3. Напишите программу, которая создает слайс из 5 целых чисел,
удаляет элемент с индексом 2 и выводит новый слайс
*/

package main

import "fmt"

func taskThree() {
	nums := []int{20, 13, 113, -5, 77}
	var indexdelete int = 2

	sliceNew := append(nums[:indexdelete], nums[(indexdelete+1):]...)

	fmt.Println(sliceNew)

}
