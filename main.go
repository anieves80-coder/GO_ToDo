package main

import (
	"fmt"
	"html/template"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var tpl *template.Template

//Data struct for data coming in from the clients ajax request
type Data struct {    
	Date string `json:"date"`
	Description string `json:"info"`
}

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))	
}

func main() {	
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", index)
	http.HandleFunc("/addRec", addRec)
	
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	r := getAll()	
	tpl.ExecuteTemplate(w, "index.gohtml", r)
}

func addRec(w http.ResponseWriter, req *http.Request) {
	var getInfo Data	
	body, _ := ioutil.ReadAll(req.Body)//reads a buffer and returns a byte array	
	json.Unmarshal(body, &getInfo)//Parses the JSON to the struct getInfo
	m := insertDB(getInfo)	
	mMap := map[string]string{"msg":m}
	
	m2,_ := json.Marshal(mMap)
	if m != "err"{
		fmt.Fprintln(w, string(m2))
	}
}