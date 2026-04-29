package main

import (
	"fmt"
)

func main() {
	// Константы
	const BaseTariff float64 = 0.45
	const HighLoadTax float64 = 0.15
	const NightDiscount float64 = 0.30

	for {
		var name string

		// Ввод названия
		fmt.Print("Введите название прибора (или 'done'): ")
		fmt.Scan(&name)

		if name == "done" {
			fmt.Println("Расчет завершен.")
			break
		}

		var power float64
		var hours float64
		var nightMode bool

		// Ввод остальных данных
		fmt.Print("Мощность (Вт): ")
		fmt.Scan(&power)

		fmt.Print("Время работы (ч): ")
		fmt.Scan(&hours)

		fmt.Print("Ночной режим (true/false): ")
		fmt.Scan(&nightMode)

		// Расчет расхода
		consumption := (power * hours) / 1000.0

		// Базовая стоимость
		cost := consumption * BaseTariff

		// Скидка за ночной режим
		if nightMode {
			cost = cost * (1 - NightDiscount)
		}

		// Налог при высоком потреблении
		if consumption > 10 {
			cost = cost * (1 + HighLoadTax)
		}

		// Классификация (switch)
		var category string
		switch {
		case power < 100:
			category = "Экономный"
		case power >= 100 && power <= 1000:
			category = "Стандартный"
		default:
			category = "Мощный"
		}

		// Вывод отчета
		fmt.Println("\n--- Отчет по прибору ---")
		fmt.Printf("Прибор: %s [Категория: %s]\n", name, category)
		fmt.Printf("Расход: %.2f кВт·ч\n", consumption)
		fmt.Printf("Итоговая стоимость: %.2f\n", cost)
		fmt.Println("------------------------")
	}
}
