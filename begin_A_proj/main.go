package main

import "fmt"

func main() {

}

// Сохраняем состояние пирата между вызовами функции
var (
	plate  = 1     // Текущая плита
	traps  = 0     // Количество ловушек, в которые попал пират
	isDead = false // Статус жизни пирата
)

func movePirate(isTrap bool) {

	if isDead {
		return
	}

	fmt.Printf("Пират переместился на плиту %d\n", plate)

	if isTrap {
		traps++
		if traps == 3 {
			fmt.Println("Пират убит")
			isDead = true
			return
		}
		fmt.Println("Пират ранен")
	}

	if plate == 10 {
		fmt.Println("Пират преодолел все ловушки")
	}
	plate++
}
