package handlers

import (
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
}

// Handle handles CSV files
func (csv *CSV) Handle(csvName string) ([]*models.JSONRecord, error) {
	validatorError := csv.Validator.Validate()
	if nil != validatorError {
		return nil, validatorError
	}
	return nil, nil
}
