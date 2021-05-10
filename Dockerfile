# Build stage
FROM golang:alpine as builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn
COPY . .
RUN go get github.com/cespare/reflex

RUN go build -o /go/bin/goframework .

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/bin/goframework .
EXPOSE 8080
CMD [ "/app/goframework" ]
