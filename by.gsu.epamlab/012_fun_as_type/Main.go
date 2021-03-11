package main

import "fmt"

func main() {

	f := func(arg int) int {
		return arg * 2
	}
	result := f(2)
	fmt.Println(result)

	result2 := mulValueOnArgFunValue(2, f)
	fmt.Println(result2)

	result3 := mulValueOnArgFunValue(3, func(arg int) int {
		return arg * 2
	})

	fmt.Println(result3)
}

func mulValueOnArgFunValue(num int, justFun func(int) int) int {
	return justFun(num)
}
