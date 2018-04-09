package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var views = template.Must(template.ParseGlob("views/*.html"))

func init() {

}

func main() {
	address := ":3000"
	http.HandleFunc("/", getMain)
	http.HandleFunc("/books", createBook)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalf("error listening: %s", err)
	}
}

func getMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, _ := template.ParseFiles("books.html")
	t.Execute(w, "Title")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	views.ExecuteTemplate(w, "book.html", struct {
		Index  float64
		Author string
		Title  string
		Theme  string
	}{
		Index:  0,
		Author: "Roger Kirkness",
		Title:  "Book Title",
		Theme:  "Theme",
	})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	index := r.FormValue("index")
	author := r.FormValue("author")
	title := r.FormValue("title")
	theme := r.FormValue("theme")
	fmt.Println(index + author + title + theme)
	redirectBack(w, r)
}

func redirectBack(w http.ResponseWriter, r *http.Request) {
	url := r.Header.Get("Referer")
	http.Redirect(w, r, url, http.StatusFound)
}
