FROM golang:1.22.2-alpine3.19 as builder

COPY go.mod go.sum /go/src/github.com/SurkovIlya/message-handler-app/
WORKDIR /go/src/github.com/SurkovIlya/message-handler-app
RUN go mod download
COPY . /go/src/github.com/SurkovIlya/message-handler-app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/message-handler-app github.com/SurkovIlya/message-handler-app


FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/SurkovIlya/message-handler-app/build/message-handler-app /usr/bin/message-handler-app

EXPOSE 8080 8080

ENTRYPOINT ["/usr/bin/message-handler-app"]