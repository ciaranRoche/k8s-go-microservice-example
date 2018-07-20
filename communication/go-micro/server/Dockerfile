FROM golang:1.8.1-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

RUN go get github.com/micro/micro && \
    go get github.com/micro/protobuf/proto && \
    go get github.com/micro/protobuf/protoc-gen-go

ENV SOURCES /go/src/github.com/ciaranRoche/advanced-cloud-native-go/communication/go-micro/server/
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

ENV CONSUL_HTTP_ADDR localhost:8500

WORKDIR ${SOURCES}
CMD ${SOURCES}server
