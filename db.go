package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

//DBres is a struct for the data returned when searching the database
type DBres struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	StringID    string             `json:"id"`
	Date        string             `json:"date"`
	Description string             `json:"description"`
}

func init() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions) //context.TODO() = returns an empty context

	if err != nil {
		log.Fatal(err)
	}

	// Tests to see if there is a proper connection to the mongodb server
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection = client.Database("gotodo").Collection("todolist")
}

func insertDB(data Data) string {
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return "err"
	}
	return "ok"
}

func getAll() []DBres {

	var result DBres
	var results []DBres

	res, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for res.Next(context.Background()) {
		err := res.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		result.StringID = result.ID.Hex()
		results = append(results, result)
	}
	res.Close(context.TODO())

	return results
}
