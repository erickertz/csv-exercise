resource "google_storage_bucket_iam_binding" "tfer--gcf-002D-sources-002D-1026843452605-002D-us-002D-central1" {
  bucket = "b/gcf-sources-1026843452605-us-central1"
}

resource "google_storage_bucket_iam_binding" "tfer--gcf-002D-sources-002D-1026843452605-002D-us-002D-east1" {
  bucket = "b/gcf-sources-1026843452605-us-east1"
}

resource "google_storage_bucket_iam_binding" "tfer--gcf-002D-sources-002D-1026843452605-002D-us-002D-west2" {
  bucket = "b/gcf-sources-1026843452605-us-west2"
}

resource "google_storage_bucket_iam_binding" "tfer--scior-002D-terraform" {
  bucket = "b/scior-terraform"
}

resource "google_storage_bucket_iam_binding" "tfer--scoir-002D-csv" {
  bucket = "b/scoir-csv"
}

resource "google_storage_bucket_iam_binding" "tfer--scoir-002D-json" {
  bucket = "b/scoir-json"
}

resource "google_storage_bucket_iam_binding" "tfer--us-002E-artifacts-002E-scoir-002D-318015-002E-appspot-002E-com" {
  bucket = "b/us.artifacts.scoir-318015.appspot.com"
}
