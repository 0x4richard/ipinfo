# run with:
# docker build -t ipinfo .
# docker run -it ipinfo:latest /bin/sh

FROM golang:1.20 as build 
WORKDIR /go/src/github.com/0x4richard/ipinfo/
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:3.15
WORKDIR /go/src/github.com/0x4richard/ipinfo/
COPY --from=build /go/src/github.com/0x4richard/ipinfo ./
COPY --from=build /go/src/github.com/0x4richard/ipinfo/ipinfo /bin/

ENTRYPOINT [ "ipinfo" ]
