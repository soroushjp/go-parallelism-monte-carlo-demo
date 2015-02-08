package main

import (
	"testing"
)

const samples = 10000000

func BenchmarkPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PI(samples)
	}
}

func BenchmarkMultiPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultiPI(samples, 4)
	}
}
