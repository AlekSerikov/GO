package main

import "fmt"

func main() {

	fmt.Println(sumFun(2, 2), sumFun2(3, 3), sumFun3(4, 4))

	a, b, c := addToEachFun(1, 2, 3)

	fmt.Println(a, b, c)

	//not my

	//sum(5, 2)
	//fmt.Println(question("word"))

	//result, _ := someFunc(-1, 4, 323)
	//fmt.Println(result)

	//someFuncTwo(10, 55)
	fmt.Println(onlyTypes(10, "23"))

}

func sumFun(a, b int) int {
	return a + b
}

func sumFun2(a, b int) (result int) {
	result = a + b
	return result
}

func sumFun3(a, b int) (result int) {
	result = a + b
	return
}

func addToEachFun(a, b, c int) (x, y, z int) {
	addedValue := 5
	x = a + addedValue
	y = b + addedValue
	z = c + addedValue
	return
}

func addToEachFun2(a, b, c int) (int, int, int) {
	addedValue := 5
	return a + addedValue, b + addedValue, c + addedValue
}

// not my

func sum(a int, b int) {

	var z int
	z = a + b
	fmt.Println(z)

}

func question(q string) bool {
	if q == "question" {
		return true
	} else {
		return false
	}
}

func someFunc(a, b, c int) (int, bool) {
	if a > 0 && b > 5 {
		return 10, true
	} else if c > 100 {
		return 100, true
	} else {
		return 0, false
	}
}

func someFuncTwo(z, _ int) {
	fmt.Println(z)
}

func onlyTypes(int, string) bool {
	return true
}
