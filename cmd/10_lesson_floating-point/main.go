package main

import (
	"fmt"
)

func main() {
	sliceNumber := []int{3, 5, -10, 15, 17, 0, -21, 0, 1, 1, 1, 15, 7, 7, 7, 13, 7, 7, 12, 14, 16}

	//fmt.Println("Задача 1: Проверка делимости числа")
	//for _, number := range sliceNumber {
	//	//n, err := checkDivision(number)
	//	//if err != nil {
	//	//	fmt.Println(err)
	//	//	continue
	//	//}
	//	fmt.Printf("Число %d %s\n", number, checkDivision(number))
	//}
	//fmt.Println()
	//
	//fmt.Println("Задача 2: Проверка диапазона")
	//for _, number := range sliceNumber {
	//	fmt.Printf("Число %d %s\n", number, rangeCheck(number))
	//}
	//fmt.Println()
	//
	//fmt.Println("Задача 3: Вычисление среднего")
	//
	//for i := 0; i < len(sliceNumber)-2; i += 3 {
	//	n1, n2, n3 := sliceNumber[i], sliceNumber[i+1], sliceNumber[i+2]
	//	fmt.Printf("Из чисел %d,%d,%d - %s\n", n1, n2, n3, averageCalculate(n1, n2, n3))
	//}
	//fmt.Println()

	fmt.Println("Задача 4: Подсчёт простых чисел")

	sliceTwo := []int{2, 3, 5, 7}
	sliceThree := []int{}
	//fmt.Println(amountNumbers(sliceNumber))
	//fmt.Println(amountNumbers(sliceTwo))
	//fmt.Println(amountNumbers(sliceThree))

	fmt.Println(sliceNumber)
	fmt.Println(amountNumbers(sliceNumber))
	fmt.Println(amountNumbers(sliceTwo))

	for i := 2; i <= 10; i++ {
		sliceThree = append(sliceThree, i)
	}
	fmt.Println(amountNumbers(sliceThree))
	sliceFour := []int{4, 6, 8}
	fmt.Println(amountNumbers(sliceFour))
	fmt.Println()
}

func checkDivision(dividend int) string {
	/*   Напишите программу, которая принимает целое число и проверяет:
	1. Делится ли оно на 3;
	2. Делится ли оно на 5;
	3. Делится ли оно на 3 и на 5 одновременно.
	Если число делится на 3 и на 5, программа должна вывести:     Число делится на 3 и 5.
	Если только на одно из чисел:                                 Делится на 3 или Делится на 5.
	Если ни на одно из чисел:                                     Число не делится на 3 и 5.
	Подсказка: вспоминаем операцию "%"
	*/
	//if dividend == 0 {
	//	return "", errors.New("0 не делится!")
	//}

	if dividend%3 == 0 && dividend%5 == 0 {
		return "делится на 3 и 5"
	} else if dividend%3 == 0 && dividend%5 != 0 {
		return "Делится на 3"
	} else if dividend%3 != 0 && dividend%5 == 0 {
		return "Делится на 5"
	} else {
		return "не делится на 3 и 5"
	}
	return ""
}

func rangeCheck(number int) string {
	/*   Напишите программу, которая принимает на вход целое число и проверяет:
	1. Попадает ли оно в диапазон от -10 до 10 включительно.
	2. Если число больше 10, выведите сообщение:     	Число больше 10.
	3. Если число меньше -10, выведите сообщение:    	Число меньше -10.
	4. Если число в диапазоне от -10 до 10, выведите:	Число находится в пределах [-10, 10].
	*/
	if number >= -10 && number <= 10 {
		return "находится в пределах [-10, 10]"
	}
	if number < -10 {
		return "меньше -10"
	}
	return "больше 10"
}

func averageCalculate(numberOne, numbertwo, numberThree int) string {
	/*   Напишите программу, которая принимает три целых числа и вычисляет их среднее значение.
	1. Если все три числа одинаковы, программа должна вывести:        					Все числа одинаковы.
	2. Если хотя бы два числа одинаковы, но не все три, выведите:     					Два числа одинаковы.
	3. Если все числа разные, программа должна вывести среднее значение в формате:   	Среднее значение: <число>.
	*/

	if numberOne == numbertwo && numberOne == numberThree {
		return "Все числа одинаковы"
	} else if numberOne == numbertwo || numberOne == numberThree || numbertwo == numberThree {
		return "Два числа одинаковы"
	} else {
		average := float64(numberOne+numbertwo+numberThree) / 3
		return fmt.Sprintf("Среднее значение: %.2f", average)
	}
}

func amountNumbers(sliceInt []int) string {
	/*  Напишите программу, которая подсчитывает количество простых чисел в заданном диапазоне от 1 до n (включительно).
	1. Если n меньше 2, программа должна вывести:     Простых чисел нет.
	2. Если n больше или равно 2, программа должна вывести количество простых чисел и их список.
	Пример вывода: 		Простые числа: 2, 3, 5, 7
						Количество простых чисел: 4.
	*/

	result := ""
	counter := 0
	for _, number := range sliceInt {

		if isPrime(number) {
			result += fmt.Sprintf("%d - Простое число\n", number)
			counter++
		} else {
			result += fmt.Sprintf("%d - Не простое число\n", number)
		}

	}
	if counter > 0 {
		result += fmt.Sprintf("Количество простых чисел - %d\n", counter)
	} else {
		result += "Простых чисел нет"
	}

	return result
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
