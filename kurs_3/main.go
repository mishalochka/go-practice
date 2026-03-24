package main

import (
	"fmt"
)

// type stringMap = map[int]string

// func main() {
// 	bookMarks := map[int]string{}
// Menu:
// 	for {
// 		variant := getMenu()
// 		switch variant {
// 		case 1:
// 			printBookmarks(bookMarks)
// 		case 2:
// 			addBookMarks(bookMarks)
// 		case 3:
// 			deleteBookMarks(bookMarks)
// 		default:
// 			break Menu
// 		}
// 	}
// }

// func getMenu() int {
// 	var variant int
// 	fmt.Println("Выберите вариант")
// 	fmt.Println("1. Посмотреть закладки")
// 	fmt.Println("2. Добавить закладку")
// 	fmt.Println("3. Удалить закладку")
// 	fmt.Println("4. Выход")
// 	fmt.Scan(&variant)
// 	return variant
// }

// func printBookmarks(bookMarks map[int]string) {
// 	if len(bookMarks) == 0 {
// 		fmt.Println("пока нет закладок")
// 	}
// 	for key, value := range bookMarks {
// 		fmt.Println(key, ": ", value)
// 	}
// }

// func addBookMarks(bookMarks map[int]string) {
// 	var newBookmarkKey int
// 	var newBookMarkValue string
// 	fmt.Println("Введите число: ")
// 	fmt.Scan(&newBookmarkKey)
// 	fmt.Println("Введите закладку: ")
// 	fmt.Scan(&newBookMarkValue)
// 	bookMarks[newBookmarkKey] = newBookMarkValue

// }

// func deleteBookMarks(bookMarks map[int]string) {
// 	var del1 int
// 	fmt.Print("Введите число вашей закладки: ")
// 	fmt.Scan(&del1)
// 	delete(bookMarks, del1)

// }

// Создать приложение, которое сначала выдаёт меню:
// -1. Посмотреть закладки
// -2. Добавить закладку
// -3. Удалить закладку
// -4. Выход
// При 1 - выводит закладки
// При 2 - 2 поля ввода названия и адреса и после, добавлениЕ
// При 3 - ввод названия и удаление по нему
// При 4 - Завершение

type hashTable = map[int]string

func add(a []string) {
	a[0] = "2"
	fmt.Println(a)
}

func main() {
	a := []string{"1"}
	add(a)
	fmt.Println(a)

	hashmap := hashTable{}
	var variant int
	fmt.Println("Болевуха для закладок")
	helloworld()
	for {
		fmt.Scan(&variant)
		switch variant {
		case 1:
			seeBookmarks(hashmap)
		case 2:
			addBookMarks(hashmap)
		case 3:
			deleteBookMarks(hashmap)
		case 4:
			return
		}
	}
}

func seeBookmarks(hash hashTable) {
	for key, value := range hash {
		fmt.Println(key, ":", value)
	}
}

func addBookMarks(hash hashTable) {
	fmt.Print("Введите ключ: ")
	var key int
	fmt.Scan(&key)
	fmt.Print("Введите закладку: ")
	var value string
	fmt.Scan(&value)
	hash[key] = value
}

func deleteBookMarks(hash hashTable) {
	fmt.Println("Введите ключ, который хотите удалить: ")
	for key, value := range hash {
		fmt.Println(key, ":", value)
	}
	var key int
	fmt.Scan(&key)
	delete(hash, key)
}

func helloworld() {
	fmt.Println("1 - Посмотреть закладки")
	fmt.Println("2 - Добавить закладку")
	fmt.Println("3 - Удалить закладку")
	fmt.Println("4 - Выход")
}
