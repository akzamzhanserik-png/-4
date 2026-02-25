package main

import (
	"fmt"
	"math"
)

func main() {
	// Задание 1
	bannerWidth := 12
	bannerHeight := 8
	bannerArea := bannerWidth * bannerHeight
	fmt.Println(bannerArea)

	halfBannerArea := bannerArea / 2
	fmt.Println(halfBannerArea)

	bannerBorderLength := (bannerWidth + bannerHeight) * 2
	fmt.Println(bannerBorderLength)

	// Задание 2
	boxCount := 29
	leftoverBoxes := boxCount % 5
	fmt.Println(leftoverBoxes)

	// Задание 3
	tempMorning := 18
	tempAfternoon := 27
	tempEvening := 21

	totalTemp := tempMorning + tempAfternoon + tempEvening
	averageTemp := totalTemp / 3

	fmt.Println(averageTemp)

	// Задание 4
	knownWords := 47
	wordsGoal := 120

	progressPercent := float64(knownWords) * 100 / float64(wordsGoal)

	fmt.Println(progressPercent, "%")

	// Задание 5
	coins := 0
	fmt.Println(coins)

	coins += 500
	fmt.Println(coins)

	coins += 1200
	fmt.Println(coins)

	coins /= 2
	fmt.Println(coins)

	coins *= 2
	fmt.Println(coins)

	coins -= 300
	fmt.Println(coins)

	// Задание 6
	participants := 42
	groupCount := 8

	participantsPerGroup := participants / groupCount
	fmt.Println(participantsPerGroup)

	// Задание 7
	fmt.Println(20 - 4*3)
	fmt.Println((20 - 4) * 3)

	/*
		Объяснение:

		В первом выражении:
			20 - 4 * 3
		Сначала выполняется умножение (4 * 3 = 12), затем вычитание:
			20 - 12 = 8

		Во втором выражении:
			(20 - 4) * 3
		Сначала выполняется выражение в скобках (20 - 4 = 16), затем умножение:
			16 * 3 = 48

		Разница в результатах возникает из-за **приоритета операций**:
			* Умножение (*) выполняется **до** вычитания (-)
			* Скобки изменяют порядок выполнения операций
	*/

	// Задание 8
	squareValue := 81.0
	sqrtResult := math.Sqrt(squareValue)
	fmt.Println(sqrtResult)

	multiplier := 5.0
	exponent := 2.0
	powResult := math.Pow(multiplier, exponent)
	fmt.Println(powResult)
}
