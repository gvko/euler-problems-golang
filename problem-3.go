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

func filterPrimeFactors(primeFactors []int) bool {
  sum := 1
  stop := false

  for _, element := range primeFactors {
    fmt.Println(element)
    sum *= element

    if sum > end {
      //TODO: remove last element and do a recursive call to this function
      primeFactors = primeFactors[:len(primeFactors)-1]
      stop = filterPrimeFactors(primeFactors)
    } else if sum == end {
      stop = true
      break
    } else {
      break
    }
  }

  return stop
}

func main() {
  primeFactors := []int{}
  stop := false

  for i := 0; i < end; i++ {
    if isPrime(i) {
      primeFactors = append(primeFactors, i)
    }

    filterPrimeFactors(primeFactors)

    if stop {
      break
    }
  }

  fmt.Println(primeFactors)
}
