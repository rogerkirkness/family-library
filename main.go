package main

import (
	"fmt"
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Book represents one physical book in the library.
type Book struct {
	gorm.Model
	Index  int    `gorm:"index:number"`
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

func main() {
	router := gin.Default()
	db, err := gorm.Open("sqlite3", "dev.db")
	if err != nil {
		panic("Failed to connect to db")
	}
	db.AutoMigrate(&Book{})
	defer db.Close()
	var views = template.Must(template.ParseGlob("views/*.html"))
	router.SetHTMLTemplate(views)
	router.GET("/books", func(c *gin.Context) {
		var books []Book
		db.Find(&books)
		c.HTML(200, "books.html", books)
	})
	router.GET("/books/:id", func(c *gin.Context) {
		var book Book
		db.Find(&book, c.Param("id"))
		c.JSON(200, book)
	})
	router.POST("/books", func(c *gin.Context) {
		author := c.PostForm("author")
		title := c.PostForm("title")
		theme := c.PostForm("theme")
		index, _ := strconv.Atoi(c.PostForm("index"))
		id := c.PostForm("id")
		if id != "" {
			var book Book
			method := c.PostForm("method")
			fmt.Println(method)
			db.First(&book, id)
			if method != "" {
				db.Delete(&book)
				c.Redirect(302, "/books")
			} else {
				book.Index = index
				book.Author = author
				book.Title = title
				book.Theme = theme
				db.Save(&book)
				c.Redirect(302, "/books")
			}
		} else {
			db.Create(&Book{
				Index:  index,
				Author: author,
				Title:  title,
				Theme:  theme,
			})
			c.Redirect(302, "/books")
		}
	})
	router.Run()
}
