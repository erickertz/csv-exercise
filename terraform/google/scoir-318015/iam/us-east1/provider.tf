provider "google" {
  project = "scoir-318015"
}

terraform {
  required_providers {
    google = {
      version = "~> 3.52.0"
    }
  }
}
