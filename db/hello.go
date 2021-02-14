package db

import (
	"context"
	"log"

	"github.com/tarathep/go-server-crud/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//HelloRepository is interface for implement
type HelloRepository interface {
	AllHello() ([]*model.Hello, error)
	InsertHello(hello model.Hello) (model.Hello, error)
}

//AllHello get datal list
func (db *MongoDB) AllHello() ([]*model.Hello, error) {

	collection := db.Database("test").Collection("hello")

	findOptions := options.Find()
	//findOptions.SetLimit(100)

	var results []*model.Hello

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		var elem model.Hello
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	return results, nil
}

//InsertHello is insert data into mongo
func (db *MongoDB) InsertHello(hello model.Hello) (model.Hello, error) {
	collection := db.Database("test").Collection("hello")
	_, err := collection.InsertOne(context.TODO(), hello)
	if err != nil {
		return hello, err
	}
	return hello, nil
}
