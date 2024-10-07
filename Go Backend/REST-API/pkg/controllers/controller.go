package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rimo02/rest-api/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// The gin.Context is one of the most critical parts of the Gin framework. It is passed into each handler function and provides everything needed to handle an HTTP request and generate a response. It holds the request and response objects, along with many helper methods to interact with them.

type Book struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title" bson:"title"`
	Author string             `json:"author" bson:"author"`
}

// var books = []Book{
// 	{ID: 1, Title: "A Tale of two", Author: "O'Henry"},
// 	{ID: 2, Title: "The Promised Neverland", Author: "ABC..."},
// }

var books []Book

func CreateBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if newBook.Title == "" || newBook.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no title or author mentioned"})
	}
	collection := database.GetCollection(database.Client, "novels")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newBook.ID = primitive.NewObjectID() //ObjectIds in MongoDB are 12-byte values that uniquely identify documents in a collection.

	result, err := collection.InsertOne(ctx, newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Book created successfully",
		"book":    newBook,
		"id":      result.InsertedID,
	})
}

// In MongoDB, a cursor is an object that references documents that are identified by a query. Cursors are used to return the results of a read operation in batches, instead of all at once. This reduces the number of server requests and memory consumption.
func GetBook(c *gin.Context) {
	c.JSON(http.StatusOK, books)

	collection := database.GetCollection(database.Client, "novels")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var books []Book
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var book Book
		if err := cursor.Decode(&book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, book)
	}
	c.JSON(http.StatusOK, books)

}

func GetBookById(c *gin.Context) {
	bookId := c.Param("bookId")
	// for _, book := range books {
	// 	if book.ID == bookId {
	// 		c.JSON(http.StatusOK, book)
	// 		return
	// 	}
	// }
	collection := database.GetCollection(database.Client, "novels")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var book Book
	objectId, err := primitive.ObjectIDFromHex(bookId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID format"})
		return
	}
	err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&book)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No book found for the given id"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)

}

// func UpdateBook(c *gin.Context) {
// 	bookId := c.Param("bookId")

// 	var updatedBook Book
// 	if err := c.BindJSON(&updatedBook); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	for i, book := range books {
// 		if book.ID == bookId {
// 			if updatedBook.Author != "" {
// 				books[i].Title = updatedBook.Title
// 			}
// 			if updatedBook.Author != "" {
// 				books[i].Author = updatedBook.Author
// 			}
// 			c.JSON(http.StatusOK, books[i])
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
// }

func UpdateBook(c *gin.Context) {
	bookId := c.Param("bookId")
	collection := database.GetCollection(database.Client, "novels")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(bookId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedFields map[string]interface{}

	if err := c.BindJSON(&updatedFields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updated := bson.M{
		"$set": updatedFields,
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, updated)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})

}

func DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")
	collection := database.GetCollection(database.Client, "novels")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(bookId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
