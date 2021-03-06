package scoircsvjson

import (
	"github.com/kelseyhightower/envconfig"
)

//Configuration contains fields for the application configuration.
type Configuration struct {
	Common
	GCP
}

//Common contains fields for the common application configuration.
type Common struct {
	AppName         string `split_words:"true" default:"scoircsvjson"`
	Environment     string `split_words:"true" required:"true"`
	LogLevel        string `split_words:"true" default:"info"`
	TestCSVName     string `split_words:"true" default:"test.csv"`
	MinCSVRecordLen int    `split_words:"true" default:"6"`
	SkipCSVHeader   bool   `split_words:"true" default:"true"`
}

// GCP is used for Google Cloud Project
type GCP struct {
	GcpApplicationCredentials string `split_words:"true" required:"false"`
	GcpProjectName            string `split_words:"true" required:"true"`
	GcpInputBucketName        string `split_words:"true" required:"true"`
	GcpOutputBucketName       string `split_words:"true" required:"true"`
}

//GetProcessedConfiguration returns the processed fields for the application configuration.
func GetProcessedConfiguration() (*Configuration, error) {
	configuration := &Configuration{}
	processError := envconfig.Process("SCOIR", configuration)
	return configuration, processError
}
