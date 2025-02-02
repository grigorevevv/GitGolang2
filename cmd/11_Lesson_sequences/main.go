package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func main() {

	fmt.Println("Задача 1: Обработка заказов в интернет-магазине")
	orders := []map[string]string{
		{"name": "Ноутбук", "status": "новый"},
		{"name": "Кофеварка", "status": "доставлен"},
		{"name": "Телефон", "status": "в обработке"},
		{"name": "Монитор", "status": "новый"},
		{"name": "Мышь", "status": "отменён"},
	}

	//statOrder := statisticOrders(orders)
	for statusName, count := range statisticOrders(orders) { // функция одни раз выполняется???
		fmt.Printf("%s: %d\n", statusName, count)
	}

	fmt.Println()

	fmt.Println("Задача 2: Работа с массивом температур")
	temps := [7]float64{12.5, 14.3, 10.8, 9.5, 8.1, 11.7, 13.2}

	fmt.Println(statisticTemperature(temps))
	fmt.Println()

	fmt.Println("Задача 3: Калькулятор команд в турнире")
	teams := map[string]int{
		"Львы":  10,
		"Тигры": 15,
		"Орлы":  12,
		"Акулы": 18,
		"Волки": 14,
	}

	fmt.Println(calculatorCommand(teams))
	fmt.Println()

	fmt.Println("Задача 4: Распределение студентов по группам")
	students := []string{"Алиса", "Борис", "Виктор", "Глеб", "Дмитрий", "Елена", "Жанна", "Захар", "Игорь"}
	//rand.Seed(time.Now().UnixNano())

	fmt.Println(groupDistribution(students))
	fmt.Println(students) /// Ну очень интересный результат))

	fmt.Println()
}

func statisticOrders(inSliceMap []map[string]string) map[string]int {
	/*  Задача 1: Обработка заказов в интернет-магазине
	У вас есть слайс заказов, каждый из которых представлен картой (map), содержащей наименование товара
	и его статус ("новый", "в обработке", "отменён", "доставлен").
	Напишите код, который: 	Считает количество заказов в каждом статусе. Выводит итоговую статистику.
	Ожидаемый вывод:  Новые заказы: 2  В обработке: 1 Отменённые: 1  Доставленные: 1
	*/
	resultMap := map[string]int{}
	for _, order := range inSliceMap {
		switch order["status"] {
		case "новый":
			resultMap["Новые заказы"]++
		case "в обработке":
			resultMap["В обработке"]++
		case "отменён":
			resultMap["Отменённые"]++
		case "доставлен":
			resultMap["Доставленные"]++
		}
	}
	return resultMap
}

func statisticTemperature(temps [7]float64) string {
	/*  Задача 2: Работа с массивом температур
		У вас есть массив температур за неделю:   temps := [7]float64{12.5, 14.3, 10.8, 9.5, 8.1, 11.7, 13.2}
		Реализуйте программу, которая:   	Определяет минимальную, максимальную и среднюю температуру.
											Выводит разницу между самой высокой и самой низкой температурой.
	Ожидаемый вывод:
			Максимальная температура: 14.3°C
			Минимальная температура: 8.1°C
			Средняя температура: 11.4°C
			Разница температур: 6.2°C
	*/
	max, min := temps[0], temps[0]
	var sum, average float64
	for _, temp := range temps {
		max = math.Max(max, temp)
		min = math.Min(min, temp)
		sum += temp
	}
	average = sum / 7
	//result += fmt.Sprintf("max - %.1f\nmin - %.1f", max, min)
	return fmt.Sprintf("Максимальная температура: %.1f°C\n"+
		"Минимальная температура: %.1f°C\n"+
		"Средняя температура: %.1f°C\n"+
		"Разница температур: %.1f°C", max, min, average, max-min)
}

func calculatorCommand(teams map[string]int) string {
	/*  Задача 3: Калькулятор команд в турнире
	У вас есть список команд и их очков в турнире:
	Напишите программу, которая: 	Определяет команду с наибольшим количеством очков.
									Определяет команду с наименьшим количеством очков.
									Выводит команды, набравшие выше среднего количества очков.
	Ожидаемый вывод:		Лидер турнира: Акулы (18 очков)
							Аутсайдер: Львы (10 очков)
							Команды выше среднего: Тигры, Орлы, Волки, Акулы
	*/
	slicepoints := []float64{}
	var leader, outsider, aboveAverage string
	aboveAverageSlice := []string{}
	for nameCommand, _ := range teams {
		slicepoints = append(slicepoints, float64(teams[nameCommand]))
	}

	max, min := slicepoints[0], slicepoints[0]
	var sum, count, average float64
	for _, point := range slicepoints {
		max = math.Max(max, point)
		min = math.Min(min, point)
		sum += float64(point)
		count++
	}
	average = sum / count

	for nameCommand, _ := range teams {
		if float64(teams[nameCommand]) == max {
			leader = fmt.Sprintf("Лидер турнира: %s (%d очков)\n", nameCommand, teams[nameCommand])
		}
		if float64(teams[nameCommand]) == min {
			outsider = fmt.Sprintf("Аутсайдер: %s (%d очков)\n", nameCommand, teams[nameCommand])
		}
		if float64(teams[nameCommand]) > average {
			aboveAverageSlice = append(aboveAverageSlice, nameCommand)
		}
	}
	aboveAverage = "Команды выше среднего: " + strings.Join(aboveAverageSlice, ", ")
	//return fmt.Sprintf("%s\n%s\n%s", leader, outsider, aboveAverage)
	return leader + outsider + aboveAverage
}

func groupDistribution(students []string) string {
	/*   Задача 4: Распределение студентов по группам
	В университете есть список студентов, которых нужно случайным образом распределить по трём группам.
	students := []string{"Алиса", "Борис", "Виктор", "Глеб", "Дмитрий", "Елена", "Жанна", "Захар", "Игорь"}
	Реализуйте программу, которая:  	Использует псевдослучайное распределение.
										Разбивает студентов по трём группам.
										Выводит списки студентов в каждой группе.
	Ожидаемый вывод (пример):		Группа 1: Алиса, Глеб, Игорь
									Группа 2: Борис, Дмитрий, Захар
									Группа 3: Виктор, Елена, Жанна
	*/

	//rand.Seed(time.Now().UnixNano())
	var groupOne, groupTwo, groupThree []string
	maxLen := len(students)
	for i := 0; i < maxLen; i++ {
		//fmt.Println(i, len(students))
		number := rand.Intn(len(students))
		//fmt.Println(len(groupOne))
		switch {
		case len(groupOne) < 3:
			groupOne = append(groupOne, students[number])
			students[number] = students[len(students)-1]
			students = students[:len(students)-1]
		case len(groupTwo) < 3:
			groupTwo = append(groupTwo, students[number])
			students[number] = students[len(students)-1]
			students = students[:len(students)-1]
		case len(groupThree) < 3:
			groupThree = append(groupThree, students[number])
			students[number] = students[len(students)-1]
			students = students[:len(students)-1]
		}

		//fmt.Println(students)

	}
	return fmt.Sprintf("Группа 1: %s\nГруппа 2: %s\nГруппа 3: %s",
		strings.Join(groupOne, ", "),
		strings.Join(groupTwo, ", "),
		strings.Join(groupThree, ", "))
}
