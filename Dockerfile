FROM golang:1.23-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o server ./cmd/app


FROM gcr.io/distroless/base-debian12:debug
WORKDIR /app
COPY --from=builder /src/server .
EXPOSE 8080
ENTRYPOINT ["./server"]
