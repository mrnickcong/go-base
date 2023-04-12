package main

import "fmt"

type Person struct {
	name string
}

func m1(person Person) {
	person.name = "李四"
}

func m2(person *Person) {
	person.name = "李四"
}

func main() {

	person := Person{name: "张三"}
	fmt.Println(person.name)

	m1(person)
	fmt.Println(person.name)

	m2(&person)
	fmt.Println(person.name)

	m1(person)
	fmt.Println(person.name)
}
