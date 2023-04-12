package main

import "fmt"

func main() {
	name := "Go Developers"
	fmt.Println("Azure for", name)
	var i int = 8
	fmt.Println(&i)

	var pi *int = &i
	fmt.Println(pi)
}
