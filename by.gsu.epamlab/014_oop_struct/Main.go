package main

import "fmt"

type EDocument struct { //class in java
	Number        string
	Date          string
	NumberOfPages int
}

func NewEDocument(number string, date string, numberOfPages int) *EDocument {
	return &EDocument{Number: number, Date: date, NumberOfPages: numberOfPages}
}

func (doc EDocument) getNumberOfPages() int {
	return doc.NumberOfPages
}

//func(doc *EDocument) getNumberOfPages() int { //указатель, ссылка
//	return doc.NumberOfPages
//}

func main() {

	doc1 := EDocument{
		Number:        "001-A",
		Date:          "10.10.2021",
		NumberOfPages: 10,
	}

	var doc2 = EDocument{
		Number:        "002-A",
		Date:          "02.02.2021",
		NumberOfPages: 5,
	}

	doc3 := new(EDocument)
	doc3.Date = "123"
	doc3.Number = "12"
	doc3.NumberOfPages = 12

	doc4 := NewEDocument("qwe", "12.12.12", 50)

	doc5 := EDocument{"name", "12.12.12", 45}

	fmt.Println(doc1, doc2, doc3, doc1.getNumberOfPages(), doc4, doc5)

}
