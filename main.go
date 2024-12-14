package main

import (
	"fmt"
	"strconv"
)

func main() {

	//taskOne()

	taskTwo()

}

/* Задача № 1
1. Работа с переменными и арифметикой
Объявите три переменные a, b, и c типа int. Пользователь вводит значения для a и b.
Присвойте переменной c результат выражения (a + b) * (a - b) и выведите все три переменные.
*/

func taskOne() {
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

/*
2. Функция с несколькими типами данных
Создайте функцию Calculate, которая принимает два числа типа float64 и строку operation.
Функция выполняет сложение, вычитание, умножение или деление в зависимости от переданного
значения operation. При делении на ноль должна возвращаться ошибка.
*/

func taskTwo() {
	var x, y float64
	var operStr string

	fmt.Print("Введите значение для переменной x: ")
	fmt.Scan(&x)
	fmt.Print("Введите значение для переменной y: ")
	fmt.Scan(&y)

	/*reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите текст: ")
	operStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}*/
	//fmt.Println("Вы ввели:", operStr)
	operStr = "сложение"
	fmt.Println(Calculate(x, y, operStr))
	operStr = "вычитание"
	fmt.Println(Calculate(x, y, operStr))
	operStr = "умножение"
	fmt.Println(Calculate(x, y, operStr))
	operStr = "деление"
	fmt.Println(Calculate(x, y, operStr))
	//fmt.Println("Значение переменной b - " + strconv.Itoa(b))
	//fmt.Println("Значение переменной c - " + strconv.Itoa(c))
}

func Calculate(x, y float64, operation string) float64 { // taskTwo

	if y == 0 {
		return 0
		//Println("Делить на ноль нельзя!"))
	}

	switch operation {
	case "сложение":
		return x + y
	case "вычитание":
		return x - y
	case "умножение":
		return x * y
	case "деление":
		return x / y
	default:
		return 0
	}

}
