FROM golang:1.13-alpine
WORKDIR /go/src/GoSampleApi
COPY . .

# RUN go get github.com/gin-gonic/gin
# RUN go install github.com/gin-gonic/gin