package main

import "fmt"

func main() {
	// Task 1
	fmt.Println("Task 1")
	// true
	fmt.Println(5 == 5)

	// true
	fmt.Println(10 != 3)

	// false
	fmt.Println(7 > 12)

	// true
	fmt.Println(15 < 20)

	// true
	fmt.Println(8 >= 8)

	// false
	fmt.Println(6 <= 4)

	// false
	fmt.Println((10 > 5) && (3 < 1))

	// true
	fmt.Println((10 > 5) || (3 < 1))

	// false
	fmt.Println(!(5 == 5))

	// true
	fmt.Println(!(7 < 3))

	// false
	fmt.Println(true && false)

	// false
	fmt.Println(false || false)

	// true
	fmt.Println(true || false)

	// true
	fmt.Println((4+6 == 10) && (9 > 2))

	// true
	fmt.Println((12/3 == 4) || (8 < 5))

	// Task 2
	fmt.Println("Task 2")
	age := 20
	hasTicket := true

	canEnter := age >= 18 && hasTicket

	fmt.Println(canEnter)

	// Task 3
	fmt.Println("Task 3")
	isLoggedIn := true
	isAdmin := false

	hasAccess := (isLoggedIn && isAdmin) || (isLoggedIn && !isAdmin)

	fmt.Println(hasAccess)
}
