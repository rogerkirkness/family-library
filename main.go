package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Book represents phsyical library book.
type Book struct {
	gorm.Model
	Index  string `form:"Index" json:"Index" gorm:"index:number"`
	Author string `form:"Author" json:"Author" gorm:"index:author"`
	Title  string `form:"Title" json:"Title" gorm:"index:title"`
	Theme  string `form:"Theme" json:"Theme" gorm:"index:theme"`
}

func main() {
	db, err := gorm.Open("sqlite3", "dev.db")
	if err != nil {
		panic("Failed to connect to db")
	}
	router := gin.Default()
	db.AutoMigrate(&Book{})
	defer db.Close()
	var views = template.Must(template.ParseGlob("views/*.html"))
	router.SetHTMLTemplate(views)
	router.GET("/books", func(c *gin.Context) {
		var books []Book
		db.Find(&books)
		contentType := c.GetHeader("Content-Type")
		if contentType == "application/json" {
			c.JSON(200, books)
		} else {
			c.HTML(200, "books", books)
		}
	})
	router.GET("/books/:id", func(c *gin.Context) {
		var book Book
		db.Find(&book, c.Param("id"))
		if book.ID != 0 {
			contentType := c.GetHeader("Content-Type")
			if contentType == "application/json" {
				c.JSON(200, book)
			} else {
				c.HTML(200, "book", book)
			}
		} else {
			c.JSON(400, gin.H{"id": c.Param("id"), "error": "not found"})
		}
	})
	router.POST("/books", func(c *gin.Context) {
		var book Book
		err := c.ShouldBind(&book)
		if err != nil {
			c.JSON(500, gin.H{"error": err})
		} else {
			db.Create(&book)
			contentType := c.GetHeader("Content-Type")
			if contentType == "application/json" {
				c.JSON(200, gin.H{"id": book.ID, "status": "created"})
			} else {
				c.Redirect(302, "http://localhost:8080/books")
			}
		}
	})
	router.PUT("/books/:id", func(c *gin.Context) {
		var book Book
		db.Find(&book, c.Param("id"))
		if book.ID != 0 {
			err := c.ShouldBind(&book)
			if err != nil {
				log.Println(err)
			} else {
				db.Save(&book)
				c.JSON(200, gin.H{"id": book.ID, "status": "updated"})
			}
		} else {
			c.JSON(400, gin.H{"id": c.Param("id"), "error": "not found"})
		}
	})
	router.DELETE("/books/:id", func(c *gin.Context) {
		var book Book
		db.Find(&book, c.Param("id"))
		if book.ID != 0 {
			db.Delete(&book)
			c.JSON(200, gin.H{"id": book.ID, "status": "deleted"})
		} else {
			c.JSON(400, gin.H{"id": c.Param("id"), "error": "not found"})
		}
	})
	router.Run()
}
