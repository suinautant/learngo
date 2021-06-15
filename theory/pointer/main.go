package main

import "fmt"

func main() {
	a := 2
	b := &a
	a = 10

	fmt.Println("value", a, b)
	fmt.Println("address", &a, &b)
	fmt.Println("address and value", &a, b)
	fmt.Println("value and pointer value", a, *b)

	*b = 20
	fmt.Println(a)
}
