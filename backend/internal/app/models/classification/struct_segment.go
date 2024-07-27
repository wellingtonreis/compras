package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subcategory struct {
	Name string `bson:"name" json:"name"`
}

type Category struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name          string             `bson:"name" json:"name"`
	Subcategories []Subcategory      `bson:"subcategories" json:"subcategories"`
}
