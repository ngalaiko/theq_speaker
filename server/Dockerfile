FROM golang:1.8

RUN mkdir server
WORKDIR ./server
ADD . ./

RUN go get -u github.com/golang/dep/...

RUN dep ensure
RUN go build

ENTRYPOINT ["./server"]
EXPOSE 7080