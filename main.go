// Package scoircsvjson contains a Google Cloud Storage Cloud Function that converts and stores a csv file to a json file.
package scoircsvjson

import (
	"context"
	"time"

	"google.golang.org/api/option"

	"github.com/erickertz/csv-exercise/app/validators"

	"github.com/erickertz/csv-exercise/app/handlers"
	"github.com/erickertz/csv-exercise/app/services"

	"cloud.google.com/go/storage"
	"github.com/erickertz/csv-exercise/app/models"
	"github.com/sirupsen/logrus"
)

// Main recieves PubSub messages when a new file is uploaded in an input storage bucket, validates the file is a CSV, converts to JSON, and uploads to an output bucket.
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
			logrus.FieldKeyFile:  "sourceLocation.file",
			logrus.FieldKeyFunc:  "sourceLocation.function",
		},
	}
	logrus.SetFormatter(formatter)
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)
	logger := logrus.WithFields(
		logrus.Fields{
			"service":     processedConfiguration.Common.AppName,
			"environment": processedConfiguration.Common.Environment,
			"input_file":  e.Name,
		},
	)

	// Set Cloud Storage Client
	var cloudStorageClient *storage.Client
	var cloudStorageClientErr error
	if processedConfiguration.GCP.GcpApplicationCredentials != "" {
		clientOption := option.WithCredentialsJSON([]byte(processedConfiguration.GCP.GcpApplicationCredentials))
		cloudStorageClient, cloudStorageClientErr = storage.NewClient(context.Background(), clientOption)
	} else {
		cloudStorageClient, cloudStorageClientErr = storage.NewClient(context.Background())
	}
	if cloudStorageClientErr != nil {
		logger.Fatalf("Failed to create cloud storage client: %v", cloudStorageClientErr)
	}

	// Set Cloud Storage Service
	cloudStorageService := services.CloudStorage{
		CloudStorageClient: cloudStorageClient,
	}

	// Set CSV Validator
	csvValidator := validators.CSV{
		Logger:       logger,
		MinRecordLen: processedConfiguration.MinCSVRecordLen,
	}

	// Set CSV Handler
	csvHandler := handlers.CSV{
		Logger:              logger,
		Validator:           &csvValidator,
		CloudStorageService: &cloudStorageService,
		SkipCSVHeader:       processedConfiguration.SkipCSVHeader,
	}

	// Set JSON Handler
	jsonHandler := handlers.JSON{
		CloudStorageService: &cloudStorageService,
	}

	processingStartTime := time.Now()
	logger.Infof("begin processing %v file starting at: %v", e.Name, processingStartTime)

	// Handle CSV Records
	jsonRecords, csvHandlerErr := csvHandler.Handle(processedConfiguration.GCP.GcpInputBucketName, e.Name)
	if nil != csvHandlerErr {
		logger.Fatalf("error handiling CSV file: %v", csvHandlerErr)
	}

	logger.Debugf("processing file: %v with records: %v", e.Name, jsonRecords)

	// Handle JSON Records
	jsonHandlerErr := jsonHandler.Handle(processedConfiguration.GCP.GcpOutputBucketName, e.Name, jsonRecords)
	if nil != jsonHandlerErr {
		logger.Fatalf("error handiling JSON file: %v", jsonHandlerErr)
	}

	processingEndTime := time.Now()
	logger.Infof("end processing %v file starting at: %v, total processing seconds: %v", e.Name, processingEndTime, processingEndTime.Unix()-processingStartTime.Unix())

	return nil

}
