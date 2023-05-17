package models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	IsAdmin	  bool               `bson:"is_admin" json:"is_admin"`
}

type UserDto struct {
	Username	string	`json:"username" validate:"required"`
	Password	string	`json:"password" validate:"required"`
}