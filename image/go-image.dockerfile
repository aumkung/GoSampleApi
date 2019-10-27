FROM golang:1.13-alpine

RUN apk add --update git curl && \
    mkdir -p $GOPATH/bin && \
    curl -fLo /bin/air https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air && \
    chmod +x /bin/air

COPY . .
WORKDIR /go/src/GoSampleApi

CMD air -c .air.conf