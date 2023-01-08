FROM golang:1.19-alpine
RUN apk add build-base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o tsg68-api github.com/jeremyawarren15/tsg68-api

EXPOSE 8090

ENTRYPOINT ["./tsg68-api", "serve", "--http=0.0.0.0:8090"]
