FROM golang:1.15.7-buster

COPY . /go/

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV GOPATH=""

CMD ["make", "run"]