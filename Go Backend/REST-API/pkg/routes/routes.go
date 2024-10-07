package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rimo02/rest-api/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *gin.Engine) {
	router.POST("/book/", controllers.CreateBook)          //create and write
	router.GET("/books/", controllers.GetBook)             // read
	router.PUT("/book/:bookId", controllers.UpdateBook)    // update
	router.GET("/books/:bookId", controllers.GetBookById)  // read
	router.DELETE("/book/:bookId", controllers.DeleteBook) // delete
}
