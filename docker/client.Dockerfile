FROM golang:1.18 AS modules
ADD go.mod go.sum /m/
RUN cd /m; go mod download

FROM golang:1.18 AS builder
COPY --from=modules /go/pkg /go/pkg
WORKDIR /build
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/client

FROM scratch
COPY --from=builder /build/main /

CMD ["/main"]