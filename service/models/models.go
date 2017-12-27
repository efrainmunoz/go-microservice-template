package models

import (
	"gopkg.in/mgo.v2/bson"
)

// User struct that represent a given user
type User struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"id"`
	FirstName    string        `json:"firstname"`
	LastName     string        `json:"lastname"`
	Email        string        `json:"email"`
	Password     string        `json:"password,omitempty"`
	HashPassword []byte        `json:"hashpassword,omitempty"`
}
