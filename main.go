package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Book represents one physical book in the library.
type Book struct {
	gorm.Model
	Index  int    `gorm:"unique;not null"`
	Author string `gorm:"index:author"`
	Title  string `gorm:"index:title"`
	Theme  string `gorm:"index:theme"`
}

// Publication year number
// Year acquired number
// Roger first read year number
// Jessica first read year number
// Borrowed true/false bool
// Borrowed by (name) string
// Borrowed date

var views = template.Must(template.ParseGlob("views/*.html"))

func main() {
	db, err := gorm.Open("sqlite3", "dev.db")
	if err != nil {
		panic("Failed to connect to db")
	}
	defer db.Close()
	db.AutoMigrate(&Book{})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			var books []Book
			db.Find(&books)
			w.Header().Set("Content-Type", "text/html")
			views.ExecuteTemplate(w, "books.html", books)
		case "POST":
		case "PUT":
		case "DELETE":
		}
	})
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
		case "POST":
			index, _ := strconv.Atoi(r.FormValue("index"))
			db.Create(&Book{
				Index:  index,
				Author: r.FormValue("author"),
				Title:  r.FormValue("title"),
				Theme:  r.FormValue("theme"),
			})
			redirectBack(w, r)
		case "PUT":
		case "DELETE":
		}
	})
	serverStartError := http.ListenAndServe(":3000", nil)
	if serverStartError != nil {
		log.Fatalf("error listening: %s", serverStartError)
	}
}

func redirectBack(w http.ResponseWriter, r *http.Request) {
	url := r.Header.Get("Referer")
	http.Redirect(w, r, url, http.StatusFound)
}
