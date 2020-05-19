
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git
ENV GOPROXY=https://goproxy.cn
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/holidayapp /app
ENTRYPOINT ./app
LABEL Name=holidayapp Version=0.0.1
EXPOSE 3000
