package main

import (
	"fmt"
	"net/url"
)

func main() {
	/*Домашняя работа 1: Анализ URL
	1. Напишите программу, которая принимает URL в виде строки и разбирает его на компоненты: протокол, сервер, путь и параметры.
	2. Если параметры присутствуют, выведите их в формате ключ=значение на отдельной строке.
	Пример ввода : https://yandex.ru/search/?text=как%20стать%20разработчиком&lr=213
	Ожидаемый вывод:
				Протокол: https
				Сервер: yandex.ru
				Путь: /search/
				Параметры:
				  text = как%20стать%20разработчиком
				  lr = 213
	*/

	fmt.Println("Задача 1: Анализ URL")

	rawURL := "https://yandex.ru/search/?text=как%20стать%20разработчиком&lr=213"

	// Разбираем URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Ошибка разбора URL:", err)
		return
	}

	// Извлекаем компоненты
	protocol := parsedURL.Scheme    // Протокол (например, https)
	server := parsedURL.Host        // Сервер (например, yandex.ru)
	path := parsedURL.Path          // Путь (например, /search/)
	parameters := parsedURL.Query() // Параметры (например, map[text:[как%20стать%20разработчиком] lr:[213]])

	// Выводим результат
	fmt.Println("Протокол:", protocol)
	fmt.Println("Сервер:", server)
	fmt.Println("Путь:", path)

	if len(parameters) > 0 {
		fmt.Println("Параметры:")
		for key, values := range parameters {
			for _, value := range values {
				fmt.Printf("  %s = %s\n", key, value)
			}
		}
	} else {
		fmt.Println("Параметры: отсутствуют")
	}

	fmt.Println()

	/*	Домашняя работа 2: Кодирование строки для URL
		1. Напишите функцию, которая принимает строку запроса от пользователя и возвращает её закодированный вариант для использования в URL.
		2. Учтите, что пробелы должны заменяться на %20, а остальные символы кодироваться в формате UTF-8.
		Пример вызова: query := "как стать разработчиком"
						encodedQuery := encodeQuery(query)
						fmt.Println(encodedQuery)
		Ожидаемый вывод:  как%20стать%20разработчиком
	*/

	fmt.Println("Задача 2: Кодирование строки для URL")

	query := "как стать разработчиком"
	params := url.Values{}
	params.Add("search", query)

	encodedQuery := params.Encode()

	fmt.Println(encodedQuery)
	fmt.Println()

	/*  Домашняя работа 3: Генерация URL для поиска
	1. Создайте программу, которая принимает строку запроса от пользователя и формирует полный URL для поиска на Яндексе.
	2. Используйте закодированный запрос в параметре text.
	Пример вызова: query := "как стать разработчиком"
					url := generateSearchURL(query)
					fmt.Println(url)
	Ожидаемый вывод:  https://yandex.ru/search/?text=как%20стать%20разработчиком
	*/

	fmt.Println("Задача 3: Генерация URL для поиска")
	baseURL := "https://yandex.ru/search/"
	params = url.Values{}

	params.Add("text", query)

	urlInstance, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	urlInstance.RawQuery = params.Encode()
	fmt.Println(urlInstance.String())
	fmt.Println()

	/*	Домашняя работа 4: Проверка корректности URL
				1. Напишите программу, которая проверяет корректность введённого URL.
				2. Проверьте, содержит ли URL следующие обязательные компоненты: протокол, сервер и путь.
				3. Если какого-то компонента не хватает, выведите сообщение об ошибке
			Пример ввода : https://example.com
		Ожидаемый вывод:   Протокол: https
							Сервер: example.com
							Путь: /
							Параметры: отсутствуют


	*/

	fmt.Println("Задача 4: Проверка корректности URL")
	queryURL := "https://example.com/search/"

	urlInstancenew, err := url.Parse(queryURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	protocol = urlInstancenew.Scheme // Протокол (например, https)
	server = urlInstancenew.Host     // Сервер (например, yandex.ru)
	path = urlInstancenew.Path       // Путь (например, /search/)
	parameters = urlInstancenew.Query()

	if len(protocol) == 0 {
		fmt.Println("Ошибка, нет протокола")
		return
	}
	if len(server) == 0 {
		fmt.Println("Ошибка, нет сервера")
		return
	}
	if len(path) == 0 {
		fmt.Println("Ошибка, нет пути")
		return
	}
	fmt.Println("Протокол:", protocol)
	fmt.Println("Сервер:", server)
	fmt.Println("Путь:", path)

	if len(parameters) > 0 {
		for k, v := range parameters {
			for _, value := range v {
				fmt.Printf("  %s = %s\n", k, value)
			}
		}
	} else {
		fmt.Println("параметры отсутствуют")
	}

	fmt.Println()

}
