FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

ENV APP_ENV=production

RUN CGO_ENABLED=0 GOOS=linux go build -o /gin-test

RUN go test -v ./...

EXPOSE 8081

CMD ["/gin-test"]