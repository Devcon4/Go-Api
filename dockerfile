FROM golang:alpine as build-env

RUN mkdir /dist
WORKDIR /dist

COPY ./src/go.mod .
COPY ./src/go.sum .

RUN go mod download

COPY ./src/ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/server

FROM alpine
COPY --from=build-env /go/bin/server /go/bin/server

COPY start.sh .
COPY wait-for-database.sh .

EXPOSE 80
ENTRYPOINT ./start.sh