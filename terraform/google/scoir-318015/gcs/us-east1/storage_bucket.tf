resource "google_storage_bucket" "tfer--gcf-002D-sources-002D-1026843452605-002D-us-002D-central1" {
  bucket_policy_only = "true"

  cors {
    max_age_seconds = "0"
    method          = ["GET"]
    origin          = ["https://*.cloud.google.com", "https://*.corp.google.com", "https://*.corp.google.com:*"]
  }

  default_event_based_hold    = "false"
  force_destroy               = "false"
  location                    = "US-CENTRAL1"
  name                        = "gcf-sources-1026843452605-us-central1"
  project                     = "scoir-318015"
  requester_pays              = "false"
  storage_class               = "STANDARD"
  uniform_bucket_level_access = "true"
}

resource "google_storage_bucket" "tfer--gcf-002D-sources-002D-1026843452605-002D-us-002D-east1" {
  bucket_policy_only = "true"

  cors {
    max_age_seconds = "0"
    method          = ["GET"]
    origin          = ["https://*.cloud.google.com", "https://*.corp.google.com", "https://*.corp.google.com:*"]
  }

  default_event_based_hold    = "false"
  force_destroy               = "false"
  location                    = "US-EAST1"
  name                        = "gcf-sources-1026843452605-us-east1"
  project                     = "scoir-318015"
  requester_pays              = "false"
  storage_class               = "STANDARD"
  uniform_bucket_level_access = "true"
}

resource "google_storage_bucket" "tfer--gcf-002D-sources-002D-1026843452605-002D-us-002D-west2" {
  bucket_policy_only = "true"

  cors {
    max_age_seconds = "0"
    method          = ["GET"]
    origin          = ["https://*.cloud.google.com", "https://*.corp.google.com", "https://*.corp.google.com:*"]
  }

  default_event_based_hold    = "false"
  force_destroy               = "false"
  location                    = "US-WEST2"
  name                        = "gcf-sources-1026843452605-us-west2"
  project                     = "scoir-318015"
  requester_pays              = "false"
  storage_class               = "STANDARD"
  uniform_bucket_level_access = "true"
}

resource "google_storage_bucket" "tfer--scior-002D-terraform" {
  bucket_policy_only          = "true"
  default_event_based_hold    = "false"
  force_destroy               = "false"
  location                    = "US-EAST1"
  name                        = "scior-terraform"
  project                     = "scoir-318015"
  requester_pays              = "false"
  storage_class               = "STANDARD"
  uniform_bucket_level_access = "true"
}

resource "google_storage_bucket" "tfer--scoir-002D-csv" {
  bucket_policy_only       = "true"
  default_event_based_hold = "false"
  force_destroy            = "false"
  location                 = "US-EAST1"
  name                     = "scoir-csv"
  project                  = "scoir-318015"
  requester_pays           = "false"

  retention_policy {
    is_locked        = "false"
    retention_period = "2678400"
  }

  storage_class               = "STANDARD"
  uniform_bucket_level_access = "true"
}

resource "google_storage_bucket" "tfer--scoir-002D-json" {
  bucket_policy_only       = "true"
  default_event_based_hold = "false"
  force_destroy            = "false"
  location                 = "US-EAST1"
  name                     = "scoir-json"
  project                  = "scoir-318015"
  requester_pays           = "false"

  retention_policy {
    is_locked        = "false"
    retention_period = "2678400"
  }

  storage_class               = "STANDARD"
  uniform_bucket_level_access = "true"
}

resource "google_storage_bucket" "tfer--us-002E-artifacts-002E-scoir-002D-318015-002E-appspot-002E-com" {
  bucket_policy_only          = "false"
  default_event_based_hold    = "false"
  force_destroy               = "false"
  location                    = "US"
  name                        = "us.artifacts.scoir-318015.appspot.com"
  project                     = "scoir-318015"
  requester_pays              = "false"
  storage_class               = "STANDARD"
  uniform_bucket_level_access = "false"
}
