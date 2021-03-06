package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountCycle returns the population count (number of set bits) of x.
func PopCountCycle(x uint64) int {
	var count byte
	for i := 0; i < 8; i++ {
		count += pc[byte(x>>(i*8))]
	}
	return int(count)
}

func PopCountSlide(x uint64) int {
	var count int

	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			count++
		}
	}
	return int(count)
}

// PopCountReset returns the population count (number of set bits) of x.
func PopCountReset(x uint64) int {
	var count int

	for ; x != 0; x &= x - 1 { // сбрасывает крайний справа ненулевой бит
		count++
	}
	return count
}
