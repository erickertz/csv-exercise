// Package scoircsvjson contains a Google Cloud Storage Cloud Function that converts and stores a csv file to a json file.
package scoircsvjson

import (
	"context"
	"time"

	"github.com/erickertz/csv-exercise/app/validators"

	"github.com/erickertz/csv-exercise/app/handlers"
	"github.com/erickertz/csv-exercise/app/services"

	"cloud.google.com/go/storage"
	"github.com/erickertz/csv-exercise/app/models"
	"github.com/sirupsen/logrus"
)

// MainGCS prints a message when a file is changed in a Cloud Storage bucket.
func Main(ctx context.Context, e models.GCSEvent) error {

	// Set App Config
	processedConfiguration, processedConfigurationErr := GetProcessedConfiguration()
	if nil != processedConfigurationErr {
		logrus.Fatalf("error getting proccessed configuration: %v", processedConfigurationErr)
	}

	// Set logger
	level, parseLevelErr := logrus.ParseLevel(processedConfiguration.Common.LogLevel)
	if parseLevelErr != nil {
		logrus.Fatalf("error setting logging level: %v", processedConfigurationErr)
	}
	formatter := &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "severity",
		},
	}
	logrus.SetFormatter(formatter)
	logrus.SetLevel(level)
	logger := logrus.WithFields(
		logrus.Fields{
			"service":     processedConfiguration.Common.AppName,
			"environment": processedConfiguration.Common.Environment,
		},
	)

	// Set Cloud Storage Client
	cloudStorageClient, cloudStorageClientErr := storage.NewClient(context.Background())
	if cloudStorageClientErr != nil {
		logger.Fatalf("Failed to create cloud storage client: %v", cloudStorageClientErr)
	}

	// Set Cloud Storage Service
	cloudStorageService := services.CloudStorage{
		CloudStorageClient:    cloudStorageClient,
		SourceBucketName:      processedConfiguration.GCP.GcpInputBucketName,
		DestinationBucketName: processedConfiguration.GCP.GcpOutputBucketName,
	}

	processingStartTime := time.Now()
	logger.Infof("begin processing %v file starting at: %v", e.Name, processingStartTime)

	// Set CSV Validator
	csvValidator := validators.CSV{
		Logger: logger,
	}

	// Set CSV Handler
	csvHandler := handlers.CSV{
		Logger:              logger,
		Validator:           &csvValidator,
		CloudStorageService: &cloudStorageService,
	}

	// Handle CSV Records
	jsonRecords, csvHandlerErr := csvHandler.Handle(e.Name)
	if nil != csvHandlerErr {
		logger.Fatalf("error handiling CSV file: %v", csvHandlerErr)
	}

	// Set JSON Handler
	jsonHandler := handlers.JSON{
		CloudStorageService: &cloudStorageService,
	}

	// Handle CSV Records
	jsonHandlerErr := jsonHandler.Handle(jsonRecords)
	if nil != jsonHandlerErr {
		logger.Fatalf("error handiling JSON file: %v", jsonHandlerErr)
	}

	processingEndTime := time.Now()
	logger.Infof("end processing %v file starting at: %v, total processing seconds: %v", e.Name, processingEndTime, processingEndTime.Unix()-processingStartTime.Unix())

	return nil

}
