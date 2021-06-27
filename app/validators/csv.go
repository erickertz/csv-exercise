package validators

import (
	"strings"

	"github.com/sirupsen/logrus"
	errors2 "gopkg.in/errgo.v2/fmt/errors"
)

// CSV provides methods to validate CSV files
type CSV struct {
	Logger       *logrus.Entry
	MinRecordLen int
}

// ValidateFile validates a CSV file name
func (csv *CSV) ValidateFileName(fileName string) error {
	CSVSuffix := ".csv"
	if !strings.HasSuffix(fileName, CSVSuffix) {
		return errors2.Newf("csv file name must end with: %v", CSVSuffix)
	}
	return nil
}

// ValidateRecord validates a CSV record
func (csv *CSV) ValidateRecord(record []string) error {
	recordLen := len(record)
	if recordLen < csv.MinRecordLen {
		return errors2.Newf("csv record has %v number values, minimal is %v", recordLen, csv.MinRecordLen)
	}
	return nil
}
