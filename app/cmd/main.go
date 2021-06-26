package main

import (
	"time"

	scoircsvjson "github.com/erickertz/csv-exercise"

	"github.com/erickertz/csv-exercise/app/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func main() {

	context := context.Background()

	processedConfiguration, processedConfigurationErr := scoircsvjson.GetProcessedConfiguration()
	if nil != processedConfigurationErr {
		logrus.Fatalf("error getting proccessed configuration: %v", processedConfigurationErr)
	}

	GCSEvent := models.GCSEvent{
		Kind:                    "",
		ID:                      "",
		SelfLink:                "",
		Name:                    processedConfiguration.TestCSVName,
		Bucket:                  "",
		Generation:              "",
		Metageneration:          "",
		ContentType:             "",
		TimeCreated:             time.Time{},
		Updated:                 time.Time{},
		TemporaryHold:           false,
		EventBasedHold:          false,
		RetentionExpirationTime: time.Time{},
		StorageClass:            "",
		TimeStorageClassUpdated: time.Time{},
		Size:                    "",
		MD5Hash:                 "",
		MediaLink:               "",
		ContentEncoding:         "",
		ContentDisposition:      "",
		CacheControl:            "",
		Metadata:                nil,
		CRC32C:                  "",
		ComponentCount:          0,
		Etag:                    "",
		CustomerEncryption: struct {
			EncryptionAlgorithm string `json:"encryptionAlgorithm"`
			KeySha256           string `json:"keySha256"`
		}{},
		KMSKeyName:    "",
		ResourceState: "",
	}

	err := scoircsvjson.Main(context, GCSEvent)
	if nil != err {
		logrus.Fatalf("error processing GCSEvent: %v", err)
	}

}
