package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

func servHome(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	rendered(w, "home.gohtml", map[string]int{"path": 1,})
}

func servTest(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	rendered(w, "index.gohtml", map[string]int{"path": 2,})
}

func rendered(w http.ResponseWriter, f string, v map[string]int) {
	tpl.ExecuteTemplate(w, f, v)
	// writer, file, v
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("views/*"))
}

func main() {
	http.HandleFunc("/", servHome)
	http.HandleFunc("/test", servTest)

	fmt.Println("Listening on port 1337")
	http.ListenAndServe(":1337", nil)
}
