FROM golang:1.18-buster as development

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go install github.com/cespare/reflex@latest 

