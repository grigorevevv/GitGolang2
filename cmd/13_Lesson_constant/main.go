package main

import (
	"fmt"
	"unicode"
)

const taxRate = 0.2

func calculateTotal(price float64) float64 {
	/*
		Задача 1: Расчёт суммы покупки с налогом
		Описание:  Напишите программу, которая вычисляет итоговую стоимость товара с учётом налога.
		Требования:			1. Объявите константу taxRate = 0.2 (налог 20%).
							2. Напишите функцию calculateTotal(price float64) float64, которая добавляет налог к цене.
							3. Проверьте её на ценах: 500.00, 1200.50.
		Ожидаемый вывод:			Итоговая цена: 600.00
									Итоговая цена: 1440.60
	*/
	return price + price*taxRate
}

func modifyArray(arr *[6]int) {
	/*  Задача 2: Указатели и изменение массива
	Описание:   	Напишите программу, в которой каждое нечётное число в массиве заменяется на -1 через указатель.
	Требования:			1. Создайте массив [6]int{3, 8, 5, 12, 7, 14}.
						2. Напишите функцию modifyArray(arr *[6]int), которая меняет нечётные числа на -1.
						3. Выведите массив до и после изменений.
	Ожидаемый вывод:				До: [3 8 5 12 7 14]
									После: [-1 8 -1 12 -1 14]
	*/
	for indesx, number := range *arr { //  итерация по указатель ? зачем тогда указатель
		if number%2 > 0 {
			arr[indesx] = -1
		}
	}
}

func amountUniqueCharInString(text string) int {
	/*    Задача 3: Подсчёт уникальных символов в строке
	Описание: Напишите функцию, которая подсчитывает количество уникальных символов в строке.
	Требования:		1. Используйте map[rune]bool для хранения встреченных символов.
					2. Проверьте на строках: "hello", "абракадабра".
	Ожидаемый вывод:			Уникальные символы в "hello": 4
								Уникальные символы в "абракадабра": 5
	*/
	runeMap := map[rune]bool{}
	for _, word := range text {
		if unicode.IsLetter(word) {
			runeMap[word] = true
		}
	}
	return len(runeMap)
}

func calculatorConsumer(hours float64) string {
	/*  	Задача 4 (сложнее): Калькулятор потребления электроэнергии
	Описание:		Напишите программу, которая рассчитывает потребление электроэнергии бытовыми приборами.
	Формула: Энергопотребление = (мощность прибора в Вт * время работы в часах) / 1000 (в кВт⋅ч).
	Требования:   	1. Используйте map[string]float64 для хранения мощностей приборов (Вт).
					2. Примите время работы в массиве часов: [5]float64{2, 3, 4, 1.5, 5}.
					3. Рассчитайте суммарное потребление.
	Ожидаемый вывод:		Потребление:
								Холодильник: 3.00 кВт⋅ч
								Чайник: 4.50 кВт⋅ч
								Телевизор: 2.00 кВт⋅ч
								Фен: 1.50 кВт⋅ч
								Кондиционер: 10.00 кВт⋅ч
								Общее потребление: 21.00 кВт⋅ч
	*/
	powerDevices := map[string]float64{
		"Холодильник": 3000,
		"Чайник":      4500,
		"Телевизор":   2000,
		"Фен":         1500,
		"Кондиционер": 10000,
	}
	var sumConsumer float64
	result := fmt.Sprintf("Потребление за %.2f ч. :\n", hours)
	for device, power := range powerDevices {
		result += fmt.Sprintf("	%s: %.2f кВт⋅ч\n", device, (power*hours)/1000)
		sumConsumer += (power * hours) / 1000
	}

	return fmt.Sprintf("%s	Общее потребление: %.2f кВт⋅ч", result, sumConsumer)
}

func main() {
	priceList := []float64{500.00, 1200.50}

	fmt.Println("Задача 1: Расчёт суммы покупки с налогом")
	for _, p := range priceList {
		fmt.Println(calculateTotal(p))
	}
	fmt.Println()

	fmt.Println("Задача 2: Указатели и изменение массива")
	intArr := [6]int{3, 8, 5, 12, 7, 14}

	fmt.Println(intArr)
	modifyArray(&intArr)
	fmt.Println(intArr)

	fmt.Println()
	fmt.Println("Задача 3: Подсчёт уникальных символов в строке")

	slceWords := []string{"hello", "абракадабра", "Hello word!d."}
	for _, word := range slceWords {
		fmt.Printf("Уникальные символы в \"%s\": %d\n", word, amountUniqueCharInString(word))
	}
	fmt.Println()
	fmt.Println("Задача 4 (сложнее): Калькулятор потребления электроэнергии")

	hoursArray := [5]float64{2, 3, 4, 1.5, 5}
	for i := 0; i < len(hoursArray); i++ {
		fmt.Println(calculatorConsumer(hoursArray[i]))
	}

	fmt.Println()
}
