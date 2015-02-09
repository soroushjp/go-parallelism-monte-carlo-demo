package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func PI(samples int) float64 {
	var inside int = 0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < samples; i++ {
		x := r.Float64()
		y := r.Float64()
		if (x*x + y*y) < 1 {
			inside++
		}
	}

	ratio := float64(inside) / float64(samples)

	return ratio * 4
}

func MultiPI(samples int) float64 {
	cpus := runtime.NumCPU()

	threadSamples := samples / cpus
	results := make(chan float64, cpus)

	for j := 0; j < cpus; j++ {
		go func() {
			var inside int
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < threadSamples; i++ {
				x, y := r.Float64(), r.Float64()

				if x*x+y*y <= 1 {
					inside++
				}
			}
			results <- float64(inside) / float64(threadSamples) * 4
		}()
	}

	var total float64
	for i := 0; i < cpus; i++ {
		total += <-results
	}

	return total / float64(cpus)
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	fmt.Println("Running Monte Carlo simulations ...\n")
	fmt.Println("Our value of Pi after 1,000,000 runs:\t\t", PI(1000000))
	fmt.Println("Our value of Pi after 1,000,000 runs (MT):\t", MultiPI(1000000))
}
