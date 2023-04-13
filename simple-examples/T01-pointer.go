package main

import "fmt"

type Person struct {
	name string
}

/*
函数的形参：Person类型
函数的实参：Person的值
传递过程：将person的值，复制传递
*/
func m1(person Person) {
	person.name = "李四"
}

/*
函数的形参：person指针类型
函数的实参：person指针
传递过程：将person指针的值，复制传递
*/
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
