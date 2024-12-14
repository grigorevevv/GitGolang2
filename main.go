package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	//taskOne()

	//taskTwo()

	//taskThree()

	//taskFour()

	taskFive()

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
	}
	fmt.Println("Вы ввели:", operStr)  */
	// Функция не видит входящий string

	operStr = "сложение"
	fmt.Println(Calculate(x, y, operStr))
	operStr = "вычитание"
	fmt.Println(Calculate(x, y, operStr))
	operStr = "умножение"
	fmt.Println(Calculate(x, y, operStr))
	operStr = "деление"
	fmt.Println(Calculate(x, y, operStr))

}

func Calculate(x, y float64, operation string) string { // taskTwo

	if y == 0 {
		//return 0
		return fmt.Sprint("Делить на ноль нельзя!")
	}

	switch operation {
	case "сложение":
		return fmt.Sprintf("сложение чисел %.2f и %.2f равно %.2f", x, y, x+y)
	case "вычитание":
		return fmt.Sprintf("разница чисел %.2f и %.2f равно %.2f", x, y, x-y)
	case "умножение":
		return fmt.Sprintf("умножение чисел %.2f и %.2f равно %.2f", x, y, x*y)
	case "деление":
		return fmt.Sprintf("деление чисел %.2f и %.2f равно %.2f", x, y, x/y)
	default:
		return fmt.Sprintf("%.2f", 0)
	}
}

/*
3. Работа с дробными числами и округлением
Программа принимает два числа типа float64 — длину и ширину прямоугольника.
Вычисляет площадь и периметр прямоугольника. Результаты округляются до двух
знаков после запятой.
*/

func taskThree() {

	var length, width float64

	fmt.Print("Введите значение длины прямоугольника length: ")
	fmt.Scan(&length)
	fmt.Print("Введите значение ширины прямоугольника width: ")
	fmt.Scan(&width)

	fmt.Printf("Периметр прямоугльника = %.2f\n", 2*(length+width))
	fmt.Printf("Площадь прямоугльника = %.2f\n", length*width)

}

/*=========================================================================================*/
/* 4. Приведение типов
Напишите программу, которая запрашивает у пользователя целое число, переводит его в float64,
делит на 3 и преобразует результат обратно в int. Выведите результат преобразования.*/
/*=========================================================================================*/
func taskFour() {
	var x, result int

	fmt.Print("Введите значение переменной x: ")

	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
	} else {
		fmt.Println("Вы ввели:", x)
	}

	result = int(float64(x) / 3)

	fmt.Println("Результат преобразования -", result)
}

/*=========================================================================================*/
/* 5. Работа с функциями и ошибками
Создайте функцию SquareRoot, которая принимает число типа float64 и возвращает
квадратный корень числа. Если передано отрицательное число, функция должна возвращать ошибку.*/
/*=========================================================================================*/
func taskFive() {
	var x float64

	fmt.Print("Введите значение переменной x: ")

	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
	} else if x < 0 {
		fmt.Println("Ошибка ввода: отрицательное число")
	} else {
		fmt.Println("Квадратный корень числа", x, "=", SquareRoot(x))

	}
}

func SquareRoot(x float64) float64 {
	return math.Sqrt(x)
}
