package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1024*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Split(bufio.ScanWords)

	scanInt := func() int {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		return val
	}

	if !scanner.Scan() {
		return
	}
	nStr := scanner.Text()
	n, _ := strconv.Atoi(nStr)

	if n <= 0 {
		fmt.Println(0)
		return
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}

	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = scanInt()
	}

	var m int
	if scanner.Scan() {
		m, _ = strconv.Atoi(scanner.Text())
	}

	var x, y []int
	if m > 0 {
		x = make([]int, m)
		for i := 0; i < m; i++ {
			x[i] = scanInt()
		}
		y = make([]int, m)
		for i := 0; i < m; i++ {
			y[i] = scanInt()
		}
	}

	lens := []int{0, 1, 2, 4}
	pts := []int{0, 1, 3, 5}

	for i := 0; i < n-1; i++ {
		if a[i]+lens[b[i]] >= a[i+1] {
			fmt.Println(0)
			return
		}
	}

	score := 0
	j := 0
	for i := 0; i < n; i++ {
		for j+1 < m && x[j+1] <= a[i] {
			j++
		}

		if j < m && x[j] <= a[i] && x[j]+y[j] >= a[i]+lens[b[i]] {
			score += pts[b[i]]
		} else {
			score -= 1
		}
	}

	if score < 0 {
		score = 0
	}

	fmt.Println(score)
}
