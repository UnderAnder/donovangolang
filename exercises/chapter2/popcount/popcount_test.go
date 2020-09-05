package popcount_test

import "testing"
import "../popcount"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(342)
	}
}

func BenchmarkPopCountCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountCycle(342)
	}
}

func BenchmarkPopCountSlide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountSlide(342)
	}
}

func BenchmarkPopCountReset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountReset(342)
	}
}
