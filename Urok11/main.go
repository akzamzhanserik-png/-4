package main

import "fmt"

func main() {
	// Task 1
	for i := 1; i <= 20; i++ {
		fmt.Println(i)
	}

	// Task 2
	sum := 0

	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Task 3
	number := 5

	for i := 1; i <= 10; i++ {
		fmt.Println(number * i)
	}

	// Task 4
	var n int
	fmt.Print("Введите n:")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	// Task 5
	var numberr int
	fmt.Print("Введите nuberr:")
	count := 0
	fmt.Scan(&numberr)

	if numberr == 0 {
		count = 1
	} else {
		for numberr != 0 {
			numberr = numberr / 10
			count++
		}
	}

	fmt.Println("Количество цифр:", count)

	// Task 6
	text := "Developer"

	for i := 0; i < len(text); i++ {
		fmt.Println(string(text[i]))
	}
	str := "Developer"
	for _, value := range str {
		fmt.Printf("Value: %c\n", value)
	}

	// Text 7
	balance := 3000
	var choice int

	for {
		fmt.Println("1 - Показать баланс")
		fmt.Println("2 - Пополнить баланс (+500)")
		fmt.Println("3 - Снять деньги (-200)")
		fmt.Println("0 - Выход")

		fmt.Print("Введите число: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Текущий баланс:", balance)
		} else if choice == 2 {
			balance += 500
			fmt.Println("Баланс пополнен. Новый баланс:", balance)
		} else if choice == 3 {
			balance -= 200
			fmt.Println("Деньги сняты. Новый баланс:", balance)
		} else if choice == 0 {
			fmt.Println("Выход из программы")
			break
		} else {
			fmt.Println("Неверная команда")
		}
	}
}
