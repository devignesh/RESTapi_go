FROM golang:1.12.7-alpine3.9
MAINTAINER vignesh

RUN mkdir /app

ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...
CMD ["./main"]

# FROM golang:1.12.7-alpine3.9
# RUN mkdir /app
# ADD . /app
# WORKDIR /app
# RUN go build -o main .
# CMD ["/app/main"]