package firebase

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/aditya37/geospatial-tracking/repository"

	gofbs "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/storage"
)

type storageBucket struct {
	storage *storage.Client
}

func NewStorageBucket(ctx context.Context, app *gofbs.App) (repository.ReaderWriterStorageBucket, error) {
	storageClient, err := app.Storage(ctx)
	if err != nil {
		return nil, err
	}

	return &storageBucket{
		storage: storageClient,
	}, nil
}

func (sb *storageBucket) UploadToStorageBucket(ctx context.Context, objectname string, file *os.File) error {
	bucket, err := sb.storage.DefaultBucket()
	if err != nil {
		return err
	}

	// writer..
	writer := bucket.Object(objectname).NewWriter(ctx)

	// add download token for anonymous access..
	writer.ObjectAttrs.Metadata = map[string]string{
		"firebaseStorageDownloadTokens": objectname,
	}
	// write file to bucket..
	if _, err := io.Copy(writer, file); err != nil {
		return err
	}

	// close writer...
	if err := writer.Close(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

//
func (sb *storageBucket) GetFileDownloadUrl(ctx context.Context, objectname string) (string, error) {
	bc, err := sb.storage.DefaultBucket()
	if err != nil {
		return "", err
	}
	attribute, err := bc.Object(objectname).Attrs(ctx)
	if err != nil {
		return "", err
	}
	return attribute.MediaLink, nil
}
