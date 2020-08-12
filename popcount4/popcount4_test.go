package popcount

import "testing"

func BenchmarkPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(342)
	}
}
