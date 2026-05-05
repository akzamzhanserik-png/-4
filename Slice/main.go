package main

import "fmt"

func main() {
	// Task1
	var numbers []int

	for {
		var n int
		fmt.Print("Введите число: ")
		fmt.Scan(&n)

		if n == 0 {
			break
		}

		numbers = append(numbers, n)
	}
	fmt.Println(numbers)
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Println("Сумма:", sum)

	// Task 2
	var values []int

	for {
		var number int
		fmt.Print("Введите число: ")
		fmt.Scan(&number)

		if number == 0 {
			break
		}
		values = append(values, number)
	}

	var evens []int
	for _, v := range values {
		if v%2 == 0 {
			evens = append(evens, v)
		}
	}
	fmt.Println("values:", values)
	fmt.Println("evens: ", evens)

	// Task 3
	var data []int

	for {
		var nn int
		fmt.Print("Введите число: ")
		fmt.Scan(&nn)

		if nn == 0 {
			break
		}
		data = append(data, nn)
	}
	fmt.Println("До удаления: ", data)
	nn := 2
	data = append(data[:nn], data[nn+1:]...)
	fmt.Println(data)

	// Task 4
	var temps []int

	for {
		var gradus int
		fmt.Print("Введите градус:")
		fmt.Scan(&gradus)

		if gradus == 0 {
			break
		}
		temps = append(temps, gradus)
	}
	fmt.Println("ALL tempps : ", temps)

	if len(temps) == 0 {
		fmt.Println("Список пуст")
		return
	}
	minVal, maxVal := temps[0], temps[0]
	for _, j := range temps {
		if j < minVal {
			minVal = j
		}
		if j > maxVal {
			maxVal = j
		}
	}
	fmt.Printf("Min: %d, Max: %d\n", minVal, maxVal)

	// Task 5
	var words []string

	for {
		var word string
		fmt.Print("Введите слово : ")
		fmt.Scan(&word)

		if word == "stop" {
			break
		}
		words = append(words, word)
	}
	fmt.Println("До разворота: : ", words)

	var reversed []string
	for i := len(words) - 1; i >= 0; i-- {
		reversed = append(reversed, words[i])
	}
	fmt.Println("После разворота:", reversed)

	// Task 6
	var numss []int

	for {
		var n int
		fmt.Print("Введите число: ")
		fmt.Scan(&n)

		if n == 0 {
			break
		}
		numss = append(numss, n)
	}

	fmt.Println("Слайс:", numss)

	sorted := true
	for i := 1; i < len(numss); i++ {
		if numss[i] < numss[i-1] {
			sorted = false
			break
		}
	}

	fmt.Println("Отсортирован по возрастанию:", sorted)

	// Task 7

	var myWishList []string
	fmt.Println("Введите ваши желания (stop - завершить):")
	for {
		var wish string
		fmt.Print("Ваше желание: ")
		fmt.Scan(&wish)
		if wish == "stop" {
			break
		}
		myWishList = append(myWishList, wish)
	}

	var friendsWishList []string
	fmt.Println("Введите желания друга (stop - завершить):")
	for {
		var wish string
		fmt.Print("Желание друга: ")
		fmt.Scan(&wish)
		if wish == "stop" {
			break
		}
		friendsWishList = append(friendsWishList, wish)
	}

	var registrationList []string
	for _, v := range myWishList {
		registrationList = append(registrationList, v)
	}
	for _, v := range friendsWishList {
		registrationList = append(registrationList, v)
	}

	fmt.Println("Мои желания:      ", myWishList)
	fmt.Println("Желания друга:    ", friendsWishList)
	fmt.Println("Общий список:     ", registrationList)

	// Task 8

	toyList := []string{"Car", "Doll", "Ball"}

	testToyList := make([]string, len(toyList))
	copy(testToyList, toyList)

	testToyList[1] = "Boat"

	fmt.Println("toyList:     ", toyList)
	fmt.Println("testToyList: ", testToyList)

}
