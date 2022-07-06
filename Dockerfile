FROM golang:1.18.3-stretch as builder

RUN mkdir -p /go/src/server
WORKDIR /go/src/server

COPY go.mod go.sum ./
RUN go mod download && go mod verify && go mod tidy

COPY . .

RUN CGO_ENABLED=0 go build server.go

FROM alpine:edge

RUN mkdir /app
COPY --from=builder /go/src/server/server /app/server
CMD ["/app/server"]
