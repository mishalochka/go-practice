package main

import (
	"fmt"
	"os"
)

// func main() {
// 	s := ""
// 	sep := ""
// 	for _, arg := range os.Args[1:] {
// 		s += sep + arg
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// func main (){
// 	s, sep := "", ""
// 	for i := 1; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	fmt.Println(strings.Join(os.Args[0:], " "))
// }

// func main() {
// 	s, sep := "", ""
// 	for i := 1; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 		fmt.Println(i, s)
// 	}

// }

// func main() {
// 	//s := ""
// 	for i := 1; i < len(os.Args); i++ {
// 		fmt.Println(os.Args[i], i)
// 	}

// }

// func main() {
// 	for i, arg := range os.Args[1:] {
// 		fmt.Println(i, arg)
// 	}
// }

// func main() {
// 	for i, arg := range os.Args {
// 		if i == 0 {
// 			continue
// 		}
// 		fmt.Println(i, arg)
// 	}
// }

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
}
