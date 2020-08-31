package main

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(342)
	}
}

func BenchmarkPopCountCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountCycle(342)
	}
}
