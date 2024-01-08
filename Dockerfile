FROM golang:1.21-alpine AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o server ./cmd/server/main.go

FROM alpine
WORKDIR /build
COPY --from=builder /build/server .
CMD ["./server"]
