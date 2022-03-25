package control

import (
	"backman/database"
	"backman/structs"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type S3Credential struct {
	ID     string
	Key    string
	Bucket string
}

var MongoClient *mongo.Client
var S3Session *session.Session
var S3Uploader *s3manager.Uploader
var contents chan structs.UploadInterface = make(chan structs.UploadInterface, 10)

func Initial(S3Info S3Credential, MongoURL string) {
	MongoClient = database.ConnectDB(MongoURL)
	S3Session = database.ConnectS3(S3Info.ID, S3Info.Key, "s3.us-west-002.backblazeb2.com", "us-west-002")
	S3Uploader = database.CreateUploader(S3Session)
	go func() {
		process(contents, MongoClient)
	}()
}

func process(Structdata <-chan structs.UploadInterface, client *mongo.Client) {
	for d := range Structdata {
		log.Println("received")
		if d.EpMeta.URl != "" {
			database.InsertDocumentIfDoesntExist(database.GetCollection(MongoClient, d.Name, "episode"), d.EpMeta)
		} else if d.MoMeta.URl != "" {
			database.InsertDocumentIfDoesntExist(database.GetCollection(MongoClient, d.Name, "movie"), d.MoMeta)
		}
		MediauplodS3(d.Media, "testfolder", S3Uploader)

	}
}

func Handle(c *fiber.Ctx) error {
	var data structs.UploadInterface
	c.BodyParser(&data)
	contents <- data
	return c.Send([]byte("dsdsdsdsd"))

}
