// Package pecommon contains common ProjectEuler utility classes and methods.
package pecommon

import (
	"sort"
	"strconv"
)

// ByteSlice attaches the methods of Interface to []byte,
// sorting in increasing order.
type ByteSlice []byte

func (p ByteSlice) Len() int { return len(p) }

// We reverse Less here to mean More so that 0's go at the right end,
// instead of getting lost in the hash.
func (p ByteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p ByteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByteSlice) Sort()              { sort.Sort(p) }

func ToDigitsInts(i int64) []int {
	digits_bytes := ByteSlice(strconv.FormatInt(int64(i), 10))
	digits_ints := make([]int, digits_bytes.Len())
	for index, digit_byte := range digits_bytes {
		digits_ints[index] = int(digit_byte - '0')
	}
	return digits_ints
}

func ToDigitsBytes(i int64) []byte {
	return ByteSlice(strconv.FormatInt(int64(i), 10))
}
