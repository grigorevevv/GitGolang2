package main

/* Задача № 1
1. Работа с переменными и арифметикой
Объявите три переменные a, b, и c типа int. Пользователь вводит значения для a и b.
Присвойте переменной c результат выражения (a + b) * (a - b) и выведите все три переменные.
*/

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, c int
	fmt.Print("Введите значение для переменной a: ")
	fmt.Scan(&a)
	fmt.Print("Введите значение для переменной b: ")
	fmt.Scan(&b)

	c = (a + b) * (a - b)
	fmt.Println("Значение переменной a - " + strconv.Itoa(a))
	fmt.Println("Значение переменной b - " + strconv.Itoa(b))
	fmt.Println("Значение переменной c - " + strconv.Itoa(c))
}
