package handlers

import (
	"github.com/erickertz/csv-exercise/app/models"
	"github.com/erickertz/csv-exercise/app/services"
)


// JSON provides methods to process JSON files
type JSON struct {
	CloudStorageService *services.CloudStorage
}


// Handle handles JSON files
func (json *JSON) Handle(jsonRecords []*models.JSONRecord) error {
	return nil
}
