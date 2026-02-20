package main

import (
	"fmt"
)

func main() {
	//Задание 1
	age := 18
	fmt.Println(age)

	age = 19
	fmt.Println(age)

	//Задание 2
	height := 180
	fmt.Println(height)

	height_in_meters := float64(height) / 100
	fmt.Println(height_in_meters)

	//Задание 3
	isStudent := true
	fmt.Println(isStudent)

	//Задание 4
	temperature := 25
	fmt.Println("Температура на улице:", temperature, "°C")

	if temperature >= 20 {
		fmt.Println("Погода теплая")
	} else {
		fmt.Println("Погода холодная")
	}

	//Задание 5
	favoriteQuote := "Учение — свет, а неучение — тьма."
	fmt.Println("Моя любимая цитата:", favoriteQuote)

	//Задание 6
	PI := 3.14
	fmt.Println(PI)
	// Попытка присвоить строку приведет к ошибке
	// PI := "3.1415" // ошибка: нельзя присвоить строку переменной, которая ранее была числом
}
