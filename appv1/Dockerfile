FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN go build -tags netgo -a -o /hello_http

FROM gcr.io/distroless/base-debian11:debug

WORKDIR /app
COPY --from=builder /hello_http /app/hello_http
USER nonroot
ENTRYPOINT ["/app/hello_http"]
