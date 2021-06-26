terraform {
  backend "gcs" {
    bucket = "terraform"
    prefix = "scoir-cloudRunctions.tfstate"
  }
}
