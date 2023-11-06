# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:alpine3.18 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /product_svc ./cmd/*.go

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /product_svc /product_svc

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/product_svc"]% 