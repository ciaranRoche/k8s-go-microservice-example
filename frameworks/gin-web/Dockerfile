FROM golang:1.8.1-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git 
RUN go get github.com/gin-gonic/gin

ENV SOURCES /go/src/github.com/ciaranRoche/advanced-cloud-native-go/frameworks/gin-web/
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

WORKDIR ${SOURCES}
CMD ${SOURCES}gin-web
EXPOSE 8080