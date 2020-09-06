package main

import (
	"../../chapter2/popcount"
	"crypto/sha256"
	"fmt"
)

func popCountArray32(x [32]byte) int {
	var count int
	for _, j := range x {
		count += popcount.PopCount(uint64(j))
	}
	return count
}

func diffBits(x, y [sha256.Size]byte) int {
	if x == y {
		return 0
	}
	var counter int
	for i := range x {
		counter += bitCount(x[i] ^ y[i])
	}
	return counter
}

func bitCount(x uint8) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Printf("diff bits: %v\n", diffBits(c1, c2))
}
