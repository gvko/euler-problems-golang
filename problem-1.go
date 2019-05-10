package main

import "fmt"

var bigNumber = 1000

func main() {
  sum := 0

  for i := 0; i < bigNumber; i++ {
    if i%3 == 0 || i%5 == 0 {
      sum += i
    }
  }

  fmt.Print(sum)
  fmt.Print("\n")
}
