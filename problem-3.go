package main

import "fmt"

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
  end := 600851475143

  for i := 0; i < end; i++ {
    if !isPrime(i) {
      continue
    }

    if end%i == 0 {
      primeFactors = append(primeFactors, i)
    }
  }

  fmt.Println(primeFactors)
  fmt.Println("The largest prime factor is: ", primeFactors[len(primeFactors)-1])
}
