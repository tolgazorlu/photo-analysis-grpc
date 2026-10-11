<a name="readme-top"></a>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <h3 align="center">PHOTO ANALYSIS</h3>
  <p align="center">
    BY TOLGA ZORLU
</div>

<!-- Features -->

## About

A gRPC backend that stores photos in Postgres and runs face/emotion detection on them with the Google Cloud Vision API, with image analysis queued through Kafka. Built January-February 2024.

## Features

- <strong>FACE DETECTION</strong>

- <strong>UPLOAD IMAGE</strong>

- <strong>UPDATE IMAGE DETAIL</strong>

- <strong>GET IMAGE DETAIL</strong>

- <strong>GET IMAGE FEED</strong>

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Technology Stack

- <strong>Docker:</strong> Container (postgres:16.1-alpine)
- <strong>Postgres16</strong> Database (postgres:16.1-alpine)
- <strong>Go</strong> Scability and Fast Backend
- <strong>gRPC</strong> Remote Procedure Call
- <strong>Protobuf: </strong> Structered Data
- <strong>Kafka:</strong> Async detection
- <strong>Google Vision API:</strong> Faces Detection

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

This is a gRPC backend (server + client) for photo upload and face analysis.
To get a local copy up and running follow these simple example steps.

### Prerequisites

- [Go](https://go.dev/doc/install)
- A running Postgres instance (the server connects on `localhost`)
- `protoc` with the Go and Go-gRPC plugins, to regenerate `proto/*.pb.go` via `make protoc`
- Google Cloud Vision API credentials (`GOOGLE_APPLICATION_CREDENTIALS`) for face detection
- A running Kafka broker if async detection is enabled (`getting-started.properties`)

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/tolgazorlu/photo-analysis-grpc
   ```
2. Enter the project directory
   ```sh
   cd photo-analysis-grpc
   ```
3. Build the server and client binaries
   ```sh
   make build
   ```
4. Run the server and client in separate terminals
   ```sh
   ./bin/server
   ```
   ```sh
   ./bin/client
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Architecture

```
proto/images.proto   # gRPC service and message definitions
server/               # gRPC server: DB, Vision API face detection, Kafka
client/               # gRPC client calling upload/update/get endpoints
Makefile              # protoc generation and build targets
```

## Status

Learning/experiment project, not actively maintained.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
