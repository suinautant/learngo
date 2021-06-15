package main

import "fmt"

func canIDrink(age int) bool {
	// switch age {
	// case 18:
	// 	return true
	// case 20:
	// 	return true
	// }
	// return false

	// switch {
	// case age >= 50:
	// 	return false
	// case age >= 18:
	// 	return true
	// }
	// return false

	switch koreanAge := age + 2; {
	case koreanAge >= 50:
		return false
	case koreanAge >= 18:
		return true
	}
	return false
}

func main() {
	result := canIDrink(50)
	fmt.Println(result)

}
