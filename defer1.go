package main

import "fmt"

func f(x *int){
  fmt.Println(*x)
}

func main() {
  fmt.Println("Inside Func main()")
  x := 1
  defer f(&x)
  x++
}
