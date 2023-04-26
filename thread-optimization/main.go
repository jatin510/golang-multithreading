package main

import (
	"log"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var MAX_INT = 100000000

// var MAX_INT = 1000000
var CONCURRENCY = 10
var totalPrimeNumbers int32 = 0
var currentNumber int32 = 0

func checkPrime(x int) {
	if x&1 == 0 {
		return
	}

	for i := 3; i <= int(math.Sqrt((float64(x)))); i++ {
		if x%i == 0 {
			return
		}
	}

	atomic.AddInt32(&totalPrimeNumbers, 1)
}

func doWork(name string, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()

	for {
		x := atomic.AddInt32(&currentNumber, 1)
		if x > int32(MAX_INT) {
			break
		}

		checkPrime(int(x))
	}

	log.Printf("thread %s completed in  %s", name, time.Since(start))
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < CONCURRENCY; i++ {
		wg.Add(1)
		go doWork(strconv.Itoa(i), &wg)
	}

	wg.Wait()

	log.Println("checking till", MAX_INT, "prime numbers found", totalPrimeNumbers+1, "took", time.Since(start), "seconds")

}
