FROM golang:1.12.5-alpine3.9

RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . .

RUN apk update && apk upgrade && apk add --no-cache bash git openssh
RUN go get -v -u github.com/gorilla/mux
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./app"]