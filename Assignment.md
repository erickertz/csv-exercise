# Assignment
## User Stories

### As a User, given a csv, I can produce a json file of valid csv records
* csv files will be placed into _`input-directory`_
    * once the application starts it watches _`input-directory`_ for any new files that need to be processed
    * files will be considered new if the file name has not been recorded as processed before.
    * file names will end in `.csv`
    * once a file has been processed, the system deletes it from the _`input-directory`_.
* csv columns and validation
    * csv files will have headers.
    * columns
        1. `INTERNAL_ID` : 8 digit positive integer. Cannot be empty.
        1. `FIRST_NAME` : 15 character max string. Cannot be empty.
        1. `MIDDLE_NAME` : 15 character max string. Can be empty.
        1. `LAST_NAME` : 15 character max string. Cannot be empty.
        1. `PHONE_NUM` : string that matches this pattern `###-###-####`. Cannot be empty.
* json files should be output to _`output-directory`_
    * file name should be the same name as the csv file with a `.json` extension
    * in the event of file name collision, the latest file should overwrite the earlier version.
    * The middle field should not exist if there is no middle name.
* json format:
```js
[
    {
        "id": <INTERNAL_ID>,
        "name": {
            "first": "<FIRST_NAME>",
            "middle": "<MIDDLE_NAME>",
            "last": "<LAST_NAME>"
        },
        "phone": "<PHONE_NUM>"
    }
]
```

#### Example

input of:

```
INTERNAL_ID,FIRST_NAME,MIDDLE_NAME,LAST_NAME,PHONE_NUM
12345678,Bobby,,Tables,555-555-5555
```

would produce:

```json
[
    {
        "id": 12345678,
        "name": {
            "first": "Bobby",
            "last": "Tables"
        },
        "phone": "555-555-5555"
    }
]
```
---

### As a User, I can produce a csv file containing validation errors
* error records should be written to a csv file in _`error-directory`_
* if errors exist, one error file is created per input file.
* processing should continue in the event of an invalid row; all errors should be collected and added to the corresponding error csv.
* an error record should contain:
    1. `LINE_NUM` : the number of the record which was invalid
    * `ERROR_MSG` : a human readable error message about what validation failed
* in the event of name collision, the latest file should overwrite the earlier version.
* the error file should match the name of the input file.

---

### As a User, I can configure input, output, and error directories

# Assignment Delivery / Response

1. csv files will be placed into _`input-directory`_
    - I used Google Cloud Storage Buckets instead of a system directory
    
2. once the application starts it watches _`input-directory`_ for any new files that need to be processed
    - I setup an google.storage.object.finalize event trigger to push messages to PubSub whenever a new file is written to the bucket. A Google Cloud Function subscribes to a PubSub topic and retrieves the messages for processing.
    
3. file names will end in `.csv`
    - I added a csv file name validator that will log fatal should a non csv file attempt to be processed.

4. files will be considered new if the file name has not been recorded as processed before
    - I setup a retention policy on the bucket that disallows overwriting or deleting of files, all files will be new.
    - https://github.com/erickertz/csv-exercise/blob/master/terraform/google/scoir-318015/gcs/us-east1/storage_bucket.tf#L79

5. once a file has been processed, the system deletes it from the _`input-directory`_.
    - the retention policy I setup on the bucket will delete the files after X amount of days, there is no reason for the application to worry about cleanup.
    - https://github.com/erickertz/csv-exercise/blob/master/terraform/google/scoir-318015/gcs/us-east1/storage_bucket.tf#L79
    
6. csv files will have headers, columns validation
    - I made the application configurable to handle files with our without headers
    - I created to csv validator to check for all of the specified conditions
    
7. json files should be output to _`output-directory`_
    - I used Google Cloud Storage Buckets instead of a system directory
    
8. in the event of file name collision, the latest file should overwrite the earlier version.
    - The retention policy I set on the destination bucket prevents collision. The application will also log error should a collision occur.
    - https://github.com/erickertz/csv-exercise/blob/master/terraform/google/scoir-318015/gcs/us-east1/storage_bucket.tf#L97
    
9. The middle field should not exist if there is no middle name.
    - omitempty was added to the json struct for middle field
    
10. error records should be written to a csv file in _`error-directory`_
    - I thought about this in combination with the requirement of "do your best to treat it like a production system". Realistically, most production systems would log in JSON format. Also, since this lives in GCP I just log to stderr so we can use Stackdrvier / Cloud Logging.
    - https://cloudlogging.app.goo.gl/rZeBLAiuua1R7MpS8
    
11. if errors exist, one error file is created per input file.
    - I added the file name to the JSON logging so we can filter logs per file that was processed in Cloud Logging
    
12. processing should continue in the event of an invalid row; all errors should be collected and added to the corresponding error csv.
    - I created a CSV validator that logs error and continues file processing on a row by row basis. See above for collecting / grouping errors per file.
    
13. an error record should contain `LINE_NUM` and `ERROR_MSG`
    - Since we are logging to Cloud Logging I stuck to their logging structure
    - https://cloud.google.com/logging/docs/structured-logging
    - https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogEntrySourceLocation
    - I was sure to include both of these fields: https://github.com/erickertz/csv-exercise/blob/master/main.go#L43
      
14. in the event of name collision, the latest file should overwrite the earlier version
    - The retention policy I set on the destination bucket will not allow collisions. In the event that a collision is detected the application will log error. 
    - https://github.com/erickertz/csv-exercise/blob/master/terraform/google/scoir-318015/gcs/us-east1/storage_bucket.tf#L97
      