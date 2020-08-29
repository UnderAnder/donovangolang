package main

import (
	"testing"
	"time"
)

func Benchmark_printArgsFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printArgsFor(time.Now())
	}
}

func Benchmark_printArgsRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printArgsRange(time.Now())
	}
}
