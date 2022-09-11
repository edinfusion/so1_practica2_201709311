FROM golang:latest  AS builder
RUN apt-get update
#NV GO111MODULE=on \
#   CGO_ENABLED=0  \
#   GOOS=linux \
#   GOARCH=amd64 
WORKDIR /go/src/backend/api
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
#RUN go install


CMD ["go", "run", "main.go"]