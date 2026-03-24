package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
)

func main() {
	// Внешний бесконечный цикл отвечает только за перезапуск игры
	for {
		// Запускаем один раунд. Функция вернет true, если игрок ввел "выход"
		userQuit := playRound()
		if userQuit {
			break // Прерываем глобальный цикл, игра окончена
		}

		fmt.Println("Хотите сыграть еще раз? (если хотите, напишите слово да):")
		var answer string
		fmt.Scan(&answer)

		// strings.ToLower позволяет принимать "Да", "да" и даже "ДА"
		if strings.ToLower(answer) != "да" {
			break // Пользователь ввел не "да", выходим из цикла
		}
	}

	fmt.Println("Спасибо за игру! До свидания!")
}

// playRound проводит одну игру. Возвращает true, если игрок прервал игру словом "выход"
func playRound() bool {
	target := rand.IntN(100) + 1
	attempts := 0 // Счетчик попыток

	fmt.Println("Компьютер загадал случайное число от 1 до 100 включительно. Угадайте его!")

	// Внутренний цикл отвечает только за попытки угадать число
	for {
		fmt.Print("Ваше предположение (либо, для завершения, введите слово выход): ")

		var input string
		fmt.Scan(&input)

		// Обработка выхода (опять же, защищаемся от разного регистра: "ВЫХОД", "выход")
		if strings.ToLower(input) == "выход" {
			return true
		}

		// Попытка перевести строку в число
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Пожалуйста, введите число или слово 'выход'")
			continue // Ошибка ввода, идем на следующий круг без засчитывания попытки
		}

		attempts++ // Увеличиваем счетчик только для корректных чисел

		// Сравниваем
		if num == target {
			fmt.Printf("Правильно! Вы угадали число с %d попытки.\n", attempts)
			return false // Раунд успешно завершен
		} else if num < target {
			fmt.Println("Загаданное число больше.")
		} else {
			fmt.Println("Загаданное число меньше.")
		}
	}
}
