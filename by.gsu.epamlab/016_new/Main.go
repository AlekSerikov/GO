package main

import "fmt"

func main() {

	p := new(int)

	*p = 10

	fmt.Println(*p)
}
