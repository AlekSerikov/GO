package main

import "fmt"

func main() {

	intCh := make(chan int, 3)

	go func() {
		fmt.Println("Go routine starts")
		intCh <- 5 // блокировка, пока данные не будут получены функцией main
		intCh <- 4
		intCh <- 2
	}()
	fmt.Println(<-intCh, <-intCh, <-intCh) // получение данных из канала
	fmt.Println("The End")
}
