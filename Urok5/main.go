package main

import "fmt"

func main() {
	// ===== Задание 1 =====
	// Количество лет обучения в школе
	schooling := 11
	fmt.Println("Задание 1:")
	fmt.Println("Лет обучения в школе:", schooling)

	// Добавился ещё один год обучения
	schooling = 12
	fmt.Println("После дополнительного года:", schooling)
	fmt.Println()

	// ===== Задание 2 =====
	// Работа со строковой переменной
	name := "Vladislav"
	fmt.Println("Задание 2:")
	fmt.Println("Имя:", name)

	// Изменяем имя
	name = "Zhanserik"
	fmt.Println("Новое имя:", name)
	fmt.Println()

	// ===== Задание 3 =====
	// Количество шагов за день
	steps := 0
	fmt.Println("Задание 3:")
	fmt.Println("Шаги утром:", steps)

	// К обеду шаги увеличились
	steps = 2000
	fmt.Println("Шаги к обеду:", steps)
	fmt.Println("Хорошая работа! Вы уже на пути к своей ежедневной цели")
	fmt.Println()

	// ===== Задание 4 =====
	// Большое число
	largeNumber := 6000000
	fmt.Println("Задание 4:")
	fmt.Println("Большое число:", largeNumber)
	fmt.Println()

	// ===== Задание 5 =====
	// Константа — длительность перемены
	const breakTime = 15
	fmt.Println("Задание 5:")
	fmt.Println("Длительность перемены:", breakTime)

	// breakTime = 20
	// Ошибка: константы в Go нельзя изменять после объявления

}
