package main

import "fmt"

func canIDrink(age int) bool {
	// if age>= 18 {
	// 	return true
	// } else {
	// 	return false
	// }

	// if age >= 18 {
	// 	return true
	// }
	// return false

	// return age >= 18

	if koreanAge := age + 2; koreanAge >= 18 {
		return true
	}
	return false
}

func main() {
	result := canIDrink(15)
	fmt.Println(result)
}
