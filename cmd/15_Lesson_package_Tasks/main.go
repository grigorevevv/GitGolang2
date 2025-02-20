package main

// импортируйте необходимые пакеты
import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {

	// -------------Задача 1: Напечатайте дату основания Яндекса-------------------------------------------//
	/*
		Импортируйте пакеты time и fmt. Напечатайте в консоль дату основания Яндекса 14.09.2000. Переменная с датой основания Яндекса
		уже задана. Вывод должен иметь вид Яндекс основан 14.09.2000. Вывод должен заканчиваться переводом строки.
		Для форматирования даты используйте time.Format(). Используйте метод для преобразования строки в дату.
		Найдите его в документации пакета time. Обработка возможной ошибки уже будет в коде.
	*/
	fmt.Println("Задача 1: Напечатайте дату основания Яндекса")

	bornYandex := "2000-09-14"
	// Ниже вызовите метод из пакета time
	timeBorn, err := time.Parse("2006-01-02", bornYandex)
	// Обработка возможной ошибки, которую возвращает метод
	if err != nil {
		panic(err)
	}
	// Здесь выводите в консоль
	fmt.Printf("Яндекс основан: %s\n", timeBorn.Format("02.01.2006"))
	fmt.Println()

	// -------------Задача 2: (сложнее): Калькулятор потребления электроэнергии-------------------------------------------//
	fmt.Println("Задача 2: (сложнее): Калькулятор потребления электроэнергии")
	powerDevices := map[string]float64{
		"Холодильник": 3000,
		"Чайник":      4500,
		"Телевизор":   2000,
		"Фен":         1500,
		"Кондиционер": 21000,
	}
	indexDevice := map[int]string{
		1: "Холодильник",
		2: "Чайник",
		3: "Телевизор",
		4: "Фен",
		5: "Кондиционер",
	}
	hoursOperatingDevice := [5]float64{2, 3, 4, 1.5, 5}

	fmt.Println(calculatorConsumer(powerDevices, indexDevice, hoursOperatingDevice))

	// -------------Задача 3: Преобразование типов и округление---------------------------------------------------------//
	fmt.Println("Задача 3: Преобразование типов и округление")

	sliceFloat := []float64{3.4, 5.6, 7.1, 2.8, 9.9}

	fmt.Println(conversionType(sliceFloat))
	fmt.Println(conversionTypeFmt(sliceFloat))

	fmt.Println()

	// -------------Задача 4 (сложнее): Автоматическое приведение типов в расчётах---------------------------------------------------------//
	fmt.Println("Задача 4 (сложнее): Автоматическое приведение типов в расчётах")

	fmt.Println(calculate(10, 3.5, "+"))
	fmt.Println(calculate(10, 3.5, "*"))
	fmt.Println(calculate("Go", 10, "+"))
	fmt.Println(calculate("Hello ", "word!", "+"))
	fmt.Println(calculate("Hello ", "word!", "*"))
	fmt.Println()

	// -------------Задача 5 : Найти разницу в датах и результат привести к определенному формату---------------------------------------------------------//
	fmt.Println("Задача 5 : Найти разницу в датах и результат привести к определенному формату")

	start := time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC)
	stop := time.Date(2025, 2, 16, 12, 0, 0, 0, time.UTC)

	fmt.Println(diffDate(start, stop))

	start = time.Date(2025, 2, 15, 12, 0, 0, 0, time.UTC)
	stop = time.Date(2025, 3, 16, 5, 0, 0, 0, time.UTC)

	fmt.Println(diffDate(start, stop))

	start = time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	stop = time.Date(2025, 3, 16, 15, 0, 0, 0, time.UTC)

	fmt.Println(diffDate(start, stop))

}

func calculatorConsumer(power map[string]float64, index map[int]string, hours [5]float64) string {

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
	Две мапы.Одна имена, другая мощность, все используют одинаковые индексы, данные вынести за функцию
	*/
	result := "Потребление: \n"
	//var sumConsumer float64                 ?? вопрос синтаксиса
	for i := 0; i < len(hours); i++ {
		sumConsumer := (hours[i] * power[index[i+1]]) / 1000
		result += fmt.Sprintf("%s: %.2f кВт⋅ч\n", index[i+1], sumConsumer)
	}
	return result
}

func conversionType(slFl []float64) []int {
	/*  Задача 3: Преобразование типов и округление
	Описание:   Есть слайс дробных чисел, но в системе хранятся только целые. Нужно преобразовать их, учитывая правила округления.
	Что нужно сделать:		1. Создайте слайс дробных чисел: [3.4, 5.6, 7.1, 2.8, 9.9].
							2. Напишите код, который преобразует их в целые числа по следующим правилам:
								Если дробная часть ≥ 0.5, округлять вверх. Если дробная часть < 0.5, округлять вниз.
							3. Выведите преобразованные числа.
	Ожидаемый вывод:				Исходные: [3.4 5.6 7.1 2.8 9.9]
									Округлённые: [3 6 7 3 10]
	*/
	resultSlice := make([]int, 0, len(slFl))
	for _, number := range slFl {
		resultSlice = append(resultSlice, int(math.Round(number)))
	}
	return resultSlice
}

func conversionTypeFmt(slFl []float64) []int { // не стоит так делать
	resultSlice := make([]int, 0, len(slFl))

	for _, number := range slFl {
		i, _ := strconv.Atoi(fmt.Sprintf("%.f", number))
		resultSlice = append(resultSlice, i)
	}
	return resultSlice
}

func calculate(a, b interface{}, op string) string {
	/* Задача 4 (сложнее): Автоматическое приведение типов в расчётах
	Описание: Некоторые системы требуют строгой типизации. Нужно реализовать автоматическое приведение типов для операций сложения и умножения.
	Что нужно сделать:		1. Напишите функцию calculate(a interface{}, b interface{}, op string) interface{}.
							2. Если a и b — числа (int или float64), выполняйте операцию (+ или *).
							3. Если a и b строки, объединяйте их при +, но выдавайте ошибку при *.
							4. Если переданы несовместимые типы, возвращайте ошибка: несовместимые типы.
	Ожидаемый вывод (пример вызова calculate(10, 3.5, "+")):				Результат: 13.5
	Ожидаемый вывод (пример вызова calculate("Go", 10, "+")):				Ошибка: несовместимые типы
	*/
	var numberOne float64
	var sumString string
	numberBool, stringBool := false, false

	switch value := a.(type) {
	case int:
		numberBool = true
		numberOne = float64(value)
	case float64:
		numberBool = true
		numberOne = value
	case string:
		stringBool = true
		sumString = value
	}

	switch value := b.(type) {
	case int:
		if numberBool && op == "+" {
			return fmt.Sprintf("Сумма: %.2f", numberOne+float64(value))
		} else if numberBool && op == "*" {
			return fmt.Sprintf("Сумма: %.2f", numberOne*float64(value))
		} else {
			return "Ошибка: несовместимые типы"
		}
	case float64:
		if numberBool && op == "+" {
			return fmt.Sprintf("Сумма: %.2f", numberOne+value)
		} else if numberBool && op == "*" {
			return fmt.Sprintf("Сумма: %.2f", numberOne*value)
		} else {
			return "Ошибка: несовместимые типы"
		}
	case string:
		if stringBool && op == "+" {
			return fmt.Sprintf("Сумма строк: %s", sumString+value)
		}
	}

	return "Ошибка: несовместимые типы"
}

func diffDate(begin, end time.Time) string {
	var hours, days, month int

	duration := end.Sub(begin)


		monthName := begin.Month()


		switch monthName {
		case 1, 3, 5, 7, 8, 10, 12:
			month = int(duration.Hours() / (24 * 31))
			days = (int(duration.Hours()) % (24 * 31)) / 24
		case 4, 6, 9, 11:
			month = int(duration.Hours() / (24 * 30))
			days = (int(duration.Hours()) % (24 * 30)) / 24
		case 2:
			month = int(duration.Hours() / (24 * 28))
			days = (int(duration.Hours()) % (24 * 28)) / 24
		}

		hours = int(duration.Hours()) % 24

		return fmt.Sprintf("продолжительность %d мес. %d дн. %d час. %d кол.часов, %d кол.дней",
			month, days, hours, int(duration.Hours()), int(duration.Hours()/24))



}
