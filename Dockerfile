# Backend Dockerfile
FROM golang:1.23 AS build
WORKDIR /src
COPY bin bin
COPY cmd cmd
COPY internal internal
COPY go.mod go.mod 
COPY go.sum go.sum
COPY Makefile Makefile

ENV GOARCH=amd64
RUN CGO_ENABLED=0 go build -o ./bin/apiserver -v ./cmd/apiserver


FROM alpine:3.20.3
WORKDIR /
COPY --from=build /src/bin/apiserver /apiserver

CMD ["/apiserver"]