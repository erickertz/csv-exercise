# SCOIR Technical Interview
This repo contains an exercise intended for engineers.

## Instructions
1. Fork this repo.
1. Using technology of your choice, complete [the assignment](./Assignment.md).
1. Update this README with
    * a [How To](#how-to) section containing any instructions needed to execute your program.
    * an [Assumptions](#assumptions) section containing documentation on any assumptions made while interpreting the requirements.
1. Before the deadline, submit a pull request with your solution.

## Expectations
1. Please take no more than 8 hours to work on this exercise. Complete as much as possible and then submit your solution.
1. This exercise is meant to showcase how you work. With consideration to the time limit, do your best to treat it like a production system.

## Assumptions

You have golang 1.13 installed and or docker / docker-compose and know how to use them.

## How To

[![architecture](https://github.com/erickertz/csv-exercise/blob/master/static/diagram.png)](https://github.com/erickertz/csv-exercise/blob/master/static/scoir.pdf)

### Architecture
https://github.com/erickertz/csv-exercise/blob/master/static/scoir.pdf

### Table of Contents
- [Environment Variables](#environment-variables)
- [Go Setup](#go-setup)
- [How to Run](#how-to-run-locally)
- [Terraform](#terraform)
- [Deploying](#deploying)
- [Contribution Guide](#contribution-guide)
- [Dev Notes](#dev-notes)
- [TODOs](#todos)

#### Environment Variables
- `SCOIR_APP_NAME` (Default: `"scoircsvjson"`)
- `SCOIR_ENVIRONMENT` (Required)
- `SCOIR_GCP_PROJECT_NAME` (Required)
- `SCOIR_GCP_INPUT_BUCKET_NAME` (Required)
- `SCOIR_GCP_OUTPUT_BUCKET_NAME` (Required)

#### Go Setup
* This repository is making use of Go Modules, as such this repository should NOT be placed in your `GOPATH` or you should force `GO111MODULE=on`.

#### How to Run Locally
1. Bare Metal:
   `export SCOIR_ENVIRONMENT="LOCAL" && SCOIR_GCP_PROJECT_NAME="scoir" && SCOIR_GCP_INPUT_BUCKET_NAME="scoir-csv" && SCOIR_GCP_OUTPUT_BUCKET_NAME="scoir-json" && go run ./app/cmd/main.go`
2. Docker:
   1. Set your service account creds as JSON in docker-compose, ex: `SCOIR_GCP_APPLICATION_CREDENTIALS: '{"type": "service_account"}'`
   2. Run docker compose: `docker-compose up --build`

#### Terraform
Terraform is saved in the following file structure: `/terraform/{provider}/{project/{service}/{region}/{resource}.tf`:

#### Deploying
##### Manual Deploy Example:
1. Install gcloud cli: [https://cloud.google.com/sdk/gcloud](https://cloud.google.com/sdk/gcloud)
2. Set project: `gcloud config set project <project-id>` (scoir for dev)
3. Deploy Function using gcloud cli:
   * DEV: `gcloud functions deploy scoir-csv-to-json --max-instances 1 --memory 512MB --timeout 540 --entry-point Main --runtime go113 --trigger-bucket scoir-csv --region us-east1 --service-account scoir-318015@appspot.gserviceaccount.com --ingress-settings all --env-vars-file ./env/dev.yaml --set-build-env-vars VERSION=master.0`

#### Contribution Guide
Please see the contribution guide in the [./.github/CONTRIBUTING.md](./.github/CONTRIBUTING.md) file

#### Dev Notes
- There is no official GCP Cloud Storage emulator so we are kinda forced to use ones in an actual project for local development :( .
- For Cloud Funcs / POC, keep max_instances to 1 to avoid unnecessary concurrent processing. See [TODOs](#todos) for long term proposal.

#### TODOs
- Use Cloud Build and Cloud Run. Cloud Functions are a great fit for this as a POC / Demo but have limitations that might not make this suitable for a production enviroment with large amounts of data to process.
- Have contribution and testing checks at PR / build time
- Write some unit tests there, cowboy
- Chunk up file and use multiple goroutines for faster input file processing
