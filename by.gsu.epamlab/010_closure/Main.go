package main

import "fmt"

func main() {

	fmt.Println(someFun("one", "two"))

	closure(0)

}

func someFun(a, b string) string {

	x := a + " " + b

	x = func(s string) string {
		return x + s
	}(" from anon fun")

	return x
}

func closure(s int) {
	fmt.Println("Start s value: ", s)

	x := func() {
		s = 100
		fmt.Println(s)
	}

	y := func() {
		s = 200
		fmt.Println(s)
	}

	y()
	x()
}
