FROM golang:alpine

WORKDIR /go/src/app

COPY . .

RUN apk add --no-cache git \
    && go get -d -v ./... \
    && go install -v ./...

CMD ["echo"]
