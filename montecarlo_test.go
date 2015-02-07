package main

import (
	"testing"
)

const sampleCount = 1000000

func BenchmarkGetPi(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetPi(sampleCount)
	}
}

func BenchmarkGetPiMulti(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetPiMulti(sampleCount)
	}
}
