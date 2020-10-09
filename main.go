package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var tpl *template.Template

//Data struct for data coming in from the clients ajax request for modification
//and getting the ID
type Data struct {
	ID   string `json:"id"`
	Data AddData
}

//AddData struct for data coming in from the clients ajax request and add it
//to the db with the use of a ID
type AddData struct {
	Date        string `json:"date"`
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
	http.HandleFunc("/delRec", delRec)
	http.HandleFunc("/updateRec", updateRec)

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	r := getAll()
	tpl.ExecuteTemplate(w, "index.gohtml", r)
}

func addRec(w http.ResponseWriter, req *http.Request) {
	var getInfo AddData
	body, _ := ioutil.ReadAll(req.Body) //reads a buffer and returns a byte array
	json.Unmarshal(body, &getInfo)      //Parses the JSON to the struct getInfo
	m := insertDB(getInfo)
	if m != "err" {
		fmt.Fprintln(w, `{"msg":"ok"}`)
	}
}

func delRec(w http.ResponseWriter, req *http.Request) {
	var getInfo Data
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &getInfo)
	deleteRec(getInfo.ID)
	fmt.Fprintln(w, `{"msg":"ok"}`)
}

func updateRec(w http.ResponseWriter, req *http.Request) {
	var getInfo Data
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &getInfo)
	m := updateDb(getInfo)
	if m != "err" {
		fmt.Fprintln(w, `{"msg":"ok"}`)
	}
}
