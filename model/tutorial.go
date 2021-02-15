package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tutorial struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Published   bool               `json:"published"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}

// EXAMPLE OUTPUT JSON
// {
// 	"title": "Node Tut #3",
// 	"description": "Tut#3 Description",
// 	"published": false,
// 	"createdAt": "2021-02-13T06:34:28.906Z",
// 	"updatedAt": "2021-02-13T06:34:28.906Z",
// 	"id": "602772f4a9ee96155cbde9b0"
// },
// {
// 	"title": "NodeJS Basic",
// 	"description": "BasicNodejs Description",
// 	"published": false,
// 	"createdAt": "2021-02-13T06:36:15.753Z",
// 	"updatedAt": "2021-02-13T06:36:15.753Z",
// 	"id": "6027735fa9ee96155cbde9b1"
// }
