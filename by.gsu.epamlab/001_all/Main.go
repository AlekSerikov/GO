package main

import "fmt"

func main() {
	var a int
	var b = 100
	c := 150

	const constant int = 100
	const constantTwo = 150

	const (
		A = iota
		B = iota
		C = iota
	)

	//array
	var array [5]int
	array[0] = 1

	var arrayTwo = [...]int{1, 2, 3}

	//slice
	var slice = []int{1, 2, 3, 4, 5}
	slice = append(slice, 6, 7, 999)
	slice2 := []int{6, 7, 8}

	//map
	var firstMap = map[int]string{}
	firstMap[1] = "String1"
	firstMap[2] = "String2"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(constant)
	fmt.Println(constantTwo)
	fmt.Println(len(arrayTwo))
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(firstMap)

}
