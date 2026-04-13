package main

import "fmt"

func main() {

	// ===================== Task 1 =====================
	var numbers [3][5]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&numbers[i][j])
		}
	}

	max_i, max_j := 0, 0
	max_value := numbers[0][0]

	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			if numbers[i][j] > max_value {
				max_value = numbers[i][j]
				max_i = i
				max_j = j
			}
		}
	}

	fmt.Println(max_i, max_j)

	// ===================== Task 2 =====================
	var snowflake [11][11]string

	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			snowflake[i][j] = "."
		}
	}

	mid := 11 / 2

	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if i == mid || j == mid || i == j || i+j == 10 {
				snowflake[i][j] = "*"
			}
		}
	}

	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			fmt.Print(snowflake[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println()

	// ===================== Task 3 =====================
	var board [8][8]string

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				board[i][j] = "."
			} else {
				board[i][j] = "*"
			}
		}
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}

	// ===================== Task 4 =====================
	var arr [4][4]int

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	var row1, row2 int
	fmt.Scan(&row1, &row2)

	for col := 0; col < 4; col++ {
		arr[row1][col], arr[row2][col] = arr[row2][col], arr[row1][col]
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	// ===================== Task 5 =====================
	var arr2 [4][4]int

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Scan(&arr2[i][j])
		}
	}

	var col1, col2 int
	fmt.Scan(&col1, &col2)

	for row := 0; row < 4; row++ {
		arr2[row][col1], arr2[row][col2] = arr2[row][col2], arr2[row][col1]
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(arr2[i][j], " ")
		}
		fmt.Println()
	}
}
