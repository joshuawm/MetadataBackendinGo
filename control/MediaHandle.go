package control

import (
	"backman/database"
	"backman/structs"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func MediauplodS3(media structs.Media, FolderName string, uploader *s3manager.Uploader) {
	if media.Gallery != nil {
		for _, value := range media.Gallery {
			log.Println(value)
			resp, err := http.Get(value)
			if err != nil {
				log.Println("failed to get from %s", value)
				continue
			}
			// f, err := os.Create("jin.jpg")
			// if err != nil {
			// 	log.Print(err)
			// 	continue
			// }
			// io.Copy(f, resp.Body)
			// f.Close()
			errs := database.InsertObjectWithS3(uploader, "gmetadata", "man/naked/huhu.jpg", resp.Body)
			log.Println(errs)
		}
	}
}
