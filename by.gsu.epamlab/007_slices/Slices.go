package main

import "fmt"

func main() {

	var slice []int

	slice = append(slice, 1, 2, 3, 4)

	var slice2 []int = append(slice, 1, 2, 3, 4)

	for _, v := range slice {

		fmt.Printf("%v", v)

	}

	fmt.Println(append(slice, slice2...))

	fmt.Println(slice)

}
