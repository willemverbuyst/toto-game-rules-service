#
## Build
##

FROM golang:1.19-buster AS build

WORKDIR /service

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /go-service

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go-service /go-service
COPY .env .

EXPOSE 9090

USER nonroot:nonroot

ENTRYPOINT ["/go-service"]