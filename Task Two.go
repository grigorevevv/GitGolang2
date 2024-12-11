package Task_Two

/*
2. Функция с несколькими типами данных
Создайте функцию Calculate, которая принимает два числа типа float64 и строку operation.
Функция выполняет сложение, вычитание, умножение или деление в зависимости от переданного
значения operation. При делении на ноль должна возвращаться ошибка.
*/

import "fmt"

func main() {
	var x, y float64
	var operStr string

	fmt.Print("Введите значение для переменной x: ")
	fmt.Scan(&x)
	fmt.Print("Введите значение для переменной y: ")
	fmt.Scan(&y)

	/*reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите текст: ")
	operStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}*/
	//fmt.Println("Вы ввели:", operStr)
	operStr = "сложение"
	fmt.Println(Calculate(x, y, operStr))
	operStr = "вычитание"
	fmt.Println(Calculate(x, y, operStr))
	operStr = "умножение"
	fmt.Println(Calculate(x, y, operStr))
	operStr = "деление"
	fmt.Println(Calculate(x, y, operStr))
	//fmt.Println("Значение переменной b - " + strconv.Itoa(b))
	//fmt.Println("Значение переменной c - " + strconv.Itoa(c))
}

func Calculate(x, y float64, operation string) float64 {

	if y == 0 {
		return 0
		//Println("Делить на ноль нельзя!"))
	}

	switch operation {
	case "сложение":
		return x + y
	case "вычитание":
		return x - y
	case "умножение":
		return x * y
	case "деление":
		return x / y
	default:
		return 0
	}

}
