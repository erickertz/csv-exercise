terraform {
  backend "gcs" {
    bucket = "scior-terraform"
    prefix = "scoir-gcs.tfstate"
  }
}
