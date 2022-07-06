##
## Development
##
FROM golang:1.18.2 as dev


##
## Builder
##
FROM golang:1.18.3-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

##
## Deploy
##
FROM alpine as prod

WORKDIR /app

COPY --from=builder /app/.env ./.env
COPY --from=builder /app/templates/* ./templates/
COPY --from=builder /app/main ./

ENV GIN_MODE release

ENTRYPOINT ["./main"]