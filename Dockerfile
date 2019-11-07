FROM golang:1.12-alpine3.9 as build
RUN apk add --no-cache git
WORKDIR $GOPATH/src/github.com/ironcore864/actsctrl
COPY . .
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o actsctrl

FROM alpine:3.9
WORKDIR /app
COPY --from=build /go/src/github.com/ironcore864/actsctrl/actsctrl /usr/local/bin/actsctrl
ENTRYPOINT ["actsctrl"]
