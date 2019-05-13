package main

import (
  "fmt"
)

func main() {
  upperLimit := 0
  n := 3

  for i := 1; i <= n; i++ {
    upperLimit *= 10
    upperLimit += 9
  }

  lowerLimit := 1 + upperLimit/10
  maxProduct := 0

  fmt.Println("Lower limit: ", lowerLimit)
  fmt.Println("Lower limit: ", upperLimit)

  for i := upperLimit; i >= lowerLimit; i-- {
    for j := i; j >= lowerLimit; j-- {
      product := i * j

      if product < maxProduct {
        break
      }

      number := product
      reverse := 0

      for {
        reverse = reverse*10 + number%10
        number /= 10

        if number != 0 {
          break
        }
      }

      if product == reverse && product > maxProduct {
        maxProduct = product
      }
    }
  }

  fmt.Println(maxProduct)
}
