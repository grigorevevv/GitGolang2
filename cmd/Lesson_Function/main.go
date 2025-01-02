package main

import (
	"fmt"
)

type Student struct {
	name  string
	marks []float64
}

func main() {
	/*  Задача 1: Фильтрация нечётных чисел
	1. Напишите функцию filterOdd, которая принимает на вход слайс целых чисел (`[]int`) и возвращает новый слайс только из нечётных чисел.
	2. В main задайте исходный слайс чисел (например, 5–7 штук), вызовите filterOdd и выведите результат.
	> Подсказка:  > - Для проверки на нечётность используйте x % 2 != 0.   > - Для добавления элемента к результату используйте append.
	Пример:  Исходные числа: [3 4 7 10 9 1]   Нечётные числа: [3 7 9 1]                   */

	fmt.Println("Задача 1: Фильтрация нечётных чисел")
	apslice := []int{3, 4, 7, 10, 9, 1, 7, 9, 0, -3, 12}

	fmt.Println(filterOdd(apslice))
	fmt.Println()
	/*   Задача 2: Переключатель сообщений (switch)
	1. Создайте функцию printMessage(level int), которая:  - Если level == 1: выводит "Info message".  - Если level == 2: выводит "Warning message".  - Если level == 3: выводит "Error message".   - Иначе: выводит "Unknown level".
	2. В main несколько раз вызовите эту функцию с разными значениями (`1, 2, 3, 5`) и выведите результат.
	> Подсказка:   > - Используйте конструкцию switch level { ... }.  > - В Go нет автоматического «проваливания» в следующую ветку, если не указать fallthrough.   Пример:  printMessage(1) -> Info message   printMessage(2) -> Warning message   printMessage(3) -> Error message    printMessage(5) -> Unknown level	 */
	fmt.Println("Задача 2: Переключатель сообщений (switch)")
	printMessage(1)
	printMessage(2)
	printMessage(3)
	printMessage(5)
	fmt.Println()
	/*  Задача 3: Удаление элемента из слайса
	1. Напишите функцию removeElement(s []int, index int) []int, которая удаляет элемент с  индексом index из слайса s и возвращает новый слайс.
	2. В main создайте слайс из 6–7 чисел, выведите его.
	3. Вызовите removeElement для индекса 2, выведите обновлённый слайс.
	4. Попробуйте удалить элемент с индексом, выходящим за границы (например, 10) и корректно обработайте   эту ситуацию (например, вернуть исходный слайс или вывести сообщение).
	> Подсказка:  > - Используйте приём s = append(s[:index], s[index+1:]...), но предварительно проверьте, что index не выходит за len(s).
	Пример:   Исходный слайс: [10 20 30 40 50 60 70]    Удаляем элемент на индексе 2...    Результат:      [10 20 40 50 60 70]     */
	fmt.Println("Задача 3: Удаление элемента из слайса")
	sliceRemove, err := removeElement(apslice, 12)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sliceRemove)
	}
	fmt.Println()
	/*Задача 4: Средний балл студентов

	1. Создайте структуру Student с полями Name string и Marks []float64.
	2. Напишите функцию averageMark(s Student) float64, которая считает среднее арифметическое оценок Marks.
	3. В main создайте срез students (тип []Student`) из 3–4 студентов, каждому заполните `Name и Marks.
	4. Пройдитесь в цикле по students и для каждого студента: - Выведите его имя,  - Средний балл (применив `averageMark`), - Комментарий: если средний балл \>= 4.5 — «Отлично», если \>= 3 — «Хорошо», иначе — «Требуется подтянуть оценки».
	Пример (условно):
	Студент: Ivan,   оценки: [5 5 4 4], средняя: 4.50 -> Отлично
	Студент: Elena,  оценки: [3 3 4],   средняя: 3.33 -> Хорошо
	Студент: Den,    оценки: [2 2 3],   средняя: 2.33 -> Требуется подтянуть оценки  */
	fmt.Println("Задача 4: Средний балл студентов")

	studentSlice := []Student{
		{name: "Ivan", marks: []float64{5, 5, 4, 4}},
		{name: "Den", marks: []float64{3, 3, 4}},
		{name: "Donald", marks: []float64{5, 2, 3}},
		{name: "Elena", marks: []float64{2, 3, 5, 5, 4}},
		{name: "Vlad", marks: []float64{2, 1, 1, 3, 5}},
	}

	var comment string
	for _, valueSlice := range studentSlice {
		avgMark := averageMark(valueSlice.marks)
		switch {
		case avgMark >= 4.5:
			comment = "Отлично"
		case avgMark >= 3:
			comment = "Хорошо"
		default:
			comment = "Требуется подтянуть оценки"
		}
		fmt.Printf("Студент: %s, оценки: %.f, средняя: %.2f -> %s\n", valueSlice.name, valueSlice.marks, avgMark, comment)
	}
	fmt.Println()

	/*  Задача 5 (посложнее): Рекурсивный вывод чисел Фибоначчи
	1. Напишите функцию fibonacci(n int) int, которая рекурсивно возвращает *n*-е число Фибоначчи:  - fibonacci(0) = 0  - fibonacci(1) = 1   - fibonacci(n) = fibonacci(n-1) + fibonacci(n-2) (для `n > 1`)
	2. В main запросите у пользователя n или задайте его в коде.
	3. Выведите результат: fibonacci(n).
	4. При больших n программа может работать медленно (особенность рекурсии), обсудите или поэкспериментируйте с этим.
	Пример:
	Введите число: 5                   5-е число Фибоначчи = 5
	Введите число: 10                   10-е число Фибоначчи = 55       */
	fmt.Println("Задача 5: Рекурсивный вывод чисел Фибоначчи")

	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n) //    возвращает *n*-е число Фибоначчи
	switch {
	case n > -1:
		fmt.Printf("%d-е число Фибоначчи = %d\n", n, fibonacci(n))
	default:
		fmt.Printf("искомое число Фибоначчи не может быть отрицательным")
	}
	fmt.Println()

}
