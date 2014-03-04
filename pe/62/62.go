package main

import (
	"fmt"
	"sort"
	"strconv"
  "github.com/ebensh/pe/common"
)

func main() {
	max := 30000
	ch := make(chan int64)
	go generate_cube_hashes(max, ch)
	cube_hash_counts := make(map[int64]int)
	hash := int64(0)
	for i := 0; i < max; i++ {
		hash = <-ch
		cube_hash_counts[hash]++
	}

	for key, value := range cube_hash_counts {
		if value == 5 {
			fmt.Println("Hash:", key)
		}
	}
}

func generate_cube_hashes(max int, ch chan int64) {
	for i := int64(1); i <= int64(max); i++ {
		ch <- hash_cube(i)
	}
}

func hash_cube(n int64) int64 {
	cubed := n * n * n
	digits := pecommon.ToDigitsBytes(cubed)
	sort.Sort(sort.Reverse(pecommon.ByteSlice(digits)))
	hash, _ := strconv.ParseInt(string(digits), 10, 0)
	fmt.Println(n, hash)
	return hash
}
