/*
Задание 3: Таблица умножения числа
Напишите программу, которая выводит таблицу умножения для числа 7.
Используйте цикл for.

Пример результата:
7 x 1 = 7
7 x 2 = 14
7 x 3 = 21
...
7 x 10 = 70
*/

package main

import "fmt"

func taskThree() {
	var value, result int

	value = 7
	for i := 1; i <= 10; i++ {
		result = value * i
		fmt.Printf("%d x %d = %d\n", value, i, result)
	}
}
