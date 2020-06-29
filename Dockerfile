FROM golang:1.14-alpine AS build
ENV CGO_ENABLED=0 GOPROXY=https://proxy.golang.org
WORKDIR /golang
RUN apk add --update --no-cache git
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -a -tags netgo -ldflags '-w' -o app cmd/app/main.go

FROM alpine:3.10
ENV TZ=Asia/Bangkok
WORKDIR /go
EXPOSE 8080
RUN apk add --update --no-cache tzdata ca-certificates
COPY --from=build /golang/app .
COPY --from=build /golang/env.yml .
ENTRYPOINT [ "./app" ]