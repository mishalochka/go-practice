package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
)

func main() {
	for range 100 {
		random = rand.IntN(100) + 1
		guesses = 0
		result := play()
		if result != random {
			fmt.Printf("Неверный ответ. Было загадано число %d, а в ответе получили число %d", random, result)
			os.Exit(-1)
		} else {
			fmt.Println(result, random)
			return
		}
	}
}

var guesses int
var random int

func guess(num int) (int, error) {
	if guesses >= 6 {
		return 0, errors.New("too many attempts")
	}
	guesses++
	if num > random {
		return -1, nil
	}
	if num < random {
		return 1, nil
	}
	return 0, nil
}

func play() int {
	min := 1
	max := 100
	mid := 0
	for min <= max {
		mid := (max + min) / 2
		result, _ := guess(mid)
		switch result {
		case 1:
			min = mid + 1
		case -1:
			max = mid - 1
		case 0:
			return mid
		}
	}
	return mid
}
