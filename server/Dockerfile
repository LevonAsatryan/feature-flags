# syntax=docker/dockerfile:1

FROM golang:1.22.2

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
COPY .env .env
RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy the source code.
COPY *.go ./
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /feature-flags

EXPOSE 8080

CMD goose -dir ./db/migrations up && /feature-flags