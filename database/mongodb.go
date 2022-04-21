package database

import (
	"backman/structs"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB(MongoURL string) *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoURL))
	if err != nil {
		log.Println("mongodb line:14")
		log.Fatal(err)
	}
	//check if connection sucess
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Panicln("mongodb line:20")
		log.Fatal(err)
	}
	log.Println("connection sucess!")
	return client
}

var DB *mongo.Database

func GetCollection(db *mongo.Client, DBName string, CollectionName string) *mongo.Collection {
	col := db.Database(DBName).Collection(CollectionName)
	return col
}

func InsertDocumentIfDoesntExist[T structs.EpisodeMetadata | map[string]string | structs.MovieMetadata](Collection *mongo.Collection, doc T, uniqueURL string) error {
	upsert := options.Update().SetUpsert(true)
	filter := bson.D{{"url", uniqueURL}}
	d := bson.M{"$set": doc}
	res, err := Collection.UpdateOne(context.TODO(), filter, d, upsert)
	log.Println("mongo action")
	log.Println(res)
	return err
}
