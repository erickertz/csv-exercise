package models

// CSVRecord is the record format of csv input
type CSVRecord struct {
	INTERNAL_ID int32
	FIRST_NAME  string
	MIDDLE_NAME string
	LAST_NAME   string
	PHONE_NUM   string
}
