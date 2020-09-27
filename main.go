package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := "test"
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}
