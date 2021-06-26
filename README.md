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

You know what they say about assumptions

## How To

### Table of Contents
- [Environment Variables](#environment-variables)
- [Go Setup](#go-setup)
- [How to Run](#how-to-run-locally)
- [Terraform](#terraform)
- [Deploying](#deploying)
- [Contribution Guide](#contribution-guide)

#### Environment Variables
- `SCOIR_APP_NAME` (Default: `"scoircsvjson"`)
- `SCOIR_ENVIRONMENT` (Required)
- `SCOIR_GCP_PROJECT_NAME` (Required)
- `SCOIR_GCP_INPUT_BUCKET_NAME` (Required)
- `SCOIR_GCP_OUTPUT_BUCKET_NAME` (Required)

#### Go Setup
* This repository is making use of Go Modules, as such this repository should NOT be placed in your `GOPATH` or you should force `GO111MODULE=on`.

#### How to Run Locally [https://github.com/GoogleCloudPlatform/functions-framework-go#features](https://github.com/GoogleCloudPlatform/functions-framework-go#features)
1. Bare Metal:
   `export SCOIR_ENVIRONMENT="LOCAL" && SCOIR_GCP_PROJECT_NAME="scoir" && SCOIR_GCP_INPUT_BUCKET_NAME="scoir-csv" && SCOIR_GCP_OUTPUT_BUCKET_NAME="scoir-json" && go run ./app/cmd/main.go`

#### Terraform

#### Deploying
##### Manual Deploy Examples:
1. Install gcloud cli: [https://cloud.google.com/sdk/gcloud](https://cloud.google.com/sdk/gcloud)
2. Set project: `gcloud config set project <project-id>` (scoir for dev)
3. Deploy Main Processor:
   * DEV: `gcloud functions deploy scoir-csv-to-json --max-instances 15 --memory 512MB --timeout 540 --entry-point Main --runtime go113 --trigger-bucket scoir-csv --region us-east1 --service-account scoir-318015@appspot.gserviceaccount.com --ingress-settings all --env-vars-file ./env/dev.yaml --set-build-env-vars VERSION=main.0`

#### Contribution Guide
Please see the contribution guide in the [./.github/CONTRIBUTING.md](./.github/CONTRIBUTING.md) file
