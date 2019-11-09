package main

import (
 "fmt"
 "math/rand"
// "time"
)

func main() {
 // rand.Seed(time.Now().UnixNano())

	fmt.Printf("Random Number : %d\n", rand.Intn(10))
}
