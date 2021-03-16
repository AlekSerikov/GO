package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/footer.html", "templates/header.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func handleFun() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8888", nil)
}

func main() {
	handleFun()

}
