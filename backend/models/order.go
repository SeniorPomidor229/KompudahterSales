package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	UserID     primitive.ObjectID   `bson:"userId" json:"userId"`
	TotalPrice float64              `bson:"totalPrice" json:"totalPrice"`
	Products   []Product 			`bson:"products" json:"products"`
	CreatedAt  time.Time            `bson:"createdAt" json:"createdAt"`
}