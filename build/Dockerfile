FROM golang:alpine as build
WORKDIR /src
COPY cmd cmd
COPY bin bin
COPY internal internal
COPY go.mod go.mod
COPY go.sum go.sum
COPY pkg pkg
RUN apk add git && go build -o spike cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/spike /app/

ENTRYPOINT ["/app/spike"]
