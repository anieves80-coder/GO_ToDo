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

func insertDB(data AddData) string {
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

func deleteRec(id string) {
	idPrimitive, _ := primitive.ObjectIDFromHex(id)
	collection.DeleteOne(context.TODO(), bson.M{"_id": idPrimitive})
}

func updateDb(data Data) string {
	idPrimitive, _ := primitive.ObjectIDFromHex(data.ID)
	filter := bson.D{{"_id", idPrimitive}}
	//bson.D = slice
	//bson.M = map
	update := bson.M{
		"$set": bson.M{
			"date":        data.Date,
			"description": data.Description,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "err"
	}

	return "ok"
}
