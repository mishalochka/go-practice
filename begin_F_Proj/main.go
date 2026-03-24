package main

import (
	"fmt"
	"log"
)

func main() {
	for {
		GetInput()
	}

}

func GetInput() (string, error) {
	fmt.Println("Введите текст: ")
	var str string

	if _, err := fmt.Scan(&str); err != nil {
		log.Fatalf("divide error: %s", err)
	}
	return str, nil
}

func CountCharacters(text string) (letters, digits, spaces, punctuation int) {
	return letters, digits, spaces, punctuation
}

func DisplayResults(letters, digits, spaces, punctuation int) {
	fmt.Printf("Количество букв:%d", letters)
	fmt.Printf("Количество цифр:%d", digits)
	fmt.Printf("Количество пробелов:%d", spaces)
	fmt.Printf("Количество букв:%d", punctuation)
}
