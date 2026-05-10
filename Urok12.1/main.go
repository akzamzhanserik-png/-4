package main

import (
	"fmt"
	"strings"
)

func main() {
	// Task 1
	fmt.Println("Task-1")
	text := "REAL MADRID"
	const author = "Zhanserik"
	fmt.Printf("Автор %s написал: %s\n", author, text)
	fmt.Printf("Длина строки text: %d\n", len(text))

	// Task 2
	fmt.Println("\nTask-2")
	message := "\"Учись так, словно тебе предстоит жить вечно\""
	if message == "" {
		fmt.Println("Строка пуста")
	} else {
		fmt.Println("Строка не пустая")
	}
	fmt.Printf("Длина строки text: %d\n", len(message))

	// Task 3
	fmt.Println("\nTask-3")
	var input string

	fmt.Print("Введите строку: ")
	fmt.Scan(&input)

	length := len([]rune(input))

	if length < 5 {
		fmt.Println("Слишком короткое слово")
	} else if length <= 10 {
		fmt.Println("Нормальная длина")
	} else {
		fmt.Println("Слишком длинное слово")
	}

	// Task 4
	fmt.Println("\nTask-4")
	word := "Golang"

	runes := []rune(word)

	first := runes[0]
	last := runes[len(runes)-1]

	fmt.Printf("Первый символ: %c\n", first)
	fmt.Printf("Последний символ: %c\n", last)

	fmt.Println("\nКаждый символ строки:")
	for _, ch := range word {
		fmt.Printf("%c\n", ch)
	}

	fmt.Printf("\nДлина строки: %d\n", len(runes))

	// Task 5

	fmt.Println("\nTask-5")
	sentence := "Аутентичность awakens пробуждает и продвигает всё вокруг Атланты"

	runess := []rune(sentence)
	lengthh := len(runess)

	countA := 0
	for _, ch := range runess {
		if ch == 'а' || ch == 'А' {
			countA++
		}
	}

	fmt.Printf("В строке %d символов и %d букв а\n", lengthh, countA)

	// Task 6
	fmt.Println("\nTask-6")
	var inputt string

	fmt.Print("Введите строку: ")
	fmt.Scan(&inputt)

	symbols := []rune(inputt)
	reversed := ""

	for i := len(symbols) - 1; i >= 0; i-- {
		reversed += string(symbols[i])
	}

	fmt.Printf("Исходная строка: %s\n", inputt)
	fmt.Printf("Перевёрнутая строка: %s\n", reversed)

	if inputt == reversed {
		fmt.Println("Это палиндром")
	} else {
		fmt.Println("Это не палиндром")
	}

	// Task 7
	fmt.Println("\nTask-7")

	var textt string

	// Ввод строки
	fmt.Print("Введите строку: ")
	fmt.Scan(&textt)

	// Количество символов
	lengthhh := len(textt)

	// Количество слов
	words := strings.Fields(textt)
	wordCount := len(words)

	// Подсчёт гласных
	vowels := "aeiouyAEIOUY"
	vowelCount := 0

	for _, char := range textt {
		if strings.ContainsRune(vowels, char) {
			vowelCount++
		}
	}

	// Вывод результата
	fmt.Printf("Длина строки: %d, слов: %d, гласных: %d\n",
		lengthhh, wordCount, vowelCount)

}
