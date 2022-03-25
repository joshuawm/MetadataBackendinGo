package database

import (
	"context"
	"log"
	"time"

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
func InsertDocumentIfDoesntExist(Collection *mongo.Collection, doc interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	upsert := options.Update()
	t := true
	upsert.Upsert = &t
	// res, err := Collection.UpdateOne(ctx, doc, doc, upsert)
	res, err := Collection.InsertOne(ctx, doc)
	log.Println("mongo action")
	log.Println(res)
	return err
}
