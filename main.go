package main

import (
	"net/http"
	"os"
)

func init() {

}

func main() {
	address := ":" + os.Getenv("PORT")
	http.HandleFunc("/", index)
	http.HandleFunc("/books", books)
	http.HandleFunc("/book/{index}", book)
}
