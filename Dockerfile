# Backend Dockerfile
FROM golang:1.22.1 AS build
WORKDIR /src
COPY bin bin
COPY cmd cmd
COPY internal internal
COPY go.mod go.mod 
COPY go.sum go.sum
COPY Makefile Makefile
COPY .env .env

ENV GOARCH=amd64
RUN CGO_ENABLED=0 make build


FROM alpine
COPY --from=build /src/bin/apiserver /bin/apiserver
COPY config.toml config.toml

CMD ["/bin/apiserver"]