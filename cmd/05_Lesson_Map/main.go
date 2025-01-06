package main

import (
	"fmt"
	"sort"
	"unicode"
)

type Student struct {
	Name   string
	Course int
}

func main() {
	/*
			Задача 1: Подсчёт частоты букв
			1. Напишите функцию countLetters(text string) map[rune]int, которая возвращает карту, где:
			Ключ (rune) — это буква (символ) из строки.
				Значение (int) — количество её вхождений.
			2. В main попросите пользователя ввести строку (или задайте её в коде), вызовите countLetters и выведите, сколько раз встречается каждая буква.
			3. Используйте цикл для перебора карты и if, чтобы пропускать пробелы, знаки пунктуации или решать, что именно считать «буквой».
			> Подсказка:
			Строка в Go — это набор рун.
				Для преобразования строки в срез рун — []rune(text).
				Считайте только символы, которые являются буквами (например, с помощью unicode.IsLetter из пакета unicode).
		Пример:
			Введите текст: Hello Go
		H: 1
		e: 1
		l: 2
		o: 2
		G: 1
	*/
	var textIn string

	textIn = "hello golang"

	for wordKey, wordValue := range countLetters(textIn) {
		fmt.Printf("%c: %d\n", wordKey, wordValue)
	}

	/*  Задача 2: Слияние двух map
	1. Создайте две карты map[string]int, которые, например, будут хранить данные о товарах и их ценах(любой набор: "apple": 100, "banana": 50 и т.д.).
	2. Напишите функцию mergeMaps(m1, m2 map[string]int) map[string]int, которая возвращает новую карту, содержащую всепары (ключ:значение) из m1 и m2.
	Если в обеих картах есть одинаковый ключ, пусть в результирующей карте будет значение из m2 (например, считаем,что m2 «главнее»).
	3. В main выведите итоговую карту.
	> Подсказка:   Результирующую карту создайте через make.
	Скопируйте в неё данные из m1, а затем из m2, «перезаписывая» при совпадении ключей.
		Пример (условный):   m1 = {"apple": 100, "banana": 50}   m2 = {"banana": 60, "pear": 150}
	Результат: {"apple": 100, "banana": 60, "pear": 150}      */
	var mapOne = map[string]int{
		"Apple":  55,
		"Banana": 35,
		"Orange": 40,
		"Pear":   45,
		"Kiwi":   60,
	}

	var mapTwo = map[string]int{
		"Grape":  70,
		"Apple":  50,
		"Peach":  80,
		"Banana": 30,
		"Mango":  90,
	}

	fmt.Println(mergeMaps(mapOne, mapTwo))

	/*	Задача 3: Слайс чисел и карта с результатами
		1. Напишите функцию analyzeNumbers(nums []int) map[string]int, которая:
		Принимает на вход срез целых чисел.
		Возвращает карту со строковыми ключами:
		"positive": количество положительных чисел,
		"negative": количество отрицательных чисел,
		"zero": количество нулей.

		2. В main создайте срез nums (10–12 целых чисел), вызовите analyzeNumbers(nums) и выведите полученную карту.
		3. Используйте циклы (for range) для обхода чисел и if-else для распределения по категориям (положительные,
		отрицательные, нули).
		> Подсказка:
		Не забудьте инициализировать значения ключей карты в 0, иначе придётся проверять существование ключей перед инкрементом.
		Можно делать так:
		res["positive"]++
		Пример:
		Исходный срез: [10, -2, 0, 7, 0, -5, 12]
		Результат: {"positive": 3, "negative": 2, "zero": 2}      */

	numbers := []int{10, -2, 0, 7, 0, -5, 12, 0, 67, 345, 0, -4, -98}

	fmt.Println(analyzeNumbers(numbers))

	/*
		Задача 4 (посложнее): Группировка студентов по курсам
		1. Создайте структуру Student с полями Name string и Course int (номер курса).
		2. В main сформируйте срез []Student из нескольких студентов (5–7 штук), у каждого заполните Name и Course.
		3. Напишите функцию groupByCourse(students []Student) map[int][]string, которая:
		Принимает срез студентов.
		Возвращает карту:
		Ключ типа int (номер курса),
		Значение типа []string — список имён студентов, которые учатся на этом курсе.
		4. Выведите результат таким образом: для каждого курса покажите список студентов.
		> Подсказка:
		В цикле for range по students берите stud.Course как ключ, а stud.Name добавляйте в срез map[key].
		Инициализируйте пустой срез, если по данному курсу ещё нет записей.
		Пример (условно):
		students = [
		  {Name: "Alice", Course: 1},
		  {Name: "Bob", Course: 2},
		  {Name: "Charlie", Course: 1},
		  {Name: "Diana", Course: 3},
		  ...
		]
		Результат (map):
		1 -> ["Alice", "Charlie"]
		2 -> ["Bob"]
		3 -> ["Diana"]
		...
	*/

	students := []Student{
		{Name: "Ivan", Course: 1},
		{Name: "Den", Course: 2},
		{Name: "Donald", Course: 3},
		{Name: "Elena", Course: 4},
		{Name: "Vlad", Course: 7},
		{Name: "Alice", Course: 1},
		{Name: "Bob", Course: 2},
		{Name: "Charlie", Course: 1},
		{Name: "Diana", Course: 3},
	}

	courses := groupByCourse(students)

	keys := make([]int, 0, len(courses))

	for key := range courses {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		fmt.Printf("%d -> %s\n", key, courses[key])
	}
}

func countLetters(text string) map[rune]int {

	var runeSlice []rune
	resultMap := make(map[rune]int)
	var counter int

	for _, r := range text {
		runeSlice = append(runeSlice, r)
	}

	for _, word := range runeSlice {
		if unicode.IsLetter(word) && len(resultMap) > 0 {
			counter = 1
			for words, _ := range resultMap {
				if words == word {
					counter = resultMap[word] + 1
				}
			}
			resultMap[word] = counter
		} else if unicode.IsLetter(word) {
			resultMap[word] = 1
		}
		//fmt.Printf("%d - %c - %d\n", index, word, word)
	}
	return resultMap
}

func mergeMaps(m1, m2 map[string]int) map[string]int {
	mergeMap := make(map[string]int)

	for name, cost := range m1 {
		mergeMap[name] = cost
	}

	for name, cost := range m2 {
		mergeMap[name] = cost
	}
	return mergeMap
}

func analyzeNumbers(nums []int) map[string]int {
	resultMap := make(map[string]int)

	for _, num := range nums {
		if num > 0 {
			resultMap["positive"]++
		} else if num < 0 {
			resultMap["negative"]++
		} else {
			resultMap["zero"]++
		}
	}
	return resultMap
}

func groupByCourse(students []Student) map[int][]string {
	resultMap := make(map[int][]string)

	for _, student := range students {
		for i := 1; i < 8; i++ {
			if student.Course == i {
				resultMap[i] = append(resultMap[i], student.Name)

			}
		}
	}

	for i := 1; i < 8; i++ {
		if len(resultMap[i]) == 0 {
			resultMap[i] = append(resultMap[i], "")
		}
	}
	return resultMap
}
