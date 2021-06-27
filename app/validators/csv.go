package validators

import "github.com/sirupsen/logrus"

// CSV provides methods to validate CSV files
type CSV struct {
	Logger *logrus.Entry
}

// Validate validates CSV files
func (csv *CSV) Validate() error {
	return nil
}
