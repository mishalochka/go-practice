// package main

// import (
// 	"errors"
// 	"fmt"
// 	"math"
// )

// const IMTPower = 2

// func main() {
// 	for {
// 		fmt.Println("Calc index body weight")
// 		userKg, userHeight := getUserInput()
// 		IMT, err := calculateIMT(userKg, userHeight)
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		outputResult(IMT)
// 		isRepeatCalculation := checkRepeatCalculation()
// 		if !isRepeatCalculation {
// 			break
// 		}
// 	}
// }

// func outputResult(imt float64) {
// 	result := fmt.Sprintf("Your index body weight: %.0f", imt)
// 	fmt.Println(result)
// 	switch {
// 	case imt < 16:
// 		fmt.Println("mini mini index")
// 	case imt < 18.5:
// 		fmt.Println("mini index")
// 	case imt < 25:
// 		fmt.Println("norm index")
// 	case imt < 30:
// 		fmt.Println("max index")
// 	default:
// 		fmt.Println("max max index")
// 	}
// }

// func calculateIMT(userKg float64, userHeight float64) (float64, error) {
// 	if userKg <= 0 || userHeight <= 0 {
// 		return 0, errors.New("No check height or veight")
// 	}
// 	IMT := userKg / math.Pow(userHeight/100, IMTPower)
// 	return IMT, nil
// }

// func getUserInput() (float64, float64) {
// 	var value1 float64
// 	var value2 float64
// 	fmt.Println("Enter the our Veight: ")
// 	fmt.Scan(&value1)
// 	fmt.Println("Enter the our Height: ")
// 	fmt.Scan(&value2)
// 	return value1, value2
// }

// func checkRepeatCalculation() bool {
// 	var userChoicse string
// 	fmt.Println("Do you make know index again (Y/N): ")
// 	fmt.Scan(&userChoicse)
// 	if userChoicse == "Y" || userChoicse == "y" {
// 		return true
// 	}
// 	return false
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"math"
// 	"strings"
// )

// const IMTPower = 2.0

// func main() {
// 	for {
// 		fmt.Println("=== Калькулятор ИМТ ===")
// 		weight, height, err := getUserInput()
// 		if err != nil {
// 			fmt.Printf("Ошибка ввода: %v. Попробуйте снова. \n", err)
// 			continue
// 		}

// 		imt, err := calculateIMT(weight, height)
// 		if err != nil {
// 			fmt.Println("Ошибка расчета:", err)
// 			continue
// 		}

// 		outputResult(imt)

// 		if !shouldRepeat() {
// 			break
// 		}
// 	}
// }

// func calculateIMT(weight, height float64) (float64, error) {
// 	if weight <= 0 || height <= 0 {
// 		return 0, errors.New("Вес и рост должны быть больше нуля")
// 	}
// 	imt := weight / math.Pow(height/100, IMTPower)
// 	return imt, nil
// }

// func getUserInput() (float64, float64, error) {
// 	var w, h float64
// 	fmt.Print("Введите ваш вес (кг): ")
// 	if _, err := fmt.Scan(&w); err != nil {
// 		return 0, 0, errors.New("Некорректный формат веса")
// 	}

// 	fmt.Print("Введите ващ рост (см): ")
// 	if _, err := fmt.Scan(&h); err != nil {
// 		return 0, 0, errors.New("Некорректный формат роста")
// 	}

// 	return w, h, nil
// }

// func outputResult(imt float64) {
// 	fmt.Printf("Ващ индекс массы тела: %.1f\n", imt)

// 	var message string
// 	switch {
// 	case imt < 16:
// 		message = "Выраженный дефицит массы тела"
// 	case imt < 18.5:
// 		message = "Недостаточная масса тела"
// 	case imt < 25:
// 		message = "Норма"
// 	case imt < 30:
// 		message = "Избыточная масса тела"
// 	default:
// 		message = "Ожирение"
// 	}
// 	fmt.Println("Вердикт:", message, imt)
// }

// func shouldRepeat() bool {
// 	var choice string
// 	fmt.Print("Хотите повторить? (Y/N): ")
// 	fmt.Scan(&choice)
// 	choice = strings.ToLower(choice)
// 	return choice == "y" || choice == "yes"
// }

package main

import (
	"fmt"
	"math"
)

func main() {
	var imt float64
	var height float64
	var weight float64
	for {
		height, weight = inputUser()
		imt = imtCalculation(weight, height)
		choiceResult(imt)
		outputResult(imt)
		if !shouldRepeat() {
			break
		}
	}
}

func inputUser() (float64, float64) {
	var weight, height float64
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&weight)
	fmt.Print("Введите свой рост: ")
	fmt.Scan(&height)
	return weight, height
}

func imtCalculation(weight, height float64) float64 {
	var imt float64
	imt = weight / math.Pow(height/100, 2)
	return imt
}

func choiceResult(imt float64) float64 {
	switch {
	case imt < 16:
		fmt.Printf("Сильный дефицит массы тела %.1f\n", imt)
	case imt < 18.5:
		fmt.Printf("Дефицит массы тела %.1f\n", imt)
	case imt < 25:
		fmt.Printf("Норма %.1f\n", imt)
	case imt < 30:
		fmt.Printf("Избыточная масса %.1f\n", imt)
	case imt < 35:
		fmt.Printf("степень ожирения %.1f\n", imt)
	case imt < 40:
		fmt.Printf("2-ая степень ожирения %.1f\n ", imt)
	default:
		fmt.Printf("3-я степень ожирения %.1f\n", imt)
	}
	return imt
}

func outputResult(imt float64) {
	fmt.Printf("Ващ индекс массы тела: %.1f\n", imt)
}

func shouldRepeat() bool {
	var choice string
	fmt.Print("Хотите повторить? (Y/n): ")
	fmt.Scan(&choice)
	return choice == "y" || choice == "yes"
}
