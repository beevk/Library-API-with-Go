package handlers

import (
	"BOOK-API/models"
	"BOOK-API/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, utils.Books)
}

func GetSingleBook(c *gin.Context) {
	id := c.Param("id")
	book, err := utils.GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	indexToRemove, err := utils.GetBookIndexById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if indexToRemove >= 0 {
		utils.Books = append(utils.Books[:indexToRemove], utils.Books[indexToRemove+1:]...)
	}
	c.IndentedJSON(http.StatusOK, utils.Books)
}

func CreateBook(c *gin.Context) {
	var newBook models.Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBook.ID = strconv.Itoa(len(utils.Books) + 1)
	fmt.Println(newBook)

	utils.Books = append(utils.Books, newBook)

	c.IndentedJSON(http.StatusCreated, newBook)
}

func RentABook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Missing 'id' query parameter."})
		return
	}

	book, err := utils.GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not available. Please try later."})
		return
	}

	book.Quantity -= 1

	c.IndentedJSON(http.StatusOK, book)
}

func ReturnABook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "missing id query parameter."})
		return
	}

	book, err := utils.GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	book.Quantity += 1

	c.IndentedJSON(http.StatusOK, book)
}
