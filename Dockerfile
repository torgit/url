FROM golang:1.11

WORKDIR /go/src/app
COPY . .

RUN go get

RUN go build

CMD ["app"]