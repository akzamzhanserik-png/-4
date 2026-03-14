package main

import "fmt"

func main() {
	// Task 1
	temperature := 15

	if temperature < 0 {
		fmt.Println("Холодно")
	} else if temperature >= 0 && temperature <= 20 {
		fmt.Println("Тепло")
	} else {
		fmt.Println("Жарко")
	}

	// Task 2
	score := 85

	if score >= 90 {
		fmt.Println("Отлично")
	} else if score >= 70 && score <= 89 {
		fmt.Println("Хорошо")
	} else if score >= 50 && score <= 69 {
		fmt.Println("Удовлетворительно")
	} else {
		fmt.Println("Не сдал")
	}

	// Task 3
	hour := 20

	switch {
	case hour >= 0 && hour <= 5:
		fmt.Println("Ночь")
	case hour >= 6 && hour <= 11:
		fmt.Println("Утро")
	case hour >= 12 && hour <= 17:
		fmt.Println("День")
	case hour >= 18 && hour <= 23:
		fmt.Println("Вечер")
	}

	// Task 4
	var number int
	fmt.Print("Введите любое число: ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Println("Четное число")
	} else {
		fmt.Println("Нечётное число")
	}

	// Task 5
	day := "Tuesday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Будний день")
	case "Saturday", "Sunday":
		fmt.Println("Выходной")
	default:
		fmt.Println("Некорректный день")
	}
	// Task 6
	balance := -50

	if balance >= 0 {
		fmt.Println("Баланс положительный")
	} else {
		fmt.Println("Баланс отрицательный")
	}
	//Task 7
	var age int
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)
	if age < 13 {
		fmt.Println("Ребёнок")
	} else if age >= 13 && age <= 17 {
		fmt.Println("Подросток")
	} else {
		fmt.Println("Взрослый")
	}
	// Task 8
	var command string
	fmt.Print("Введите команду (start, stop, restart): ")
	fmt.Scan(&command)

	switch command {
	case "start":
		fmt.Println("Команда старт выполнена")
	case "stop":
		fmt.Println("Команда стоп выполнена")
	case "restart":
		fmt.Println("Команда перезапуск выполнена")
	default:
		fmt.Println("Неизвестная команда")
	}

	// Task 9
	grade := 2
	switch grade {
	case 5:
		fmt.Println("A")
	case 4:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	case 2:
		fmt.Println("D")
	case 1:
		fmt.Println("F")
	default:
		fmt.Println("Некорректная оценка")

	}
}
