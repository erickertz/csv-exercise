resource "google_cloudfunctions_function" "tfer--us-002D-east1_scoir-002D-csv-002D-to-002D-json" {
  available_memory_mb = "512"
  entry_point         = "Main"

  environment_variables = yamldecode(file("../../../../../env/dev.yaml"))

  event_trigger {
    event_type = "google.storage.object.finalize"

    failure_policy {
      retry = "false"
    }

    resource = "projects/_/buckets/scoir-csv"
  }

  ingress_settings = "ALLOW_ALL"

  labels = {
    deployment-tool = "cli-gcloud"
  }

  max_instances         = "1"
  name                  = "scoir-csv-to-json"
  project               = "scoir-318015"
  region                = "us-east1"
  runtime               = "go113"
  service_account_email = "scoir-318015@appspot.gserviceaccount.com"
  timeout               = "540"
}
