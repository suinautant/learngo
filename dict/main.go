package main

import (
	"fmt"

	"github.com/suinautant/learngo/dict/mydict"
)

func errPrint(e error) {
	if e != nil {
		fmt.Println(e)
	}

}

func main() {
	// using fuc dictionary.Update()
	// change definition of map
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	// err := dictionary.Add(baseWord, "First")
	// errPrint(err)
	// if err != nil {
	// fmt.Println(err)
	// }
	// err := dictionary.Update("baseWord", "Second")
	dictionary.Update(baseWord, "Second")
	// errPrint(err)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	rtnWord, _ := dictionary.Search(baseWord)
	fmt.Println(rtnWord)
	err := dictionary.Delete(baseWord)
	errPrint(err)
	rtnWord2, err2 := dictionary.Search(baseWord)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(rtnWord2)
	}

	// dictionary := mydict.Dictionary{"first": "First word"}
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// func dictionary.Add()
	// dictionary := mydict.Dictionary{"first": "First word"}
	// fmt.Println(dictionary)
	// word := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, err := dictionary.Search(word)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(hello, " is found")

	// // Add same word
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	// hello2, err2 := dictionary.Search(word)
	// if err != nil {
	// 	fmt.Println(err2)
	// }
	// fmt.Println(hello2, " is found")
}
