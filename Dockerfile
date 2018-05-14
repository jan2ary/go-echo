FROM golang:latest as builder

WORKDIR /go/src/app/

COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /go/src/app/app .
EXPOSE 9000
CMD ["./app"]
