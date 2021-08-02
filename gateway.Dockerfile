##
## Build
##

FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /gateway cmd/gateway/main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /gateway /gateway

EXPOSE 8090

USER nonroot:nonroot

ENTRYPOINT ["/gateway"]
