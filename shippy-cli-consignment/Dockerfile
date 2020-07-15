
FROM golang:1.3 as builder

ENV GO111MODULE=on

RUN go get -u github.com/BelkevichAndry/go-microservice/shippy-cli-consignment

WORKDIR /go/src/github.com/BelkevichAndry/go-microservice/shippy-cli-consignment

RUN go get && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY consignment.json /app/consignment.json
COPY --from=builder /go/src/github.com/BelkevichAndry/go-microservice/shippy-cli-consignment .

ENTRYPOINT ["./shippy-cli-consignment", "consignment.json"]
CMD ["./shippy-cli-consignment"]

