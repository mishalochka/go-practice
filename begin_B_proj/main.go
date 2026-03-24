package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	rollDice(8)
}

// Функция теперь возвращает две строки: глагол и существительное
func getRollWords(count int) (string, string) {
	mod100 := count % 100
	if mod100 >= 11 && mod100 <= 14 {
		return "потребовалось", "бросков"
	}

	mod10 := count % 10
	switch mod10 {
	case 1:
		return "потребовался", "бросок"
	case 2, 3, 4:
		return "потребовалось", "броска"
	default:
		return "потребовалось", "бросков"
	}
}

func rollDice(val int) {
	var sum int
	for i := 1; sum != val; i++ {
		dice1 := rand.IntN(6) + 1
		dice2 := rand.IntN(6) + 1
		sum = dice1 + dice2

		if sum == val {
			verb, noun := getRollWords(i)
			fmt.Printf("Выпало %d и %d, в сумме %d, на это %s %d %s.\n", dice1, dice2, sum, verb, i, noun)
			return
		} else {
			fmt.Printf("Выпало %d и %d, в сумме %d, бросаем еще раз.\n", dice1, dice2, sum)
		}
	}
}
