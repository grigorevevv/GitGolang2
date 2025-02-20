package main

import (
	"fmt"
)

func main() {
	values := []interface{}{42, "Go", 3.14, false, []int{1, 2, 3}, map[int]string{}}

	for _, v := range values {
		switch v.(type) {
		case int, float64:
			fmt.Println("Это число")
		case string:
			fmt.Println("Это строка")
		case bool:
			fmt.Println("Это булево значение")
		case []int:
			fmt.Println("Это slice")
		default:
			fmt.Println("Неизвестный или сложный тип")
		}
	}

	processValue(10)
	processValue(5.5)
	processValue("Golang")
	processValue([]byte{1, 2, 3}) // Тип не поддерживается

	fmt.Println()
	// -------------Задача 1: Анализ значений в слайсе----------------------------------------//
	fmt.Println("Задача 1: Анализ значений в слайсе")
	sliceValue := []interface{}{42, "hello", 3.14, "world", 100, 7.5, 8, "go"}
	fmt.Println(sliceValue)
	fmt.Println(AnalyzingValueSlice(sliceValue))
	fmt.Println()
	// ---------------------------------------------------------------------------------------//

	// -------------Задача 2: Проверка соответствия типов данных--------------------------------//
	fmt.Println("Задача 2: Проверка соответствия типов данных")
	sliceOne := []interface{}{42, "hello", 3.14}
	sliceTwo := []interface{}{"world", 100, 7.5, 8, "go"}

	for _, varOne := range sliceOne {
		for _, varTwo := range sliceTwo {
			fmt.Println(chekcingDataTypes(varOne, varTwo))
		}
	}
	//fmt.Println(chekcingDataTypes(a, b))
	fmt.Println()
	// ---------------------------------------------------------------------------------------//

}

func processValue(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("Удвоенное значение: %d\n", v*2)
	case float64:
		fmt.Printf("Квадрат числа: %.2f\n", v*v)
	case string:
		fmt.Printf("Длина строки: %d символов\n", len(v))
	default:
		fmt.Println("Тип не поддерживается")
	}
}

// Задача 1: Анализ значений в слайсе
func AnalyzingValueSlice(vSlice []interface{}) string {
	/*   Задача 1: Анализ значений в слайсе
	Описание:  Дан слайс с разными типами данных. Нужно посчитать, сколько в нём строк, целых чисел и дробных чисел
	            , а также вывести их средние значения для числовых типов.
	Что нужно сделать:    1. Создайте слайс с разными типами данных: []interface{}{42, "hello", 3.14, "world", 100, 7.5, 8, "go"}.
	                      2. Используйте switch type, чтобы определить, какие значения принадлежат к числовым типам (целые и дробные).
	                      3. Выведите количество строк, целых чисел и чисел с плавающей точкой.
	                      4. Посчитайте и выведите среднее значение отдельно для целых чисел и отдельно для чисел с плавающей точкой.
	Ожидаемый вывод:              Строк: 3
	                              Целых чисел: 3 (среднее: 50)
	                              Чисел с плавающей точкой: 2 (среднее: 5.32)
	*/
	var countInt, countFloat, countString, sumInt int
	var sumFloat float64
	for _, v := range vSlice {
		switch value := v.(type) {
		case int:
			countInt++
			sumInt += value
		case float64:
			countFloat++
			sumFloat += value
		case string:
			countString++
		}
	}
	avargeInt := sumInt / countInt
	avargeFloat := sumFloat / float64(countFloat)
	return fmt.Sprintf("Строк: %d\nЦелых чисел: %d (среднее: %d)\nЧисел с плавающей точкой: %d (среднее: %.2f)",
		countString, countInt, avargeInt, countFloat, avargeFloat)
}

// -------------Задача 2: Проверка соответствия типов данных---------------------------------//
func chekcingDataTypes(a, b interface{}) string {
	/*    Задача 2: Проверка соответствия типов данных
	  Описание:     Есть две переменные разного типа. Нужно проверить, можно ли их сложить,и если да — выполнить операцию,
	                если нет — вывести сообщение об ошибке.
	  Что нужно сделать:    1. Создайте две переменные: одна может быть int, float64 или string, другая — int, float64 или string.
	                        2. Проверьте, можно ли их сложить (числа между собой, строки между собой).
	                        3. Если сложение возможно, выведите результат. Если нет — выведите сообщение об ошибке.
	Ожидаемый вывод (пример для 10 и 3.5):          Сумма: 13.5
	Ожидаемый вывод (пример для "Go" и 10):         Ошибка: невозможно сложить разные типы данных
	*/
	var sumFloat float64
	var sumString string
	numberBool, stringBool := false, false

	switch value := a.(type) {
	case int:
		numberBool = true
		sumFloat += float64(value)
	case float64:
		numberBool = true
		sumFloat += value
	case string:
		stringBool = true
		sumString += value
	}

	switch value := b.(type) {
	case int:
		if numberBool {
			return fmt.Sprintf("Сумма: %.2f", sumFloat+float64(value))
		} else {
			return "Ошибка: невозможно сложить разные типы данных"
		}
	case float64:
		if numberBool {
			return fmt.Sprintf("Сумма: %.2f", sumFloat+value)
		} else {
			return "Ошибка: невозможно сложить разные типы данных"
		}
	case string:
		if stringBool {
			return fmt.Sprintf("Сумма строк: %s", sumString+value)
		}
	}

	return "Ошибка: невозможно сложить разные типы данных"
}
