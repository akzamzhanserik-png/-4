package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Task 1
	toolUsage := map[string]int{
		"Go":     3,
		"VSCode": 5,
		"Git":    2,
	}
	for tool, count := range toolUsage {
		fmt.Println(tool, count)
	}

	// Task 2
	buildStatus := make(map[string]bool)
	buildStatus["build"] = true
	buildStatus["run"] = false

	if buildStatus["build"] {
		fmt.Println("Сборка прошла успешно")
	}

	// Task 3
	var name string
	fmt.Print("Введите ваше имя: ")
	fmt.Scan(&name)

	userInfo := make(map[string]string)
	userInfo["name"] = name
	userInfo["isLoggedIn"] = "true"

	fmt.Printf("Пользователь %s вошёл в систему\n", userInfo["name"])

	// Task 4
	cpuLoad := make(map[int]int)
	cpuLoad[1] = 40
	cpuLoad[2] = 65
	cpuLoad[3] = 30

	maxCore := 0
	maxLoad := 0

	for key, value := range cpuLoad {
		if value > maxLoad {
			maxLoad = value
			maxCore = key
		}
	}

	fmt.Println("Ядро с максимальной загрузкой:", maxCore)
	fmt.Println("Загрузка:", maxLoad)

	// Task 5
	examResults := map[string]int{
		"Aruzhan": 85,
		"Dias":    92,
		"Alina":   78,
	}
	for namem, score := range examResults {
		if score >= 80 {
			fmt.Printf("clear%s сдал экзамен\n", namem)
		} else {
			fmt.Printf("%s не сдал экзамен\n", namem)
		}
	}

	// Task 6
	words := []string{"go", "is", "fast"}

	wordLength := map[string]int{}

	for _, word := range words {
		wordLength[word] = len(word)
	}

	for word, length := range wordLength {
		fmt.Printf("'%s' - длина: %d\n", word, length)
	}

	// Task 7
	menu := map[string]int{
		"Burger": 2500, "Pizza": 3200, "Tea": 500,
	}
	var eda string
	fmt.Print("Название блюда : ")
	fmt.Scan(&eda)

	if price, ok := menu[eda]; ok {
		fmt.Printf("%s стоит %d тенге\n", eda, price)
	} else {
		fmt.Println("Блюдо не найдено")
	}

	// Task 8
	loginAttempts := map[string]int{
		"admin": 2,
		"guest": 0,
	}

	loginAttempts["admin"]++

	if loginAttempts["admin"] > 2 {
		fmt.Println("Доступ заблокирован")
	} else {
		fmt.Println("Доступ разрешён")
	}

	// Task 9
	sales := [2][3]int{
		{100, 150, 200}, // магазин 1
		{80, 120, 160},  // магазин 2
	}

	total := map[int]int{}

	for i, store := range sales {
		for _, daySale := range store {
			total[i+1] += daySale
		}
	}

	for store, sum := range total {
		fmt.Printf("Магазин %d — общие продажи: %d\n", store, sum)
	}

	// Task 10
	numbers := []int{4, 7, 2, 9, 5}
	numberStatus := map[int]string{}

	for _, num := range numbers {
		if num%2 == 0 {
			numberStatus[num] = "even"
		} else {
			numberStatus[num] = "odd"
		}
	}

	for num, status := range numberStatus {
		fmt.Printf("%d — %s\n", num, status)
	}

	// Task 11
	defaultConfig := map[string]string{
		"host": "localhost",
		"port": "8080",
		"mode": "production",
	}

	currentConfig := map[string]string{
		"host": "localhost",
		"port": "8080",
		"mode": "production",
	}

	if reflect.DeepEqual(defaultConfig, currentConfig) {
		fmt.Println("Конфигурации совпадают")
	} else {
		fmt.Println("Конфигурации отличаются")
	}

	currentConfig["mode"] = "debug"

	if reflect.DeepEqual(defaultConfig, currentConfig) {
		fmt.Println("Конфигурации совпадают")
	} else {
		fmt.Println("Конфигурации отличаются")
	}

}
