package main

import (
	"fmt"
)

func main() {

	var array [3]int

	array[0] = 1
	array[1] = 2
	array[2] = 3

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

}
