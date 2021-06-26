terraform {
  backend "gcs" {
    bucket = "terraform"
    prefix = "scoir-iam.tfstate"
  }
}
