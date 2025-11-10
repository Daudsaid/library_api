package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"` // <- string
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien", Quantity: 3},
	{ID: "2", Title: "The Hobbit", Author: "J. R. R. Tolkien", Quantity: 2},
	{ID: "3", Title: "The Silmarillion", Author: "J. R. R. Tolkien", Quantity: 1},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookBy(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id not found"})
		return
	}
	book, err := getBookBy(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	book.Quantity++
	c.IndentedJSON(http.StatusOK, book)
}

func getBookBy(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id not found"})
		return
	}
	book, err := getBookBy(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not available"})
		return
	}
	book.Quantity--
	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	r := gin.Default()
	_ = r.SetTrustedProxies(nil)
	r.GET("/books", getBooks)
	r.GET("/books/:id", bookByID)
	r.PATCH("/checkout", checkoutBook)
	r.PATCH("/return", returnBook)
	r.POST("/books", createBook)
	r.Run(":1880")
}
