package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	/*
		Задача 1: Подсчёт гласных и согласных
		1. Напишите функцию: 	func countVowelsAndConsonants(s string) (int, int)
		Которая возвращает два числа: Количество гласных (vowels) и Количество согласных (consonants)
		2. Считаем, что мы работаем только с английским алфавитом (включая верхний и нижний регистр).
		3. В main протестируйте на разных строках (например, "Hello World", "GoLang").
		> Подсказка:
		Гласные: A, E, I, O, U (и строчные a, e, i, o, u).
		Для проверки, является ли символ буквой, используйте unicode.IsLetter(rune).
		Не включайте символы вроде пробелов, цифр, знаков препинания и т. д.
		Пример:
		"Hello"
		Гласных = 2 (e, o)
		Согласных = 3 (H, l, l)
	*/
	fmt.Println("Задача 1: Подсчёт гласных и согласных")
	var text string

	text = "Hello golang "
	vowels, consonants := countVowelsAndConsonants(text)
	fmt.Printf("В слове %s, количество гласных - %d, и количество согласных - %d\n", text, vowels, consonants)
	fmt.Println()

	/*
			Задача 2: Поиск подстроки
			1. Напишите функцию: func findOccurrences(s, sub string) int   Которая возвращает, сколько раз подстрока sub встречается в строке s.
			2. В main проверьте на разных примерах. Например: s = "banana"   sub = "ana"   Результат: 2.
		> Подсказка:  Реализуйте поиск вручную, перебирая индексы от 0 до len(s) - len(sub).
		Либо используйте готовые методы из пакета strings, но лучше для понимания сделать самостоятельно.
		Уточните, нужна ли «перекрывающаяся» логика (в "aaaa" подстрока "aa" встречается либо 3, либо 2 раза).
	*/

	fmt.Println("Задача 2: Поиск подстроки")
	fmt.Println(findOccurrences(text, "l")) //  вручную, перебирая индексы
	fmt.Println()
	fmt.Println(findOccurrences_2(text, "l")) // методом из пакета strings.Split
	fmt.Println()
	fmt.Println(strings.Count(text, "l")) // методом из пакета strings.Count
	fmt.Println()

	/*
		Задача 3: Сортировка слов в предложении
		1. Задайте строку, например: text := "go is an open source programming language"
		2. Разбейте её на слова: words := strings.Split(text, " ")
		3. Отсортируйте массив слов в лексикографическом порядке:  sort.Strings(words)
		4. Выведите результат в main.
		Пример:
		Исходная строка: "go is an open source programming language"
		Массив слов (до сортировки):
		["go", "is", "an", "open", "source", "programming", "language"]
		После сортировки:
		["an", "go", "is", "language", "open", "programming", "source"]
	*/
	text = "go is an open source programming language"
	fmt.Println("Задача 3: Сортировка слов в предложении")
	fmt.Println(sortWordSlice(text))
	fmt.Println()

	/*
		Задача 4 (посложнее): Статистика слов и сортировка по частоте
		1. Напишите функцию:  func wordFrequency(text string) map[string]int  Которая разбивает строку на слова и возвращает мапу «слово -> число вхождений».
		2. В main: Возьмите или спросите у пользователя строку,  Вызовите wordFrequency, получите мапу с подсчётом каждого слова.
		3. Отсортируйте слова по убыванию количества (Value) и выведите в порядке «самое частое слово» → «самое редкое слово».
		При одинаковой частоте выводите в любом порядке (или дальше сортируйте по алфавиту при равенстве).
		> Подсказка:  Считайте слова, например, через:  words := strings.Split(text, " ")  Делайте freq[word]++ для каждого слова.
		Для сортировки по значению в мапе используйте дополнительную структуру и sort.Slice.
		Пример:  Текст:  "banana apple apple banana banana pear"
		Статистика (до сортировки):
		{
		  "banana": 3,
		  "apple": 2,
		  "pear": 1
		}

		После сортировки:
		banana -> 3
		apple  -> 2
		pear   -> 1
	*/
	fmt.Println("Задача 4: Статистика слов и сортировка по частоте")

	text = "banana apple cherry apple cherry banana pear cherry"
	fmt.Println("До сотрировки")
	freq := wordFrequency(text)
	fmt.Println(freq)

	//--------  Первый способ----------------------------
	keys := make([]int, 0, len(freq))

	for _, words := range freq {
		keys = append(keys, words)
	}
	sort.Ints(keys)
	fmt.Println("После сотрировки")

	for i := len(keys); i > 0; i-- {
		for key, count := range freq {
			if i == count {
				fmt.Printf("%s -> %d\n", key, count)
			}
		}
	}

	fmt.Println()

	//--------  Второй способ----------------------------
	fmt.Println("После сотрировки - 2")
	type key_value struct {
		Key   string
		Value int
	}

	var sorted_struct []key_value

	for key, value := range freq {
		sorted_struct = append(sorted_struct, key_value{key, value})
	}

	sort.Slice(sorted_struct, func(i, j int) bool {
		return sorted_struct[i].Value > sorted_struct[j].Value
	})

	for _, Fruits_quantity := range sorted_struct {
		fmt.Printf("%s -> %d\n", Fruits_quantity.Key, Fruits_quantity.Value)
	}

	//fmt.Println(sorted_struct)

}

func countVowelsAndConsonants(s string) (int, int) {
	var v, c int
	vowels := "aeiouAEIOU"

	for _, r := range s {
		if unicode.IsLetter(r) && strings.ContainsRune(vowels, r) {
			v++
		} else if unicode.IsLetter(r) {
			c++
		}
	}
	return v, c
}

func findOccurrences(s, sub string) int {
	var count, sum int

	for i := 0; i < (len(s) - len(sub)); i++ {
		count = 0
		for y := 0; y < len(sub); y++ {
			if s[i+y] == sub[0+y] {
				count++
			}
		}
		if count == len(sub) {
			sum++
		}
	}
	return sum
}

func findOccurrences_2(s, sub string) int {
	//t := strings.Split(s, sub)
	//result := (len(t) - 1)
	//return result
	return len(strings.Split(s, sub)) - 1
}

func sortWordSlice(textin string) []string {
	result := strings.Split(textin, " ")
	sort.Strings(result)
	return result
}

func wordFrequencyTEST(text string) map[string]int {

	resultMap := make(map[string]int)
	//var resultmap map[string]int

	resultMap["word"] = 1

	return resultMap
}

func wordFrequency(text string) map[string]int {

	freq := make(map[string]int)

	result := strings.Split(text, " ")

	for _, word := range result {
		freq[word]++
	}

	return freq
}
