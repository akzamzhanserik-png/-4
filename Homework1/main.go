package main

import (
	"fmt"
)

func main() {
	// Константы
	const BaseRate float64 = 5.50
	const TaxRate float64 = 0.12
	const DistanceRate float64 = 2.0
	const FragileFee float64 = 0.2

	// Переменные для ввода
	var name string
	var weight float64
	var distance float64
	var fragileCount int

	// Ввод данных
	fmt.Print("Введите имя: ")
	fmt.Scan(&name)

	fmt.Print("Вес груза (кг): ")
	fmt.Scan(&weight)

	fmt.Print("Дистанция (км): ")
	fmt.Scan(&distance)

	fmt.Print("Количество хрупких упаковок: ")
	fmt.Scan(&fragileCount)

	// Расчеты
	baseCost := (weight*BaseRate)*(1+FragileFee*float64(fragileCount)) + (distance * DistanceRate)
	totalCost := baseCost * (1 + TaxRate)

	// Вывод отчета
	fmt.Println("\n--- Отчет о доставке ---")
	fmt.Printf("Отправитель: %s\n", name)
	fmt.Printf("Итоговая стоимость: %.2f\n", totalCost)
}
