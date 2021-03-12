package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

// рекурсивный вывод списка
func printNodeValue(n *node) {

	fmt.Println(n.value)
	if n.next != nil {
		printNodeValue(n.next)
	}
}
func main() {

	first := node{value: 4}
	second := node{value: 5}
	third := node{value: 6}

	first.next = &second
	second.next = &third

	var current *node = &first

	printNodeValue(current) //same

	for current != nil { //same
		fmt.Println(current.value)
		current = current.next
	}

	fmt.Println()

	personIgor := person{
		name: "Igor",
		age:  45,
	}

	//method presentation

	fmt.Println(personIgor)

	personIgor.updateAge(100)

	fmt.Println(personIgor)
}

//method

type person struct {
	name string
	age  int
}

// если указываем * , то ссылаемся на текущий объект и будем его менять!!!
func (p *person) updateAge(newAge int) {
	p.age = newAge
}
