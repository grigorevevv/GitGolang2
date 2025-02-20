package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	K1 = 0.035
	K2 = 0.029
)

var (
	Format     = "20060102 15:04:05" // формат даты и времени
	StepLength = 0.65                // длина шага в метрах
	Weight     = 75.0                // вес кг
	Height     = 1.75                // рост м
	Speed      = 1.39                // скорость м/с
)

/*
	 	parsePackage разбирает входящий пакет в параметре data.  Возвращаемые значения:
			t — дата и время, указанные в пакете
			steps — количество шагов
	 		ok — true, если время и шаги указаны корректно, и false — в противном случае
*/

func AcceptPackage(data string, storage []string) []string {
	/*	1. Используйте parsePackage для разбора пакета
				t, steps, ok := parsePackage(data)
		  		выведите сообщение в случае ошибки также проверьте количество шагов на равенство нулю
		2. Получите текущее UTC-время и сравните дни выведите сообщение, если день в пакете t.Day()
				не совпадает с текущим днём		*/

	//t, steps, ok := parsePackage(data)
	t, step, ok := parsePackage(data)
	if !ok {
		showMessage("ошибочный формат пакета")
		return storage
	}
	if step == 0 {
		return storage
	}

	now := time.Now().UTC()
	// выводим ошибку, если время в пакете больше текущего времени
	if t.After(now) {
		showMessage(`некорректное значение времени`)
		return storage
	}

	// проверки для непустого storage
	if len(storage) > 0 {
		/* 3. Достаточно сравнить первые len(Format) символов пакета с len(Format) символами последней записи storage
		если меньше или равно, то ошибка — некорректное значение времени     */

		dateLastPackage, err := time.Parse(Format, storage[len(storage)-1][:17])
		if err != nil {
			fmt.Println(err)
		}
		if dateLastPackage.After(t) || dateLastPackage == t {
			showMessage(`некорректное значение времени`)
			return storage
		}
		//fmt.Println(dateLastPackage)

		// смотрим, наступили ли новые сутки: YYYYMMDD — 8 символов
		if data[:8] != storage[len(storage)-1][:8] {
			// если наступили,то обнуляем слайс с накопленными данными
			storage = storage[:0]
		}
	}
	// остаётся совсем немного
	// 5. Добавить пакет в storage
	timeString := t.Format(Format)
	stepString := strconv.Itoa(step)

	storage = append(storage, timeString+","+stepString)

	// 6. Получить общее количество шагов

	sumSteps := float64(stepsDay(storage))
	//fmt.Println(sumSteps)

	// 7. Вычислить общее расстояние (в метрах)

	distance := sumSteps * StepLength
	//fmt.Println(distance)

	// 8. Получить потраченные килокалории
	calorie := calories(distance)
	//fmt.Println(caloie)

	// 9. Получить мотивирующий текст

	textMotiv := achievement(distance)
	//fmt.Println(textMotiv)

	// 10. Сформировать и вывести полный текст сообщения

	fmt.Printf("Время: %s.\nКоличество шагов за сегодня: %.f.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n%s\n\n",
		t.Format("15:04:05"), sumSteps, distance/1000, calorie, textMotiv)

	// 11. Вернуть storage
	return storage
}

func parsePackage(data string) (t time.Time, steps int, ok bool) {
	/* 	1. Разделите строку на две части по запятой в слайс ds
	2. Проверьте, чтобы ds состоял из двух элементов         */

	ds := strings.Split(data, ",")

	if len(ds) != 2 {
		//fmt.Println("ds состоял не из двух элементов")
		return
	}
	var err error
	// получаем время time.Time
	t, err = time.Parse(Format, ds[0])
	if err != nil {
		//fmt.Println("Ошибка времени", err)
		return
	}
	// получаем количество шагов
	steps, err = strconv.Atoi(ds[1])
	if err != nil || steps < 0 {
		//fmt.Println("Ошибка в значении количество шагов", err)
		return
	}
	// отмечаем, что данные успешно разобраны
	ok = true
	return
}

// showMessage выводит строку и добавляет два переноса строк
func showMessage(s string) {
	fmt.Printf("%s\n\n", s)
}

// stepsDay перебирает все записи слайса, подсчитывает и возвращает общее количество шагов
func stepsDay(storage []string) int {
	/*  тема оптимизации не затрагивается, поэтому можно использовать parsePackage для каждого элемента списка */
	var sumStep int
	for _, onePack := range storage {
		steps, err := strconv.Atoi(onePack[18:])
		if err != nil {
			fmt.Println(err)
		}
		sumStep += steps
	}
	return sumStep
}

// calories возвращает количество килокалорий, которые потрачены на  прохождение указанной дистанции (в метрах) со скоростью 5 км/ч
func calories(distance float64) float64 {
	/* Энергозатраты (ккал/мин) = 0,035 * m + (v*v/h) * 0,029 * m
	m — вес тела (кг);			h — рост (м); 			v — скорость ходьбы (м/с).      */

	energyCost := K1*Weight + (Speed*Speed/Height)*K2*Weight

	timeStep := (distance / Speed) / 60
	result := energyCost * timeStep

	return result
}

// achievement возвращает мотивирующее сообщение в зависимости от  пройденного расстояния в километрах
func achievement(distance float64) string {

	switch {
	case distance >= 6500:
		return "Отличный результат! Цель достигнута."
	case distance >= 3900:
		return "Неплохо! День был продуктивный."
	case distance >= 2000:
		return "Завтра наверстаем!"
	case distance < 2000:
		return "Лежать тоже полезно. Главное — участие, а не победа!"
	}
	return ""
}

func main() {
	/* Вы можете сразу проверить работу функции AcceptPackage на небольшом тесте. Если запустить программу после 05:00 UTC, то последнее
	сообщение должно быть таким:
			Время: 04:45:21.
			Количество шагов за сегодня: 16956.
			Дистанция составила 11.02 км.
			Вы сожгли 664.23 ккал.
			Отличный результат! Цель достигнута.
	*/

	now := time.Now().UTC()
	today := now.Format("20060102")

	//данные для самопроверки
	input := []string{
		"01:41:03,-100",
		",3456",
		"12:40:00, 3456 ",
		"something is wrong",
		"02:11:34,678",
		"02:11:34,792",
		"17:01:30,1078",
		"03:25:59,7830",
		"04:00:46,5325",
		"04:45:21,3123",
	}

	var storage []string
	//storage = AcceptPackage("20250219 00:11:33,100", storage)
	for _, v := range input {
		fmt.Println("Пакет:  ", v)
		storage = AcceptPackage(today+" "+v, storage)
	}

	//Для проверки содержимого пакета
	//for _, pack := range storage {
	//	fmt.Println(pack)
	//}
}
