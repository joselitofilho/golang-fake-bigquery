# Start from golang base image
FROM golang:alpine as builder

# Enable go modules
ENV GO111MODULE=on

# Install git. (alpine image does not have git in it)
RUN apk update && apk add --no-cache git curl

# Set current working directory
WORKDIR /app

# Note here: To avoid downloading dependencies every time we
# build image. Here, we are caching all the dependencies by
# first copying go.mod and go.sum files and downloading them,
# to be used every time we build the image if the dependencies
# are not changed.

# Copy go mod and sum files
COPY go.mod ./
# COPY go.sum ./

# Download all dependencies.
RUN go mod download

# Now, copy the source code
COPY . .

# Note here: CGO_ENABLED is disabled for cross system compilation
# It is also a common best practise.

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o golang-fake-bigquery cmd/golang-fake-bigquery/main.go

# Get the discovery file
RUN curl https://www.googleapis.com/discovery/v1/apis/bigquery/v2/rest > discovery.json

# Finally our multi-stage to build a small image
# Start a new stage from scratch
FROM scratch

# Copy the Pre-built binary file
COPY --from=builder /app/golang-fake-bigquery .

# Copy the discovery file
COPY --from=builder /app/discovery.json .

# Run executable
CMD ["./golang-fake-bigquery", "-discovery-json-path", "discovery.json"]