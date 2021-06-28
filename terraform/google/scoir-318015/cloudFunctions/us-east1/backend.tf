terraform {
  backend "gcs" {
    bucket = "scior-terraform"
    prefix = "scoir-cloudRunctions.tfstate"
  }
}
