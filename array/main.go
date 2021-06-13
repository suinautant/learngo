package main

import "fmt"

func main() {
	// names := [5]string{"aa", "bb", "cc"}
	// names[4] = "ee"
	// fmt.Println(names)

	names := []string{"aa", "bb", "cc"}
	fmt.Println(names)
	names = append(names, "dd")
	names = append(names, "ee", "ff")
	// varArr := []string{"new", "old", "variety"}
	// names = append(names, varArr) // err
	// names = append(names, {"new", "old"}) // err
	fmt.Println(names)
}
