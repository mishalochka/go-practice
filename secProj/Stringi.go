package main

import (
	"bufio"
	"fmt"
	"os"
)

////////////DUP1

// func main() {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}

// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

////////DUP3

// func main() {
// 	counts := make(map[string]int)
// 	for _, filename := range os.Args[1:] {
// 		data, err := ioutil.ReadFile(filename)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
// 			continue
// 		}
// 		for _, line := range strings.Split(string(data), "\n") {
// 			counts[line]++
// 		}
// 	}
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

/////// DUP 2

// func main() {
// 	counts := make(map[string]int)
// 	files := os.Args[1:]
// 	if len(files) == 0 {
// 		countLines(os.Stdin, counts)
// 	} else {
// 		for _, arg := range files {
// 			f, err := os.Open(arg)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
// 				continue
// 			}
// 			countLines(f, counts)
// 			f.Close()
// 		}
// 	}
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// func countLines(f *os.File, counts map[string]int) {
// 	input := bufio.NewScanner(f)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// }

// Ex 1.4
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			counts[input.Text()]++
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}

	} else {

		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			input := bufio.NewScanner(f)
			for input.Scan() {
				counts[input.Text()]++
				if counts[input.Text()] > 1 {
					fmt.Println(arg)
					break
				}
			}
			f.Close()
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}
