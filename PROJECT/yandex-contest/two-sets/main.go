package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	nums := make([]int64, 40)
	for i := 0; i < 40; i++ {
		if scanner.Scan() {
			nums[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		}
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	seen := make(map[int64]uint64)

	for {
		mask := rng.Uint64() & ((1 << 40) - 1)

		if mask == 0 {
			continue
		}

		var sum int64 = 0
		for i := 0; i < 40; i++ {
			if (mask & (1 << i)) != 0 {
				sum += nums[i]
			}
		}

		if prevMask, ok := seen[sum]; ok && prevMask != mask {
			intersection := prevMask & mask

			m1 := prevMask ^ intersection
			m2 := mask ^ intersection

			if m1 == 0 || m2 == 0 {
				continue
			}

			printMask(m1)
			printMask(m2)
			return
		}

		seen[sum] = mask

		if len(seen) > 3000000 {
			seen = make(map[int64]uint64)
		}
	}
}

func printMask(mask uint64) {
	var indices []int
	for i := 0; i < 40; i++ {
		if (mask & (1 << i)) != 0 {
			indices = append(indices, i+1)
		}
	}

	fmt.Println(len(indices))

	for i, idx := range indices {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(idx)
	}
	fmt.Println()
}
