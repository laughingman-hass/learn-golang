FROM golang:1.15.5-alpine

WORKDIR /usr/src/app

RUN apk add --no-cache build-base git

