#
## Build
##

FROM golang:1.19-buster AS build

WORKDIR /server

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY *.go ./

RUN go build -o /go-server

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go-server /go-server

EXPOSE 9090

USER nonroot:nonroot

ENTRYPOINT ["/go-server"]