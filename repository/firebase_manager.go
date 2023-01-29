package repository

import (
	"context"
	"os"
)

type ReaderWriterStorageBucket interface {
	UploadToStorageBucket(ctx context.Context, objectname string, file *os.File) error
	GetFileDownloadUrl(ctx context.Context, objectname string) (string, error)
}
