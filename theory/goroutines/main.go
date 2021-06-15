package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := []string{"nico", "flynn", "yoon", "larry", "pages", "nolan"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println("Waiting for messages")
	for i := 0; i < len(people); i++ {
		fmt.Println("Received this message:", <-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 3)
	c <- person + " is sexy"
}
