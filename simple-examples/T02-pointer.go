package main

import "fmt"

func mf(m *int) {
	//解引用 给变量m赋值
	*m = 8
}

func nf(n int) {
	n = 8
}

func main() {

	var i int = 0
	fmt.Println(i)

	var pi *int
	pi = &i
	fmt.Println(pi)

	m := 0
	mf(&m)
	fmt.Println(m)

	n := 0
	nf(n)
	fmt.Println(n)
}
