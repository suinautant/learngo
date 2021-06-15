package main

import "fmt"

func superAdd(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// 	total += numbers[i]
	// }
	return total
}

func main() {
	result := superAdd(6, 5, 4, 3, 2, 1)
	fmt.Println(result)
}
