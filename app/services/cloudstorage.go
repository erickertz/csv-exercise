package services

import (
	"cloud.google.com/go/storage"
)


// CloudStorage provides methods to interact with Google Cloud Storage
type CloudStorage struct {
	CloudStorageClient    *storage.Client
	SourceBucketName      string
	DestinationBucketName string
}
