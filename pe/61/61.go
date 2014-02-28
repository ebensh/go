package main

import (
	"fmt"
	"sort"
	"strconv"
)

// ByteSlice attaches the methods of Interface to []byte,
// sorting in increasing order.
type ByteSlice []byte

func (p ByteSlice) Len() int { return len(p) }

// We reverse Less here to mean More so that 0's go at the right end,
// instead of getting lost in the hash.
func (p ByteSlice) Less(i, j int) bool { return p[i] > p[j] }
func (p ByteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByteSlice) Sort()              { sort.Sort(p) }

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
	digits := []byte(strconv.FormatInt(cubed, 10))
	sort.Sort(ByteSlice(digits))
	hash, _ := strconv.ParseInt(string(digits), 10, 64)
	fmt.Println(n, hash)
	return hash
}
