package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Error create file: ", err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}

// func WriteFile(content string, name string) {
// 	data := "Hello, golang! "
// 	err := os.WriteFile("ex1.txt", []byte(data), 0644)
// 	if err != nil {
// 		fmt.Println("Ошибка записи файла: ", err)
// 	}

// }

// func ReadFile() {
// 	data, err := os.ReadFile("ex1.txt")
// 	if err != nil {
// 		fmt.Println("Error reading file: ", err)
// 		return
// 	}
// 	fmt.Println("Содержимое: ")
// 	fmt.Println(string(data))
// }
