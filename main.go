package main

import (
	_ "net/http"

	"github.com/gin-gonic/gin"

	"bookshelf/controllers"
	"bookshelf/models"
)

func main() {
	r := gin.Default()

	// r.GET("/", func(c *gin.Context){
	// 	c.JSON(http.StatusOK, gin.H{"data": "hello gin web server"})
	// })

	models.ConnectToDatabase()

	r.GET("/books", controllers.FindAllBook)
	r.GET("/books/:id", controllers.FindBookById)
	r.POST("/books", controllers.CreateNewBook)

	r.Run(":8081")
}
