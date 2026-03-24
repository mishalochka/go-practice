package main

import "fmt"

//В цикле спрашиваем число
//Добавлять его в массив
//Вывести массив
//Вывести его сумму

func main() {

	array := []int{}
	for {
		inputValue := scan()
		if inputValue == 0 {
			break
		}
		array = append(array, inputValue)
	}
	summ(array)
	fmt.Println(array)
}

func scan() int {
	var value int
	fmt.Print("Enter the number: ")
	fmt.Scan(&value)
	return value
}

func summ(array []int) {
	sum := 0
	for _, value := range array {
		sum += value
	}
	fmt.Println(sum)
}
