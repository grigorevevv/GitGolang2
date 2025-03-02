package main

import (
	"fmt"
	"strings"
	"unicode"
)

func analysisShoppingList(listSh []string) string {
	/* Задача 1: Анализ списка покупок
	    	Вы разрабатываете систему управления покупками. У вас есть слайс покупок, где каждая покупка — это строка с наименованием товара.
		   	Напишите программу, которая:	Подсчитывает количество уникальных товаров в списке.
											Определяет самый часто встречающийся товар.
			shoppingList := []string{"яблоки", "молоко", "хлеб", "яблоки", "сыр", "молоко", "яблоки", "хлеб", "сок", "сыр"}
			Ожидаемый вывод (пример):				Уникальных товаров: 5
		   											Самый популярный товар: яблоки (3 раза)
	*/
	mapShop := make(map[string]int)
	for _, product := range listSh {
		mapShop[product]++
	}
	var popularProduct string
	popularProductSlice := make([]string, 0, len(mapShop))
	var unique, maxUnique int
	maxBool := true
	for product, sum := range mapShop {
		//result += fmt.Sprintf("%s - %d\n", product, sum)
		unique++
		if maxBool {
			maxUnique = sum
			maxBool = false
		}
		if maxUnique < sum {
			maxUnique = sum
			popularProductSlice = nil
			popularProduct = fmt.Sprintf("%s (%d %s)", product, maxUnique, formatMaxUnique(maxUnique))
			popularProductSlice = append(popularProductSlice, popularProduct)
		} else if maxUnique == sum {
			popularProduct = fmt.Sprintf("%s (%d %s)", product, maxUnique, formatMaxUnique(maxUnique))
			popularProductSlice = append(popularProductSlice, popularProduct)
		}
		if len(popularProductSlice) > 1 {
			popularProduct = "Самые популярные товары: " + strings.Join(popularProductSlice, ", ")
		} else {
			popularProduct = "Самый популярный товар: " + strings.Join(popularProductSlice, ", ")
		}
	}
	return fmt.Sprintf("Уникальных товаров: %d\n%s",
		unique, popularProduct)
}

func formatMaxUnique(n int) string {
	if (n%10 == 2 || n%10 == 3 || n%10 == 4) && (n%100 < 10 || n%100 > 20) {
		return "раза"
	} else {
		return "раз"
	}
}

func verificationPassword(password string) string {
	/*	Задача 2: Проверка корректности пароля
		Напишите функцию, которая проверяет, является ли переданный пароль надежным.
		Пароль считается надежным, если:		Его длина не менее 8 символов.
												В нём есть хотя бы одна цифра.
												В нём есть хотя бы одна заглавная буква.
												В нём есть хотя бы один специальный символ (!@#$%^&*).
		Пример вызова:			passwords := []string{"hello", "Pass123", "Secure!123", "GoLang@2023"}
		Ожидаемый вывод:				"hello" - ненадежный
										"Pass123" - ненадежный
										"Secure!123" - надежный
										"GoLang@2023" - надежный
	*/
	passwordSlice := []rune(password)

	if len(passwordSlice) < 8 {
		return "ненадежный"
	}

	isDigitBool, isUpperBool, isContainsRune := false, false, false

	for _, word := range passwordSlice {
		if unicode.IsDigit(word) {
			isDigitBool = true
		}
		if unicode.IsUpper(word) {
			isUpperBool = true
		}
		if strings.ContainsRune("!@#$%^&*", word) {
			isContainsRune = true
		}
	}
	if isDigitBool && isUpperBool && isContainsRune {
		return "надежный"
	} else {
		return "ненадежный"
	}
}

func main() {

	/* Задача 1: Анализ списка покупок  ----------------------------------------------------------------------------- */
	fmt.Println("Задача 1: Анализ списка покупок")
	shoppingList := []string{"яблоки", "молоко", "хлеб", "яблоки", "сыр", "молоко", "яблоки", "хлеб", "сок", "сыр", "молоко",
		"молоко", "молоко"}

	fmt.Println(analysisShoppingList(shoppingList))
	fmt.Println()
	/*  --------------------------------------------------------------------------------------------------------------*/

	/* Задача 2: Проверка корректности пароля  ---------------------------------------------------------------------- */
	fmt.Println("Задача 2: Проверка корректности пароля")

	passwords := []string{"hello", "Pass123", "Secure!123", "GoLang@2023", "A3!A33#", "S1ecuresecure*"}

	for _, password := range passwords {
		fmt.Printf("Пароль: %s - %s\n", password, verificationPassword(password))
	}
}
