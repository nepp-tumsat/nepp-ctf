# This Dockerfile is refer to https://docs.docker.com/guides/golang/build-images/#multi-stage-builds

# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.23 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /server

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /server /server

EXPOSE 50001

USER nonroot:nonroot

ENTRYPOINT ["/server"]