package main

import (
	"backman/control"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("failed to parse .env")
		log.Fatal(err)
	}
	S3ID := os.Getenv("S3ID")
	S3Key := os.Getenv("S3KEY")
	S3Bucket := os.Getenv("S3Bucket")
	MongoURL := os.Getenv("MongodbUrl")
	control.Initial(control.S3Credential{S3ID, S3Key, S3Bucket}, MongoURL)
	// log.Print(S3ID, S3Key, S3Bucket, MongoURL)
	app := fiber.New()
	app.Post("/upload", control.Handle)
	app.Listen(":9090")

}
