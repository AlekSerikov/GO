package main

import "fmt"

func AddValueByPointer(x *int) {
	*x = (*x) * 2
}

func main() {

	//pointer in function
	startValue := 12

	AddValueByPointer(&startValue)

	fmt.Println(startValue) //start value is changed = 24

	//pointer

	x := 15

	var pointer *int = &x

	fmt.Println(*pointer)
}
