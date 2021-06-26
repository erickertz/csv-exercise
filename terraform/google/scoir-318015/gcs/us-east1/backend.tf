terraform {
  backend "gcs" {
    bucket = "terraform"
    prefix = "scoir-gcs.tfstate"
  }
}
