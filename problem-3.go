package main

import "fmt"

var end = 13195

func isPrime(value int) bool {
	prime := true

	if value == 0 || value == 1 {
		return false
	}

	for i := 2; i <= value/2; i++ {
		if value%i == 0 {
			prime = false
			break
		}
	}

	return prime
}

func main() {
	primeFactors := []int{}
	sum := 1
	stop := false

	for i := 0; i < end; i++ {
		sum = 1

		if isPrime(i) {
			primeFactors = append(primeFactors, i)
		}

		for _, element := range primeFactors {
			sum *= element

			if sum > end {
				//TODO: remove last element and do a recursive call to this function
				//stop = true
				//fmt.Print("sum: ", sum)
				//break
			}
		}

		if stop {
			break
		}
	}

	fmt.Print(stop, "\n")
	fmt.Print(primeFactors)
}
