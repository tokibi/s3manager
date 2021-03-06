# S3 Manager

[![Go Report Card](https://goreportcard.com/badge/github.com/mastertinner/s3manager?style=flat-square)](https://goreportcard.com/report/github.com/mastertinner/s3manager)
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fmastertinner%2Fs3manager%2Fbadge&style=flat-square)](https://github.com/mastertinner/s3manager/actions)
[![Docker Build](https://img.shields.io/docker/cloud/build/mastertinner/s3manager.svg?style=flat-square)](https://hub.docker.com/r/mastertinner/s3manager)

A Web GUI written in Go to manage S3 buckets from any provider.

![Screenshot](https://raw.githubusercontent.com/mastertinner/s3manager/master/screenshot.png)

## Usage

### Configuration

The application can be configured with the following environment variables:

- `ENDPOINT`: The endpoint of your S3 server (defaults to `s3.amazonaws.com`)
- `ACCESS_KEY_ID`: Your S3 access key ID (required)
- `SECRET_ACCESS_KEY`: Your S3 secret access key (required)
- `USE_SSL`: Whether your S3 server uses SSL or not (defaults to `true`)
- `PORT`: The port the s3manager app should listen on (defaults to `8080`)

### Build and Run Locally

1.  Run `make build`
1.  Execute the created binary and visit <http://localhost:8080>

### Run Docker image

1. Run `docker run -p 8080:8080 -e 'ACCESS_KEY_ID=XXX' -e 'SECRET_ACCESS_KEY=xxx' mastertinner/s3manager`

### Deploy to Cloud Foundry

1.  Modify `deployments/cf/*` to your liking
1.  Run `make deploy-cf`

## Development

### Lint Code

1. Run `make lint`

### Run Tests

1.  Run `make test`

### Build Docker Image

The image is available on [Docker Hub](https://hub.docker.com/r/mastertinner/s3manager/)

1.  Run `make build-docker`

### Run Locally for Testing

There is an example [docker-compose.yml](https://github.com/mastertinner/s3manager/blob/master/docker-compose.yml) file that spins up an S3 service and the s3manager. You can try it by issuing the following command:

```shell
$ docker-compose up
```
