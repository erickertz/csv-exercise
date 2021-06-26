resource "google_storage_bucket_acl" "tfer--gcf-002D-sources-002D-1026843452605-002D-us-002D-central1" {
  bucket = "gcf-sources-1026843452605-us-central1"
}

resource "google_storage_bucket_acl" "tfer--gcf-002D-sources-002D-1026843452605-002D-us-002D-east1" {
  bucket = "gcf-sources-1026843452605-us-east1"
}

resource "google_storage_bucket_acl" "tfer--gcf-002D-sources-002D-1026843452605-002D-us-002D-west2" {
  bucket = "gcf-sources-1026843452605-us-west2"
}

resource "google_storage_bucket_acl" "tfer--scior-002D-terraform" {
  bucket = "scior-terraform"
}

resource "google_storage_bucket_acl" "tfer--scoir-002D-csv" {
  bucket = "scoir-csv"
}

resource "google_storage_bucket_acl" "tfer--scoir-002D-json" {
  bucket = "scoir-json"
}

resource "google_storage_bucket_acl" "tfer--us-002E-artifacts-002E-scoir-002D-318015-002E-appspot-002E-com" {
  bucket = "us.artifacts.scoir-318015.appspot.com"
}
