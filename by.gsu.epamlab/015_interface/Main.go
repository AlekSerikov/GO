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

	doc := new(Document)
	pcard := new(PersonCard)

	doc.Date = "1.10.2018"
	doc.NumberOfPages = 5
	doc.Number = "A - 100"

	pcard.Date = "1.10.2018"
	pcard.Age = 21
	pcard.FirstName = "User"
	pcard.LastName = "Test"

	sl := []PrintInterface{doc, pcard}

	for _, v := range sl {
		v.PrintMessage()
	}

}
