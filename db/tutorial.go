package db

import (
	"context"
	"errors"
	"log"

	"github.com/tarathep/go-server-crud/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

//TutorialRepository is contain interface
type TutorialRepository interface {
	Create(tutorial model.Tutorial) error
	FindAll(title string) ([]*model.Tutorial, error)
	FindOne(id uuid.UUID) (model.Tutorial, error)
	Update(id uuid.UUID) error
	Delete(id uuid.UUID) error
	DeleteAll() error
	FindAllPublished() ([]model.Tutorial, error)
}

func (db *MongoDB) Create(tutorial model.Tutorial) error {
	collection := db.Database("bokie").Collection("tutorials")
	_, err := collection.InsertOne(context.TODO(), tutorial)
	if err != nil {
		return err
	}
	return nil
}

func (db *MongoDB) FindAll(title string) ([]*model.Tutorial, error) {

	collection := db.Database("bokie").Collection("tutorials")

	findOptions := options.Find()

	var results []*model.Tutorial

	// filter := bson.D{{Key: "foo", Value: 99}}
	filter := bson.M{"title": bson.M{"$regex": title}}

	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem model.Tutorial
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

func (db *MongoDB) FindOne(id uuid.UUID) (model.Tutorial, error) {
	var tutorial model.Tutorial
	return tutorial, errors.New("err")
}

func (db *MongoDB) Update(id uuid.UUID) error {

	return errors.New("err")
}

func (db *MongoDB) Delete(id uuid.UUID) error {

	return errors.New("err")
}

func (db *MongoDB) DeleteAll() error {

	return errors.New("err")
}

func (db *MongoDB) FindAllPublished() ([]model.Tutorial, error) {

	var tutorials []model.Tutorial
	return tutorials, errors.New("err")
}
