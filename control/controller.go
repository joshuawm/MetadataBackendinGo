package control

import (
	"backman/database"
	"backman/database/gromimplement"
	"backman/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type S3Credential struct {
	ID     string
	Key    string
	Bucket string
}

var MongoClient *mongo.Client
var S3Session *session.Session
var S3Uploader *s3manager.Uploader
var bucketName string
var contents chan structs.UploadInterface = make(chan structs.UploadInterface, 10)
var BFilterMan database.BFilter = database.BFilter{}
var GormDB *gorm.DB

func Initial(S3Info S3Credential, MongoURL string) {
	MongoClient = database.ConnectDB(MongoURL)
	S3Session = database.ConnectS3(S3Info.ID, S3Info.Key, "s3.us-west-002.backblazeb2.com", "us-west-002")
	S3Uploader = database.CreateUploader(S3Session)
	GormDB = gromimplement.ConnectDBCockroach()
	BFilterMan.Create("redis-13493.c54.ap-northeast-1-2.ec2.cloud.redislabs.com:13493", "E4z3Z8JABf3eFrAhvEHeR1cd8FC5z9eE", 0, "links")
	go func() {
		process(contents, MongoClient)
	}()
}

func process(Structdata <-chan structs.UploadInterface, client *mongo.Client) {
	for d := range Structdata {
		log.Println("received")
		err := GormHandler(d)
		if err != nil {
			log.Panicln(fmt.Sprintf("got an error: %s", err.Error()))
		}
		if d.EpMeta.URl != "" {
			database.InsertDocumentIfDoesntExist(database.GetCollection(MongoClient, d.Name, "episode"), d.EpMeta, d.Name)
		} else if d.MoMeta.URl != "" {
			database.InsertDocumentIfDoesntExist(database.GetCollection(MongoClient, d.Name, "movie"), d.MoMeta, d.Name)
		}
		MediauplodS3(d.Media, d.Name, S3Uploader)

	}
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var data structs.UploadInterface
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("failed to parse json"))

	}

	// contents <- data
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("received"))

}
