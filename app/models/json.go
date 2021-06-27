package models

// JSONRecord is the record format of json output
type JSONRecord struct {
	ID    int            `json:"id"`
	Name  JSONRecordName `json:"name"`
	Phone string         `json:"phone"`
}

// JSONRecordName is the Name record format of json output
type JSONRecordName struct {
	First  string `json:"first"`
	Middle string `json:"middle,omitempty"`
	Last   string `json:"last"`
}
