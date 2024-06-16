##############################
#           build            #
##############################
FROM golang:1.22.3-alpine AS  builder

WORKDIR /build

RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main cmd/main.go


##############################
#       production           #
##############################
FROM alpine:3.15 as runner

WORKDIR /app

COPY --from=builder /build/main /build/migrate ./
RUN mkdir schema
COPY ./schema ./schema