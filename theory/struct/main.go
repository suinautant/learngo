package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	// nico := person{"nico", 18, favFood}
	// nico := person{name: "nico", age: 18, favFood: favFood}
	nico := person{name: "nico", favFood: favFood}
	fmt.Println(nico)
	fmt.Println(nico.name)
	fmt.Println(nico.age)
	fmt.Println(nico.favFood)
	fmt.Println(nico.favFood[0])
	fmt.Println(nico.favFood[1])

	for _, v := range nico.favFood {
		fmt.Println(v)
	}

	teacher := &nico
	teacher.name = "sir nico"
	fmt.Println("*teacher : ", *teacher)
	fmt.Println("teacher : ", teacher)
	fmt.Println("&teacher : ", &teacher)
	fmt.Println("nico : ", nico)
}
