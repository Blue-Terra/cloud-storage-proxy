</br>
<p align="center">
    <img
        style="width:20%;height:auto;"
        src="./docs/img/logo.png">
    </img>
    <div align="center">
        <h3 style="font-size:26px;line-height:40px">
            Blue Terra Engineering
            <br/>
            <br/>
            ☁️ Cloud Storage Proxy 
            <br/>
        </h3>
        <h4>
            <br/>
            <br/>
            🚦 Warning: This Blue Terra Repository Is Currently In Development
            <br/>
            <br/>
            Use This Code At Your Own Risk
        </h4>
        <br/>
    </div>
</p>


# I. Introduction

This is boilerplate for a containerized proxy service for GCP Cloud Storage written in golang. 

It is useful if you need access to cloud storage blobs on a public client. 

# II. Requirements

1) [Golang](https://go.dev/)
2) [GCP Cloud Storage Service Account](https://cloud.google.com/storage)

# III. Configuration

1) Add your service account to the top level of the directory as a filed named: `serviceAccountKey.json`. 
2) Add add these variables to your enviroment:

    - "PORT=<SOME_PORT>"
    - "SERVICE-API-KEY=<SOME_API_KEY>"
    
# IV. Service Overview

This is a REST Api that retrieves blobs from GCP Storage. It contains a single endpoint:

1) getBlob - takes in `BUCKET` and `BLOB` string fields in the header, queries gcs for the blobs 
data and returns the data from the blob in the response body. 

The `API-KEY` request header field should be passed into all requests for authentication.

# 

# V. Development

To run this service run: 

    go run ./src/main.go 

# VI. Building 

To build the service run:

    go build ./src/main.go 

# VII. Docker

We have included a `docker-compose.yaml` file at the top level.

```
version: "3.9"
services:
  cloud-stroage-proxy:
    build:
      context: "."
      dockerfile: "Dockerfile"
    ports:
      - "8080:8080"
    environment:
      - "PORT=8080"
      - "SERVICE-API-KEY=86c5ceb27e1bf441130299c0209e5f35b88089f62c06b2b09d65772274f12057"
```

To docker compose execute your service you can run:

    docker-compose up --build

after initializing your Docker daemon.

