/*
Задание 4: Факториал числа
Напишите программу, которая вычисляет факториал числа n (произведение всех чисел от 1 до `n`).
Число n должно быть задано в коде, и его можно менять для проверки работы программы.

Пример результата для n = 5:
Факториал 5 равен 120

Пример результата для n = 7:
Факториал 7 равен 5040
*/

package main

import "fmt"

func taskFour() {
	var factorial, sum int

	factorial = 5
	sum = 1
	for i := 1; i <= factorial; i++ {
		sum = sum * i
		fmt.Println(sum)
	}
	fmt.Printf("Факториал %d равен %d", factorial, sum)
}
