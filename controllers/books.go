package controllers

import (
	"bookshelf/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// For create a new book
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// For update a book
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GET all books
func FindAllBook(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// Create a new book
func CreateNewBook(c *gin.Context) {
	var input CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}
	models.DB.Create(&book)

	c.JSON(http.StatusCreated, gin.H{"data": book})

}

// Find by id
func FindBookById(c *gin.Context){
	var book models.Book

	if err := models.DB.Where("id= ?", c.Param("id")).First(&book).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book Not Found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Update book with id
func UpdateBookById(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book Not Found!"})
		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	models.DB.Model(&book).Updates(input)
	c.JSON(http.StatusAccepted, gin.H{"data": book})
}

// Delete book with id
func DeleteBookById(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book Not Found!"})
		return 
	}

	models.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": "Deleted!"})
}