FROM golang:1.15

ENV GO111MODULE=off

RUN go get k8s.io/code-generator; exit 0
WORKDIR /go/src/k8s.io/code-generator
RUN go get -d ./...

RUN mkdir -p /go/src/github.com/Marcos30004347/kubernetes-posts
VOLUME /go/src/github.com/Marcos30004347/kubernetes-posts

