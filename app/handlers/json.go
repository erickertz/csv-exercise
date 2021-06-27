package handlers

import (
	"strings"

	"github.com/erickertz/csv-exercise/app/services"
)

// JSON provides methods to process JSON files
type JSON struct {
	CloudStorageService *services.CloudStorage
}

// Handle handles JSON files
func (json *JSON) Handle(destinationBucketName string, destinationFileName string, jsonRecords string) error {
	destinationFileName = strings.Replace(destinationFileName, ".csv", ".json", -1)
	destinationBucket := json.CloudStorageService.GetBucket(destinationBucketName)
	destinationObject := json.CloudStorageService.GetObject(destinationBucket, destinationFileName)
	return json.CloudStorageService.UploadObject(destinationObject, "application/json", nil, jsonRecords)
}
