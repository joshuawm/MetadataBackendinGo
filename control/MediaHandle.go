package control

import (
	"backman/database"
	"backman/structs"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func MediauplodS3(media structs.Media, BaseDirectoryName string, uploader *s3manager.Uploader) {
	if media.Trailer != "" {
		go downloadNupload(media.Thumbnail, fmt.Sprintf("%s/trailer.mp4", BaseDirectoryName), uploader, database.InsertObjectWithS3)
	}
	if media.Poster != "" {
		go downloadNupload(media.Poster, fmt.Sprintf("%s/poster.jpg", BaseDirectoryName), uploader, database.InsertObjectWithS3)
	}
	if media.Thumbnail != "" {
		go downloadNupload(media.Thumbnail, fmt.Sprintf("%s/thumbnail.jpg", BaseDirectoryName), uploader, database.InsertObjectWithS3)
	}
	if len(media.Gallery) > 0 {
		for i, v := range media.Gallery {
			go downloadNupload(v, fmt.Sprintf("%s/gallery/%s.jpg", BaseDirectoryName, fmt.Sprint(i)), uploader, database.InsertObjectWithS3)
		}
	}
}

func downloadNupload(url string, folder string, uploader *s3manager.Uploader, f func(*s3manager.Uploader, string, string, io.Reader) error) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	return f(uploader, bucketName, folder, resp.Body)
}
