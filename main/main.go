// https://www.hackerrank.com/challenges/priyanka-and-toys/problem

package main

import (
	"fmt"
	"sort"
)

func toys(w []int32) int32 {
	sort.Slice(w, func(i, j int) bool {
		return w[i] < w[j]
	})

	var containerCount = 0
	var idx = 0

	for {
		var curr = w[idx]
		var max = curr + 4

		for ; idx < len(w); idx++ {
			if w[idx] > max {
				break
			}
		}

		containerCount++

		if idx >= len(w) {
			break
		}
	}

	return int32(containerCount)
}

func main() {
	fmt.Println(toys([]int32{1, 2, 3, 21, 7, 12, 14, 21}))
}
