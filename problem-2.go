package main

import "fmt"

func main() {
  end := 4000000
  termA := 1
  termB := 2
  nextTerm := 0
  sum := termB

  for {

    nextTerm = termA + termB
    termA = termB
    termB = nextTerm

    if nextTerm%2 == 0 {
      sum += nextTerm
    }

    fmt.Print(nextTerm, "\n")

    if termA >= end || termB >= end {
      break
    }
  }

  fmt.Print("sum: ", sum, "\n")
}
