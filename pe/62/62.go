package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"

	pecommon "github.com/ebensh/pe/common"
)

func main() {
	max := 30000
	ch := make(chan int64, 50)

	cube_hash_counts := make(map[int64]int)
	//hash := int64(0)
	// waitgroup, mutex
	var wg sync.WaitGroup
	for i := 0; i < max; i++ {
		wg.Add(1)
		go func(i int) {
			generate_cube_hashes(i, ch)
			wg.Done()
		}(i)
	}

	done := make(chan int)
	go func() {
		for hash := range ch {
			cube_hash_counts[hash]++
		}
		done <- 1
	}()

	wg.Wait()
	close(ch)
	<-done

	for key, value := range cube_hash_counts {
		if value == 5 {
			fmt.Println("Hash:", key)
		}
	}
}

func generate_cube_hashes(max int, ch chan int64) {
	//for i := int64(1); i <= int64(max); i++ {
	ch <- hash_cube(int64(max))
	//}
}

func hash_cube(n int64) int64 {
	cubed := n * n * n
	digits := pecommon.ToDigitsBytes(cubed)
	sort.Sort(sort.Reverse(pecommon.ByteSlice(digits)))
	hash, _ := strconv.ParseInt(string(digits), 10, 0)
	fmt.Println(n, hash)
	return hash
}
