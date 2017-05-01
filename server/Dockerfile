FROM golang:1.8

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

WORKDIR /go/src/github.com/ngalayko/theq_speaker/server
ADD . ./

RUN go get -u github.com/golang/dep/...
RUN dep ensure

RUN go build

ENTRYPOINT ["./server"]
EXPOSE 7080