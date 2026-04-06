package main

import "fmt"

func main() {
	//Task 1
	runningExercises := [...]string{"Спринт 100м", "Бег на длинную дистанцию"}
	walkingExercises := [...]string{"Прогулка в парке", "Скандинавская ходьба"}
	fmt.Println(len(runningExercises))
	fmt.Println(len(walkingExercises))

	// Task 2
	subjectsList := [3]string{"Физика", "Химия", "География"}
	first := subjectsList[0]
	last := subjectsList[len(subjectsList)-1]
	fmt.Println(first)
	fmt.Println(last)
	subjectsList[1] = "Биология"
	fmt.Println(subjectsList)

	// Task 3
	person := [3]string{"Tom", "35", "New York"}
	name := person[0]
	age := person[1]
	home := person[2]

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(home)

	// Task 4
	numbersList := [5]int{1, 2, 3, 4, 5}

	isNumberPresent := false

	for _, number := range numbersList {
		if number == 3 {
			isNumberPresent = true
			break
		}
	}

	if isNumberPresent {
		fmt.Println("Число 3 есть в массиве")
	} else {
		fmt.Println("Число 3 отсутствует в массиве")
	}

	// Task 5
	friendsList := [5]string{"Aibek", "Daulet", "Bekbolat", "Nursultan", "Arman"}
	found := false
	for _, name := range friendsList {
		if name == "Bekbolat" {
			found = true
			break
		}
	}
	if found {
		fmt.Println("Мне очень повезло")
	} else {
		fmt.Println("Мне не повезло")
	}

	// Task 6
	numbers1 := [3]int{1, 2, 3}
	numbers2 := [3]int{1, 2, 4}

	if numbers1 == numbers2 {
		fmt.Println("списки одинаковы")
	} else {
		fmt.Println("они не одинаковы")
	}

	// Task 7
	myWishList := [3]string{"ноутбук", "наушники", "книга"}
	friendsWishList := [3]string{"телефон", "часы", "рюкзак"}
	var registrationList []string

	for _, item := range myWishList {
		registrationList = append(registrationList, item)
	}
	for _, item := range friendsWishList {
		registrationList = append(registrationList, item)
	}
	fmt.Println(registrationList)

	// Task 8
	toyList := [3]string{"Car", "Doll", "Ball"}
	testToyList := toyList
	testToyList[1] = "Boat"
	fmt.Println(toyList)
	fmt.Println(testToyList)

}
