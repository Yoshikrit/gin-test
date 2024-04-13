FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV APP_ENV=production

RUN CGO_ENABLED=0 GOOS=linux go build -o /gin-test

EXPOSE 8081

USER nonroot:nonroot

CMD ["/gin-test"]