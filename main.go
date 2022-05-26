package main

import (
	"backman/control"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
	log.Print(S3ID, S3Key, S3Bucket, MongoURL)
	app := mux.NewRouter()
	app.HandleFunc("/api/v1/upload", control.Handle).Methods("POST")
	app.HandleFunc("/api/v1/redis/bf/exist", control.RedisHandleGet).Methods("GET")
	app.HandleFunc("/api/v1/gorm/pg/schemas", control.AllSchemaHandle).Methods("GET")
	app.HandleFunc("/api/v1/redis/bf/put", control.RedisHandlePut).Methods("PUT")
	app.HandleFunc("/api/v1/gorm/pg/schema/create", control.CreateSchemaHandle)

	http.ListenAndServe(":9090", app)

}
