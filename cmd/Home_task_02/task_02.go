/*
Задание 2: Оценка температуры
Напишите программу, которая запрашивает у пользователя значение температуры (целое число)
и выводит сообщение в зависимости от её значения:
- Если температура ниже 0, напечатать: "Холодно".
- Если температура от 0 до 20 (включительно), напечатать: "Прохладно".
- Если температура выше 20, напечатать: "Тепло".
Используйте ветвление if-else.

Пример результата:    Введите температуру: 15
                       Прохладно
*/

package main

import "fmt"

func taskTwo() {
	var temperature int

	fmt.Println("Введите температуру: ")
	fmt.Scan(&temperature)

	if temperature < 0 {
		fmt.Println("Холодно")
	} else if temperature >= 0 && temperature <= 20 {
		fmt.Println("Прохладно")
	} else {
		fmt.Println("Тепло")
	}
	/*	case temperature < 0 :

			fmt.Println("Холодно")
		case temperature >= 0 && temperature <= 20:
			fmt.Println("Прохладно")
		case temperature > 20:
			fmt.Println("Тепло")
		}*/

}
