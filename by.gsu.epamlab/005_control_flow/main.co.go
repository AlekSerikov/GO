package main

import "fmt"

func main() {

	//if / else

	a, b, c := 10, 20, 30

	if a > b {
		fmt.Println("a")
	} else if b < c {
		fmt.Println("b")
	} else {
		fmt.Println("c")
	}

	x := 12

	switch x {
	case 12:
		fmt.Println("Двенадцать")
	case 15:
		fmt.Println("Пятнадцать")
	default:
		fmt.Println("No value")
	}

	for x := 1; x < 100; x++ {

		if x%10 == 0 {
			fmt.Println("%10", x)
			continue //прерываем текущую итерацию цикла for
			// break - остановка цикла
		}

		fmt.Println(x)
	}

}
