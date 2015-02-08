package main

import (
	"fmt"
	"math/rand"
	"time"
)

func PI(samples int) float64 {
	var inside int = 0

	for i := 0; i < samples; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if (x*x + y*y) < 1 {
			inside++
		}
	}

	ratio := float64(inside) / float64(samples)

	return ratio * 4
}

func MultiPI(samples int, threads int) float64 {
	threadSamples := samples / threads
	results := make(chan float64, threads)

	for j := 0; j < threads; j++ {
		go func() {
			var inside int
			for i := 0; i < threadSamples; i++ {
				x, y := rand.Float64(), rand.Float64()

				if x*x+y*y <= 1 {
					inside++
				}
			}
			results <- float64(inside) / float64(threadSamples) * 4
		}()
	}

	var total float64
	for i := 0; i < threads; i++ {
		total += <-results
	}

	return total / float64(threads)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("Running Monte Carlo simulations ...\n")
	fmt.Println("Our value of Pi after 1,000,000 runs:\t\t", PI(1000000))
	fmt.Println("Our value of Pi after 1,000,000 runs (MT):\t", MultiPI(1000000, 4))
}
