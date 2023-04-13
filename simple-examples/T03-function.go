package main

import "fmt"

/*
* anybody的形参：f func(word string) 就是函数签名
 */
func anybodySay(f func(word string)) {
	f("hello")
}
func main() {

	//1、声明了一个函数变量
	var f1 = func(word string) {
		fmt.Println(word, "f1")
	}
	//函数变量的使用，即：函数变量可以当做一个变量一样，传递
	anybodySay(f1)

	//2、声明了一个函数签名
	var ff func(word string)
	//函数签名必须有实现，才能当做参数传递
	ff = f1
	anybodySay(ff)

	//3、给已经声明的函数传递参数，然后再调用
	f1("hhh")
	anybodySay(f1)

}
