FROM golang:1.11.2 AS build
WORKDIR /src
COPY cmd/inspopular/main.go .
RUN go get -d && CGO_ENABLED=0 GOOS=linux go build -o inspopular main.go

FROM alpine:3.8
LABEL maintainer="d94.zaragoza@gmail.com"
WORKDIR /app
COPY --from=build /src/inspopular .
RUN apk add --no-cache ca-certificates 
ENTRYPOINT ["./inspopular"]
