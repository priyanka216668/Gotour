package main

import "fmt"

func main() {
  {
    defer func() {
      fmt.Println("block: defer runs")
    }()
    fmt.Println("block: ends")
  }
  fmt.Println("main: ends")
}
