
package main

import "fmt"

func swap(x, y *string) {
	var temp string
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	var a , b string = "hello", "world"
	swap(&a, &b)
	fmt.Println(a, b)
}
