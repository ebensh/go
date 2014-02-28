package main

import (
	"fmt"
  "sort"
  "strconv"
)

// ByteSlice attaches the methods of Interface to []byte,
// sorting in increasing order.
type ByteSlice []byte

func (p ByteSlice) Len() int           { return len(p) }
func (p ByteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p ByteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByteSlice) Sort() { sort.Sort(p) }

func main() {
  //cube_counts := make(map[int]int)
  for i := int64(1); i <= 10; i++ {
    cubed := i * i * i
    digits := []byte(strconv.FormatInt(cubed, 10))
    sort.Sort(ByteSlice(digits))
    sorted_digits := string(digits)
    fmt.Println(sorted_digits)
  }
}
