resource "google_storage_notification" "tfer--scoir-002D-csv-002F-cloud-002D-functions-002D-scoir-002D-318015-002D-f681c495f6625f20ap-002D-tp" {
  bucket         = "scoir-csv"
  payload_format = "JSON_API_V1"
  topic          = "//pubsub.googleapis.com/projects/f681c495f6625f20ap-tp/topics/cloud-functions-sbvmyxrfmufpgjr3yv7bx46bbi"
}
