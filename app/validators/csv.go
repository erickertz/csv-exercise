package validators

import (
	"regexp"
	"strconv"
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
	validateRecordInternalIDErr := csv.validateRecordInternalID(record[0], true, 8)
	if nil != validateRecordInternalIDErr {
		return validateRecordInternalIDErr
	}
	validateRecordFirstNameErr := csv.validateRecordFirstName(record[1], true, 15)
	if nil != validateRecordFirstNameErr {
		return validateRecordFirstNameErr
	}
	validateRecordMiddleNameErr := csv.validateRecordMiddleName(record[2], false, 15)
	if nil != validateRecordMiddleNameErr {
		return validateRecordMiddleNameErr
	}
	validateRecordLastNameErr := csv.validateRecordLastName(record[3], true, 15)
	if nil != validateRecordLastNameErr {
		return validateRecordLastNameErr
	}
	validateRecordPhoneNumErr := csv.validateRecordPhoneNum(record[4], true)
	if nil != validateRecordPhoneNumErr {
		return validateRecordPhoneNumErr
	}
	return nil
}

func (csv *CSV) validateRecordInternalID(record string, required bool, fieldLen int) error {
	fieldName := "INTERNAL_ID"
	if required && csv.isEmpty(record) {
		return errors2.Newf("%v cannot be empty", fieldName)
	}
	if len(record) != fieldLen {
		return errors2.Newf("%v must be exactly %v characters", fieldName, fieldLen)
	}
	if !csv.isPositiveInt(record) {
		return errors2.Newf("%v must be a positive integer", fieldName)
	}
	return nil
}

func (csv *CSV) validateRecordFirstName(record string, required bool, maxLen int) error {
	fieldName := "FIRST_NAME"
	if required && csv.isEmpty(record) {
		return errors2.Newf("%v cannot be empty", fieldName)
	}
	if csv.isMaxLen(record, maxLen) {
		return errors2.Newf("%v can not exceed %v characters", fieldName, maxLen)
	}
	return nil
}

func (csv *CSV) validateRecordMiddleName(record string, required bool, maxLen int) error {
	fieldName := "MIDDLE_NAME"
	if required && csv.isEmpty(record) {
		return errors2.Newf("%v cannot be empty", fieldName)
	}
	if csv.isMaxLen(record, maxLen) {
		return errors2.Newf("%v can not exceed %v characters", fieldName, maxLen)
	}
	return nil
}

func (csv *CSV) validateRecordLastName(record string, required bool, maxLen int) error {
	fieldName := "LAST_NAME"
	if required && csv.isEmpty(record) {
		return errors2.Newf("%v cannot be empty", fieldName)
	}
	if csv.isMaxLen(record, maxLen) {
		return errors2.Newf("%v can not exceed %v characters", fieldName, maxLen)
	}
	return nil
}

func (csv *CSV) validateRecordPhoneNum(record string, required bool) error {
	fieldName := "PHONE_NUM"
	pattern := "^\\d{3}\\-\\d{3}\\-\\d{4}$"
	if required && csv.isEmpty(record) {
		return errors2.Newf("%v cannot be empty", fieldName)
	}
	matched, matchedErr := regexp.Match(pattern, []byte(record))
	if nil != matchedErr {
		return errors2.Newf("% validation error %v: %v", fieldName, record, matchedErr)
	}
	if !matched {
		return errors2.Newf("%v %v does not match pattern: %v", fieldName, record, pattern)
	}
	return nil
}

func (csv *CSV) isEmpty(record string) bool {
	return len(record) == 0
}

func (csv *CSV) isMaxLen(record string, maxLen int) bool {
	return len(record) > maxLen
}

func (csv *CSV) isPositiveInt(record string) bool {
	recordInt, err := strconv.Atoi(record)
	if nil == err && recordInt >= 0 {
		return true
	}
	return false
}
