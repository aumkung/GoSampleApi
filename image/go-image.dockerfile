FROM golang:1.13-alpine
# WORKDIR /go/src/GoSampleApi
# RUN go get github.com/pilu/fresh
# CMD go run main.go
# EXPOSE 1323

# RUN go get github.com/gin-gonic/gin
# RUN go install github.com/gin-gonic/gin

# create image from the official Go image
# FROM golang:alpine
RUN apk add --update tzdata bash wget curl git;
# Create binary directory, install glide and fresh
RUN mkdir -p $$GOPATH/bin && curl https://glide.sh/get | sh && go get github.com/pilu/fresh
# define work directory
ADD . /go/src/GoSampleApi
WORKDIR /go/src/GoSampleApi
CMD fresh main.go
# COPY . .
# serve the app
# CMD glide update && fresh -c runner.conf main.go