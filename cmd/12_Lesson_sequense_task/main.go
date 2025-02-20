package main

import "fmt"

/*
Описание:
Напишите программу, которая вычисляет площадь круга, используя константу π. Затем реализуйте функцию swap(a, b *int)
, которая меняет местами два числа по ссылке.
Требования:
1. Объявите константу pi = 3.141592.
2. Создайте функцию circleArea(radius float64) float64, которая возвращает площадь круга.
3. Объявите два целых числа, передайте их в функцию swap(a, b *int), и выведите результат до и после вызова.
Ожидаемый вывод (пример):
Площадь круга радиусом 5: 78.54
До swap: a = 10, b = 20
После swap: a = 20, b = 10
*/

const pi = 3.141592

func circleArea(radius float64) float64 {
	return radius * radius * pi
}

func swap(a, b *int) {
	*a, *b = *b, *a
	*a = *a * 5
}

func main() {
	radius := 5.0

	fmt.Printf("Площадь круга радиусом %.1f: %.1f\n", radius, circleArea(radius))

	a, b := 10, 20
	fmt.Printf("До swap: a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("После swap: a = %d, b = %d\n", a, b)

	brandName := "Champion"
	fmt.Println("Адрес переменной brandName:", &brandName)

	brandName = "Cat"
	fmt.Println("Адрес переменной brandName:", &brandName, "значение", brandName)

}
