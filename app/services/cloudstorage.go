package services

import (
	"io"
	"strings"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

// CloudStorage provides methods to interact with Google Cloud Storage
type CloudStorage struct {
	CloudStorageClient *storage.Client
}

// GetBucket gets a Google Cloud Storage bucket
func (cloudStorage *CloudStorage) GetBucket(bucketName string) *storage.BucketHandle {
	return cloudStorage.CloudStorageClient.Bucket(bucketName)
}

// GetObject gets a Google Cloud Storage bucket object
func (cloudStorage *CloudStorage) GetObject(bucket *storage.BucketHandle, objectName string) *storage.ObjectHandle {
	return bucket.Object(objectName)
}

// ReadObjectAsString gets a Google Cloud Storage bucket object contents as string
func (cloudStorage *CloudStorage) ReadObjectAsString(object *storage.ObjectHandle) (string, error) {
	reader, readerErr := object.NewReader(context.Background())
	if nil != readerErr {
		return "", readerErr
	}
	buf := new(strings.Builder)
	_, bufErr := io.Copy(buf, reader)
	if nil != bufErr {
		return "", bufErr
	}
	return buf.String(), nil
}

// UploadObject uploads a Google Cloud Storage object to a bucket
func (cloudStorage *CloudStorage) UploadObject(object *storage.ObjectHandle, contentType string, metadata map[string]string, content string) error {
	writer := object.NewWriter(context.Background())
	writer.ContentType = contentType
	writer.Metadata = metadata
	_, writeErr := writer.Write([]byte(content))
	if nil != writeErr {
		return writeErr
	}
	closeErr := writer.Close()
	if nil != closeErr {
		return closeErr
	}
	return nil
}
