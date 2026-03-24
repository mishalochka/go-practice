package main

import (
	"fmt"
)

func main() {
	printDiamond(3)
}

func printTable(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			fmt.Printf("%d x %d = %d", i, j, i*j)
			if j == num {
				continue
			} else {
				fmt.Print("\t")
			}
		}
		fmt.Print("\n")
	}
}

func printDiamond(n int) {
	fmt.Println("Мой бриллиант:")

	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		fmt.Print("#")
		if i > 1 {
			for z := 1; z <= (i-1)*2-1; z++ {
				fmt.Print(" ")
			}
			fmt.Print("#")
		}

		fmt.Print("\n")
	}

	for i := n - 1; i >= 1; i-- {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		fmt.Print("#")
		if i > 1 {
			for z := 1; z <= (i-1)*2-1; z++ {
				fmt.Print(" ")
			}
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}

// for i := 1; i <= n; i++ {
// 	for j := 1; j <= n; j++ {
// 		if i >= j {
// 			fmt.Printf(" ")
// 		}
// 		if n == j {
// 			fmt.Print("#\n")
// 		}
// 	}
// }
