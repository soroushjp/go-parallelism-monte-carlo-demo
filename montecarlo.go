package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetPi(samples int) float64 {
	var inside int = 0

	for i := 0; i < samples; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if (x*x + y*y) < 1 {
			inside++
		}
	}

	return float64(inside) / float64(samples) * 4
}

func GetPiMulti(samples int) float64 {
	NCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(NCPU)

	results := make(chan float64, NCPU)
	wg := sync.WaitGroup{}

	for j := 0; j < NCPU; j++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			var pi float64
			var inside int = 0
			var threadSamples = samples / NCPU
			for i := 0; i < threadSamples; i++ {
				x := rand.Float64()
				y := rand.Float64()
				if (x*x + y*y) < 1 {
					inside++
				}
			}
			pi = float64(inside) / float64(threadSamples) * 4
			results <- pi
		}()
	}

	wg.Wait()
	var piTotal float64
	for t := 0; t < NCPU; t++ {
		piTotal += <-results
	}
	pi := piTotal / float64(NCPU)

	return pi
}

func main() {
	fmt.Println("Our value of Pi after 100 runs:\t\t\t\t\t", GetPi(100))
	fmt.Println("Our value of Pi after 1,000 runs:\t\t\t\t", GetPi(1000))
	fmt.Println("Our value of Pi after 10,000 runs:\t\t\t\t", GetPi(10000))
	fmt.Println("Our value of Pi after 100,000 runs:\t\t\t\t", GetPi(100000))
	fmt.Println("Our value of Pi after 1,000,000 runs:\t\t\t\t", GetPi(1000000))

	fmt.Println("Our value of Pi after 1,000,000 runs: (multithreaded)\t\t", GetPiMulti(1000000))
}
