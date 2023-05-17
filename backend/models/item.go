package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct{
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
}

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	PhotoUrl    string             `bson:"photo_url" json:"photo_url"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price"`
	Category    Category           `bson:"category" json:"category"`
}
