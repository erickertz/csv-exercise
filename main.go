// Package scoircsvjson contains a Google Cloud Storage Cloud Function that converts and stores a csv file to a json file.
package scoircsvjson

import (
	"context"

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

	logger.Infof("Processing file: %s", e.Name)
	return nil

}
