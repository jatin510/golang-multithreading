package main

import (
	"log"
	"math"
	"time"
)

var MAX_INT = 100000000
var totalPrimeNumbers int32 = 0

func checkPrime(x int) {
	if x&1 == 0 {
		return
	}

	for i := 3; i <= int(math.Sqrt((float64(x)))); i++ {
		if x%i == 0 {
			return
		}
	}

	totalPrimeNumbers++
}

func main() {
	start := time.Now()

	for i := 3; i < MAX_INT; i++ {
		checkPrime(i)
	}

	log.Println("checking till", MAX_INT, "prime numbers found", totalPrimeNumbers+1, "took", time.Since(start), "seconds")
}
