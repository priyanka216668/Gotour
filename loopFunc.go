package main

import "fmt"

func printOddNumber(limit int) {
  fmt.Println("Odd Numbers from 0 To", limit)
  for i := 0 ; i <= limit; i++ {
    if  i % 2 != 0 {
      fmt.Printf( "%d ", i)
    }
  }
}

func main() {
  limit := 10
  printOddNumber(limit)
}
