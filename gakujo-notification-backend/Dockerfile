FROM golang:1.18 AS build

WORKDIR /work
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN \
    GOOS=linux \
    go build -o app .

FROM ubuntu:22.04

WORKDIR /work
COPY --from=build /work/app /work

ENTRYPOINT [ "/work/app" ]
