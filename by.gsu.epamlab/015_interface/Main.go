package main

import "fmt"

/*
 Полиморфизм
1 - Возможность объектов с одинаковой спецификацией иметь различную реализацию.
2 - Способность обьекта использовать методы производного класса
	который не существует на момент создания базового.
3 - Это свойство классов иметь одинаковые методы,
	которые будут работать по-разному в контексте объектов.
*/

type Document struct {
	Date          string
	Number        string
	NumberOfPages int
}

type PersonCard struct {
	Date      string
	FirstName string
	LastName  string
	Age       int
}

type PrintInterface interface {
	PrintMessage()
}

func (d *Document) PrintMessage() {
	fmt.Println("Im Document!")
}

func (p *PersonCard) PrintMessage() {
	fmt.Println("Im PersonCard!")
}

func main() {

	var doc PrintInterface = &Document{"1.10.2018", "A - 100", 5}
	var pcard PrintInterface = &PersonCard{"1.10.2018", "User", "Test", 21}

	sl := []PrintInterface{doc, pcard}

	for _, v := range sl {
		v.PrintMessage()
	}

}
