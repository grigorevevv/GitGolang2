package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var ErrResponse = errors.New("Ошибка запроса ")

func main() {
	/* 	Задача 1: Проверка доступности ресурса
		Описание:
	Напишите программу, которая отправляет HEAD запрос на заданный URL, чтобы проверить доступность ресурса.
	Если ресурс доступен (статус код 200–299), выведите сообщение об успешной проверке. Если нет — выведите код ошибки.
		Пример работы:
	Введите URL: https://jsonplaceholder.typicode.com       Ресурс доступен!
	Введите URL: https://example.invalid             		Ошибка: код ответа 404
	Подсказка: 	Используйте метод http.Head.
	*/

	fmt.Println("Задача 1: Проверка доступности ресурса")

	//sliceURL := []string{
	//	"https://jsonplaceholder.typicode.com",
	//	"https://example.invalid",
	//	"https://hakim.se/404",
	//	"https://yandex.ru/search/?text=что+такое+backend&lr=213",
	//	"https://jsonplaceholder.typicode.com/todos",
	//}
	//
	//for _, urls := range sliceURL {
	//	ok, err := checkResource(urls)
	//	if err != nil {
	//		fmt.Errorf("не удалось проверить доступность: %v", err)
	//	}
	//	if ok {
	//		fmt.Println("Ресурс доступен!")
	//	}
	//}

	fmt.Println()

	/* 	Задача 2: Отправка POST-запроса
		Описание:
	Реализуйте программу, которая отправляет POST запрос на сервер для создания новой задачи.
	Заголовок Content-Type должен быть установлен в application/json. Тело запроса передайте в виде строки.
	Пример данных для отправки:		{"userId":1,"title":"Купить молоко","completed":false}
	Пример работы:					Ответ сервера: код 201, задача создана.
	Подсказка:						Используйте URL: https://jsonplaceholder.typicode.com/todos.
	*/

	fmt.Println("Задача 2: Отправка POST-запроса")

	//fmt.Println(requestPost(sliceURL[4]))

	fmt.Println()

	/*  Задача 3: Пинг сервиса
		Описание:
	Напишите программу, которая отправляет GET запрос на заданный URL. Если сервер возвращает статус 200
	, программа выводит время отклика в миллисекундах. Если статус другой — выведите сообщение об ошибке.
	Пример работы:					Введите URL: https://jsonplaceholder.typicode.com				Ответ: 20 ms
	Введите URL: https://example.invalid				Ошибка: сервер недоступен.
	Подсказка:			Используйте time.Now() для вычисления времени запроса.
	*/

	fmt.Println("Задача 3: Пинг сервиса")

	//for _, urls := range sliceURL {
	//	resp, err := pingService(urls)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(resp)
	//}
	fmt.Println()
	/*
		Задача 4: Пагинация с GET-запросами
		Описание:
		Создайте программу, которая запрашивает список задач с сервера по страницам. На каждой странице отображайте до 5 задач.
		Добавьте возможность навигации с помощью команд next и prev.
		Пример работы (страница 1):

		Задача 1: Delectus aut autem
		Задача 2: Quis ut nam facilis et officia qui
		Задача 3: Fugiat veniam numquam impedit
		Задача 4: Et porro tempora
		Задача 5: Laboriosam mollitia et enim quasi adipisci

		Введите команду: next

		Подсказка: 	Используйте URL: https://jsonplaceholder.typicode.com/todos.
		Для эмуляции пагинации передавайте параметры ?_start=0&_limit=5, изменяя значения start при навигации
		baseURL := "https://jsonplaceholder.typicode.com/todos"
	*/

	fmt.Println("Задача 4: Пагинация с GET-запросами")

	for i := 0; i < 4; i++ {
		bodyText, pageOut := paginationRequest(i)
		fmt.Printf("страница %d\n%s", pageOut, bodyText)
	}

	//fmt.Println(paginationRequest(1))

	fmt.Println()

}

func checkResource(s string) (bool, error) {
	response, err := http.Head(s)

	if err != nil {
		return false, fmt.Errorf("Oшибка  : %v", err)

	}

	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		return true, nil
	} else {
		return false, nil
	}
}

func requestPost(url string) string {
	// Данные для отправки
	data := map[string]interface{}{
		"userId":    1,
		"title":     "Купить молоко",
		"completed": "false",
	}

	// Преобразуем данные в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Sprintln("Ошибка преобразования данных в JSON:", err)
	}

	//Создаём HTTP-запрос
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Sprintln("Ошибка отправки POST-запроса:", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode == http.StatusCreated {
		return fmt.Sprintln("Запрос выполнен успешно! Статус:", resp.Status)
	} else {
		return fmt.Sprintln("Запрос выполнен с ошибкой! Статус:", resp.Status)
	}

}

func pingService(s string) (string, error) {
	timeStart := time.Now()

	response, err := http.Head(s)
	timeEnd := time.Now()

	if err != nil {
		return "", ErrResponse
	}

	if response.StatusCode >= 200 && response.StatusCode <= 299 {

		duration := timeEnd.Sub(timeStart)
		return fmt.Sprintf("Ответ: %dms\n", duration.Milliseconds()), nil
	} else {
		return fmt.Sprintf("Ошибка: код ответа %d\n", response.StatusCode), nil
	}
}

func paginationRequest(page int) (string, int) {
	// определяем URL для обращения и его параметры
	baseURL := "https://jsonplaceholder.typicode.com/todos"
	params := url.Values{}

	pageStr := fmt.Sprintf("0%d", page*5)
	// для параметров задаются только их имена, значения оставляются пустыми
	params.Add("_start", pageStr)
	params.Add("_limit", "5")

	// преобразуем URL-строку в объект и заполняем параметры
	urlInstance, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Sprintln(err), page
	}
	urlInstance.RawQuery = params.Encode()

	// выполняем запрос и обрабатываем ответ
	response, err := http.Get(urlInstance.String())
	if err != nil {
		return fmt.Sprintln(err), page
	}
	fmt.Println(response.Status)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Sprintln(err), page
	}

	//fmt.Println(string(body))

	Datajson := make([]map[string]interface{}, 0)

	err = json.Unmarshal(body, &Datajson)
	if err != nil {
		fmt.Println(err)
	}

	var listPage string
	for _, bodyValue := range Datajson {
		listPage += fmt.Sprintf("Задача %.f: %s\n", bodyValue["id"], bodyValue["title"])
	}
	return listPage, page
}
