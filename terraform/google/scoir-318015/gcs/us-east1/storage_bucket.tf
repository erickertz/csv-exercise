resource "google_storage_bucket" "tfer--scior-002D-terraform" {
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
