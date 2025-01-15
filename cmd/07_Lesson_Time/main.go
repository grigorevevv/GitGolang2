package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Задание 1. Массив с текущим временем
		Создайте массив из трех элементов типа time.Time.
		Заполните его:
		1. Текущим временем (метод `time.Now()`)
		2. Временем, полученным добавлением 2 часов к текущему (метод `Add`)
		3. Временем, преобразованным в UTC (метод `UTC`)
		Выведите весь массив с помощью fmt.Println.
	*/
	fmt.Println("Задание 1. Массив с текущим временем")
	var arrayTime [3]time.Time
	dateFormat := "02.01.2006 15:04:05"
	now := time.Now()
	arrayTime[0] = now
	arrayTime[1] = now.Add(2 * time.Hour)
	arrayTime[2] = now.UTC()

	for element, tm := range arrayTime {
		switch element {
		case 0:
			fmt.Println("Текущее время - ", tm.Format(dateFormat))
		case 1:
			fmt.Println("Прибавим 2 часа - ", tm.Format(dateFormat))
		case 2:
			fmt.Println("Преобразуем в UTC - ", tm.Format("02.01.2006 15:04:05")) // оставил для сравнения
		}
	}
	fmt.Println()

	/*
		Задание 2. Слайс с длительностями
		Создайте слайс из трех элементов типа time.Duration.
		Заполните его длительностями: 30 минут, 1 час и 90 секунд.
		Добавьте к слайсу новую длительность — 2 часа (функция `append`).
		Выведите итоговый слайс с помощью fmt.Println.
	*/

	fmt.Println("Задание 2. Слайс с длительностями")

	sliceDuration := []time.Duration{
		time.Duration(30 * time.Minute),
		time.Duration(1 * time.Hour),
		time.Duration(90 * time.Second),
	}

	sliceDuration = append(sliceDuration, time.Duration(2*time.Hour))

	fmt.Println(sliceDuration)
	fmt.Println()

	/*
		Задание 3. Сравнение времени
		Создайте две переменные типа time.Time:
		1. Первая должна хранить текущее время (метод `time.Now()`)
		2. Вторая — время 3 часа назад (используйте метод Add с отрицательной длительностью)
		Выведите, какое из этих двух значений больше, используя метод After.
	*/

	fmt.Println("Задание 3. Сравнение времени")
	var tm1, tm2 time.Time

	tm1 = time.Now()
	tm2 = tm1.Add(-3 * time.Hour)

	fmt.Printf("Есть две даты - %s, %s\n", tm1.Format(dateFormat), tm2.Format(dateFormat))

	if tm1.After(tm2) {
		fmt.Printf("Ниабольшая из них - %s", tm1.Format(dateFormat))
	} else {
		fmt.Printf("Ниабольшая из них - %s", tm2.Format(dateFormat))
	}

	fmt.Println()

	/*
		Задание 4. Форматирование времени с массивами и слайсами
		1. Создайте массив из двух элементов типа time.Time: текущего времени и времени 1 день назад.
		2. Создайте слайс строк и заполните его, отформатировав элементы массива в виде:
		ГГГГ-ММ-ДД ЧЧ:ММ:СС (метод Format и шаблон `2006-01-02 15:04:05`).
		3. Выведите слайс с помощью fmt.Println.
		4. Добавьте новую строку в слайс, содержащую текущую дату и время в UTC, отформатированное аналогичным образом.
		5. Выведите обновленный слайс.
	*/
	fmt.Println("Задание 4. Форматирование времени с массивами и слайсами")
	arrTime := [2]time.Time{}
	//var arrTime [2]time.Time

	//now = time.Now()
	arrTime[0] = time.Now()
	arrTime[1] = time.Now().Add(-24 * time.Hour)

	fmt.Println(arrTime)

	slcTime := make([]string, 0, len(arrTime))

	for _, dt := range arrTime {
		slcTime = append(slcTime, dt.Format(dateFormat))
	}

	fmt.Println(slcTime)
	slcTime = append(slcTime, time.Now().UTC().Format(dateFormat))
	fmt.Println(slcTime)

	fmt.Println()

}
