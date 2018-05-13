FROM golang:latest

WORKDIR /go/src/app/

COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest

WORKDIR /root/
COPY --from=0 /go/src/app/app .
CMD ["./app"]
