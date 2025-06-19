FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache git curl make
RUN go install github.com/air-verse/air@latest
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz -C /usr/local/bin

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p /app/tmp

EXPOSE 8080

CMD ["air", "-c", "/app/air.toml"]
