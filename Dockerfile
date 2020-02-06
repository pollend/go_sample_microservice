# swagger builder -------------------------------------------------------------------
FROM quay.io/goswagger/swagger as swagger-builder
WORKDIR /app
COPY . .
RUN swagger generate spec -m -o /app/swagger.json

# golang builder -------------------------------------------------------------------
FROM golang:1.12.16-alpine as builder
RUN apk add --no-cache ca-certificates git
RUN apk add --no-cache ca-certificates make

RUN go get github.com/google/wire/cmd/wire
WORKDIR /go/src

COPY go.mod .
RUN export GO111MODULE=on && go mod download

COPY . .
RUN export GO111MODULE=on && /go/bin/wire gen
RUN export GO111MODULE=on && go build -o ./out/app wire_gen.go

# release container -------------------------------------------------------------------
FROM alpine as release
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/out/app /app
COPY --from=builder /go/src/public /public
COPY --from=swagger-builder /app/swagger.json ./public/docs/swagger.json

EXPOSE 5050
WORKDIR /
CMD ["/app"]
