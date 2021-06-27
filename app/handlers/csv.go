package handlers

import (
	csvReader "encoding/csv"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/erickertz/csv-exercise/app/models"
	"github.com/erickertz/csv-exercise/app/services"
	"github.com/erickertz/csv-exercise/app/validators"
	"github.com/sirupsen/logrus"
)

// CSV provides methods to process CSv files
type CSV struct {
	Logger              *logrus.Entry
	Validator           *validators.CSV
	CloudStorageService *services.CloudStorage
	SkipCSVHeader       bool
}

// Handle handles CSV files
func (csv *CSV) Handle(sourceBucketName string, csvName string) (string, error) {
	validatorErr := csv.Validator.ValidateFileName(csvName)
	if nil != validatorErr {
		return "", validatorErr
	}
	sourceBucket := csv.CloudStorageService.GetBucket(sourceBucketName)
	sourceObject := csv.CloudStorageService.GetObject(sourceBucket, csvName)
	sourceObjectString, sourceObjectStringErr := csv.CloudStorageService.ReadObjectAsString(sourceObject)
	if nil != sourceObjectStringErr {
		return "", sourceObjectStringErr
	}
	sourceObjectJSON, sourceObjectJSONErr := csv.parseObjectString(sourceObjectString)
	if nil != sourceObjectJSONErr {
		return "", sourceObjectJSONErr
	}
	sourceObjectJSONString, sourceObjectJSONStringMarshalErr := json.Marshal(sourceObjectJSON)
	if nil != sourceObjectJSONStringMarshalErr {
		return "", sourceObjectJSONStringMarshalErr
	}
	return string(sourceObjectJSONString), nil
}

// parseObjectString parses source bucket CSV and returns as JSON
func (csv *CSV) parseObjectString(sourceObjectString string) ([]*models.JSONRecord, error) {
	var jsonRecords []*models.JSONRecord
	fileReader := csvReader.NewReader(strings.NewReader(sourceObjectString))
	records, recordsErr := fileReader.ReadAll()
	if nil != recordsErr {
		return nil, recordsErr
	}
	for index, record := range records {
		if csv.SkipCSVHeader && index == 0 {
			csv.Logger.Debugf("skipping csv header")
			continue
		}
		validateRecordErr := csv.Validator.ValidateRecord(record)
		if nil != validateRecordErr {
			csv.Logger.Warnf("invalid CSV record: %v", validateRecordErr)
			continue
		}
		IDInt, IDIntErr := strconv.Atoi(record[0])
		// All data should be validated at this point so error should never happen, but we shouldn't ignore either way
		if nil != IDIntErr {
			csv.Logger.Warnf("invalid CSV record ID: %v", IDIntErr)
			continue
		}
		jsonRecord := models.JSONRecord{
			ID: IDInt,
			Name: models.JSONRecordName{
				First:  record[1],
				Middle: record[2],
				Last:   record[3],
			},
			Phone: record[4],
		}
		jsonRecords = append(jsonRecords, &jsonRecord)
	}
	return jsonRecords, nil
}
