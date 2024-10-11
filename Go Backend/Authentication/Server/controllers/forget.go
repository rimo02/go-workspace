package controllers

import (
	"auth-go/database"
	"auth-go/models"
	"auth-go/utils"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

func ForgetPassword(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var existingUser models.User
	collection := database.GetCollection(database.Client, "user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found for that email"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	hashPass, errHash := utils.GenerateHashPassword(user.Password)
	if errHash != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating password"})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"password": hashPass}}
	_, err = collection.UpdateOne(ctx, bson.M{"email": existingUser.Email}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "password updated"})
	c.Redirect(http.StatusOK, "/login")
}
