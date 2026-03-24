package main

import (
	"Kurs_4/account"
	"fmt"
)

func main() {
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleAccount()

		default:
			return
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант: ")
	fmt.Println("1. Создать аккаунт ")
	fmt.Println("2. Найти аккаунт ")
	fmt.Println("3. Удалить аккаунт ")
	fmt.Println("4. Выйти")
	fmt.Scan(&variant)
	return variant
}

func findAccount() {

}

func deleAccount() {

}

func createAccount() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите url: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("некорректный urLL или логин")
		return
	}
	vault := account.NewVault()
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)
	return value
}
