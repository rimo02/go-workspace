package models

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Role string             `json:"role" bson:"role"`
	jwt.RegisteredClaims
}
